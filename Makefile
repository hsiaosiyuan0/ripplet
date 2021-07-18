ripplet: cmd/ripplet/*.go internal/*/*.go go.mod
	go build ./cmd/ripplet

install: ./ripplet
	cp ./ripplet /usr/local/bin/ripplet

clean:
	go run scripts/clean/mod.go

antlr:
	go run scripts/antlr/mod.go

test: cmd/ripplet/*.go internal/*/*.go go.mod
	go test ./...
