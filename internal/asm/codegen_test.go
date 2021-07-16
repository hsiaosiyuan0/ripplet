package asm

import (
	"strings"
	"testing"

	parser "github.com/hsiaosiyuan0/ripplet/internal/grammar"
	utils "github.com/hsiaosiyuan0/ripplet/internal/utils"
)

func assertOpcodesEq(t *testing.T, code string, opcodes string) {
	tree := parser.Parse(code)

	s := NewSymTabListener([]string{"print"})
	symtab, err := s.Resolve(&tree)
	if err != nil {
		t.Fatal(err)
	}

	c := NewCodegenVisitor(symtab)
	c.Visit(tree)

	ops := strings.Split(opcodes, "\n")
	opsTrim := make([]string, 0)
	for _, op := range ops {
		if op == "" {
			continue
		}
		opsTrim = append(opsTrim, strings.TrimSpace(op))
	}

	utils.AssertEqual(t, strings.TrimSpace(c.Finalize().Dump()), strings.Join(opsTrim, "\n"), "")
}

func TestVarDec(t *testing.T) {
	assertOpcodesEq(t, `
	a := 1
	b := 0xff
	`, `
	CONST
	INDEX_0
	STORE
	INDEX_0
	CONST
	INDEX_1
	STORE
	INDEX_1
`)
}

func TestPrintHello(t *testing.T) {
	assertOpcodesEq(t, `
	a := "hello world"
	print(a)
	`, `
	CONST
	INDEX_0
	CONCAT
	INDEX_1
	STORE
	INDEX_0
	LOAD_GL
	NAME_print
	LOAD
	INDEX_0
	CALL
	CNT_2
`)
}
