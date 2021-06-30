# ripplet: cmd/ripplet/*.go pkg/*/*.go internal/*/*.go go.mod
# 	CGO_ENABLED=0 go build ./cmd/ripplet

clean:
	go run scripts/clean/mod.go

antlr:
	go run scripts/antlr/mod.go
