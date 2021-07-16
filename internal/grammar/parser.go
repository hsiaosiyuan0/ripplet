package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func Parse(code string) IProgramContext {
	input := antlr.NewInputStream(code)
	lexer := NewRippletLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewRippletParser(stream)
	p.BuildParseTrees = true
	tree := p.Program()
	return tree
}
