package asm

import (
	"testing"

	parser "github.com/hsiaosiyuan0/ripplet/internal/grammar"
	utils "github.com/hsiaosiyuan0/ripplet/internal/utils"
)

func TestNestedScopes(t *testing.T) {
	tree := parser.Parse(`
	a := 1    				// 0
	fn b() {  				// 0
		c := 1  				// 1
	
		d := () => { 		// 1
			e := 1 				// 2
		}
	}	
	`)

	s := NewSymTabListener(make([]string, 0))
	symtab, err := s.Resolve(&tree)
	if err != nil {
		t.Fatal(err)
	}

	utils.AssertEqual(t, symtab.Scopes[1].HasBinding("a"), true, "")
	utils.AssertEqual(t, symtab.Scopes[1].HasBinding("b"), true, "")
	utils.AssertEqual(t, symtab.Scopes[2].HasBinding("c"), true, "")
	utils.AssertEqual(t, symtab.Scopes[2].HasBinding("d"), true, "")
	utils.AssertEqual(t, symtab.Scopes[3].HasBinding("e"), true, "")
}

func TestUpval(t *testing.T) {
	tree := parser.Parse(`
	a := 1    				// 0
	fn b() {  				// 0
		a
	}	
	`)

	s := NewSymTabListener(make([]string, 0))
	symtab, err := s.Resolve(&tree)
	if err != nil {
		t.Fatal(err)
	}

	utils.AssertEqual(t, symtab.Scopes[1].HasBinding("a"), true, "")
	utils.AssertEqual(t, symtab.Scopes[1].HasBinding("b"), true, "")
	utils.AssertEqual(t, symtab.Scopes[2].HasLocal("a"), false, "")
}

func TestUndef(t *testing.T) {
	tree := parser.Parse(`
	a := 1    				// 0
	fn b() {  				// 0
		c
	}	
	`)

	s := NewSymTabListener(make([]string, 0))
	_, err := s.Resolve(&tree)
	if err == nil {
		t.Fatal("should not pass")
	}
}
