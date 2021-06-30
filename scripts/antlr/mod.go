package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	utils "github.com/hsiaosiyuan0/ripplet/scripts/utils"
)

func patchLexer() {
	lexer := "./internal/grammar/ripplet_lexer.go"
	bytes, err := ioutil.ReadFile(lexer)
	if err != nil {
		log.Fatalf("failed to patch lexer: %s", err)
	}
	fixed := strings.Replace(string(bytes), "*RippletLexerBase", "RippletLexerBase", 1)

	err = ioutil.WriteFile(lexer, []byte(fixed), 0644)
	if err != nil {
		log.Fatalf("failed to patch lexer: %s", err)
	}

	fmt.Println("patch lexer: ok")
}

func main() {
	utils.Shell("antlr", "-Dlanguage=Go", "./grammar/RippletLexer.g4", "-o", "./internal/")
	utils.Shell("antlr", "-Dlanguage=Go", "./grammar/RippletParser.g4", "-lib", "./internal/grammar", "-o", "./internal/")
	patchLexer()
}
