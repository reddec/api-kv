package dbo

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"strings"

	migrate "github.com/rubenv/sql-migrate"
	_ "modernc.org/sqlite"
)

//go:embed migrations
var migrations embed.FS

func Connect(file string) (*sql.DB, error) {
	q := "_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_time_format=sqlite"
	if strings.Contains(file, ":memory:") {
		q += "&cache=shared"
	}
	db, err := sql.Open("sqlite", file+"?"+q)
	if err != nil {
		return nil, fmt.Errorf("open DB: %w", err)
	}
	_, err = migrate.Exec(db, "sqlite3", migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrations,
		Root:       "migrations",
	}, migrate.Up)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("migrate: %w", err), db.Close())
	}
	return db, nil
}
