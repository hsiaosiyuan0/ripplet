ripplet: cmd/ripplet/*.go pkg/*/*.go internal/*/*.go go.mod
	go build ./cmd/ripplet

clean:
	go run scripts/clean/mod.go

antlr:
	go run scripts/antlr/mod.go

test: cmd/ripplet/*.go pkg/*/*.go internal/*/*.go go.mod
	go test ./...
