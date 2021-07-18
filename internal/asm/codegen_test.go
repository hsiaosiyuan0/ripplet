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
	OPD_0
	STORE
	OPD_0
	CONST
	OPD_1
	STORE
	OPD_1
`)
}

func TestPrintHello(t *testing.T) {
	assertOpcodesEq(t, `
	a := "hello world"
	print(a)
	`, `
	CONST
	OPD_0
	CONCAT
	OPD_1
	STORE
	OPD_0
	LOAD_EXT
	NAME_print
	LOAD
	OPD_0
	CALL
	CNT_2
`)
}
