package main

import (
	"fmt"
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	asm "github.com/hsiaosiyuan0/ripplet/internal/asm"
	parser "github.com/hsiaosiyuan0/ripplet/internal/grammar"
	// vm "github.com/hsiaosiyuan0/ripplet/internal/vm"
)

func parse(code string) parser.IProgramContext {
	input := antlr.NewInputStream(code)
	lexer := parser.NewRippletLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewRippletParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Program()
	return tree
}

func main() {
	tree := parse(`
	a := 1 
	b := 2
	c := 3
  d := a + b * c

	// CONST
	// INDEX
	// STORE
	// INDEX
	// CONST
	// INDEX
	// STORE
	// INDEX
	// LOAD
	// INDEX
	// LOAD
	// INDEX
	// CALL
	// INDEX
	`)

	s := asm.NewSymTabListener()
	symtab, err := s.Resolve(&tree)
	if err != nil {
		log.Fatal(err)
	}

	c := asm.NewCodegenListener(symtab)
	c.Resolve(&tree)
	fn := c.Finalize().Fn
	fmt.Println(fn.Dump())

	// vm := vm.NewVm()
	// vm.Exec(&fn)

	// p := parser.NewJSONParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// p.BuildParseTrees = true
}
