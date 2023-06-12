package commands

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/reddec/api-kv/internal/dbo"
	"github.com/reddec/api-kv/internal/impl"
	"github.com/reddec/api-kv/internal/server/api"
)

type ServeCMD struct {
	Bind            string        `short:"b" long:"bind" env:"BIND" description:"API binding address" default:"127.0.0.1:8080"`
	Dir             string        `short:"d" long:"dir" env:"DIR" description:"Directory to store data" default:"data"`
	Token           string        `short:"t" long:"token" env:"TOKEN" description:"Authorization token" default:"changeme"`
	BatchSize       int64         `long:"batch-size" env:"BATCH_SIZE" description:"Batch size for keys iteration" default:"20"`
	CleanupInterval time.Duration `long:"cleanup-interval" env:"CLEANUP_INTERVAL" description:"Cleanup interval for expired items" default:"1m"`
}

func (cmd *ServeCMD) Execute([]string) error {
	if err := os.MkdirAll(cmd.Dir, 0755); err != nil {
		return fmt.Errorf("create dir: %w", err)
	}
	db, err := dbo.Connect("file:" + filepath.Join(cmd.Dir, "values.db"))
	if err != nil {
		return fmt.Errorf("connect to DB: %w", err)
	}
	defer db.Close()
	sqdb := dbo.New(db)
	handler, err := api.NewServer(&impl.Server{
		BatchSize: cmd.BatchSize,
		DB:        sqdb,
	}, impl.StaticToken(cmd.Token))
	if err != nil {
		return fmt.Errorf("create server: %w", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	// run cleanup old expired records
	go func() {
		for range time.NewTicker(cmd.CleanupInterval).C {
			if err := sqdb.DeleteExpired(ctx); err != nil {
				log.Println("delete expired failed:", err)
			}
		}
	}()

	// run API server
	server := &http.Server{Addr: cmd.Bind, Handler: impl.FixContentType(handler)}
	go func() {
		<-ctx.Done()
		_ = server.Close()
	}()
	log.Println("ready on", cmd.Bind)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("serve: %w", err)
	}
	return nil
}
