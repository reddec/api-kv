LOCALBIN := $(PWD)/.local/bin
# tools
GORELEASER := $(LOCALBIN)/goreleaser
OGEN := $(LOCALBIN)/ogen
SQLC := $(LOCALBIN)/sqlc

$(LOCALBIN):
	mkdir -p "$(LOCALBIN)"

$(GORELEASER): $(LOCALBIN)
	GOBIN="$(LOCALBIN)" go install -v github.com/goreleaser/goreleaser@latest

$(OGEN): $(LOCALBIN)
	GOBIN="$(LOCALBIN)" go install -v github.com/ogen-go/ogen/cmd/...@v0.68.4

$(SQLC): $(LOCALBIN)
	GOBIN="$(LOCALBIN)" go install -v github.com/kyleconroy/sqlc/cmd/sqlc@v1.18.0

gen: $(OGEN) $(SQLC)
	go generate ./...
	# generate db
	cd internal/dbo && $(SQLC) generate
	# generate public client
	$(OGEN) --target api/client --package client --no-server --clean  --no-webhook-client --no-webhook-server --convenient-errors=off openapi.yaml
	# generate internal server
	$(OGEN) --target internal/server/api --no-client --clean  --no-webhook-client --no-webhook-server --convenient-errors=off openapi.yaml