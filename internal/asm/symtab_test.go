package asm

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/hsiaosiyuan0/ripplet/internal/grammar"
)

func parse(code string) parser.IProgramContext {
	input := antlr.NewInputStream(code)
	lexer := parser.NewRippletLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewRippletParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Program()
	return tree
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestNestedScopes(t *testing.T) {
	tree := parse(`
	a := 1    				// 0
	fn b() {  				// 0
		c := 1  				// 1
	
		d := () => { 		// 1
			e := 1 				// 2
		}
	}	
	`)

	s := NewSymTabListener()
	symtab, err := s.Resolve(&tree)
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, symtab.Scopes[0].HasBinding("a"), true, "")
	assertEqual(t, symtab.Scopes[0].HasBinding("b"), true, "")
	assertEqual(t, symtab.Scopes[1].HasBinding("c"), true, "")
	assertEqual(t, symtab.Scopes[1].HasBinding("d"), true, "")
	assertEqual(t, symtab.Scopes[2].HasBinding("e"), true, "")
}

func TestUpval(t *testing.T) {
	tree := parse(`
	a := 1    				// 0
	fn b() {  				// 0
		a
	}	
	`)

	s := NewSymTabListener()
	symtab, err := s.Resolve(&tree)
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, symtab.Scopes[0].HasBinding("a"), true, "")
	assertEqual(t, symtab.Scopes[0].HasBinding("b"), true, "")
	assertEqual(t, symtab.Scopes[1].HasLocal("a"), false, "")
}

func TestUndef(t *testing.T) {
	tree := parse(`
	a := 1    				// 0
	fn b() {  				// 0
		c
	}	
	`)

	s := NewSymTabListener()
	_, err := s.Resolve(&tree)
	if err == nil {
		t.Fatal("should not pass")
	}
}
