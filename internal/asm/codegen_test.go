package asm

import (
	"strings"
	"testing"

	parser "github.com/hsiaosiyuan0/ripplet/internal/grammar"
	utils "github.com/hsiaosiyuan0/ripplet/internal/utils"
)

func trimLines(str string) string {
	lines := strings.Split(str, "\n")
	ret := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		ret = append(ret, strings.TrimSpace(line))
	}
	return strings.Join(ret, "\n")
}

func assertOpcodesEq(t *testing.T, code string, opcodes string) {
	tree := parser.Parse(code)

	s := NewSymTabListener([]string{"print"})
	symtab, err := s.Resolve(&tree)
	if err != nil {
		t.Fatal(err)
	}

	c := NewCodegenVisitor(symtab)
	c.Visit(tree)

	utils.AssertEqual(t, trimLines(c.Finalize().Dump()), trimLines(opcodes), "")
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

func TestStoreArr(t *testing.T) {
	assertOpcodesEq(t, `
  a := [1, 2, 3]
  a[1] = 5
	`, `
  CONST
  OPD_0
  CONST
  OPD_1
  CONST
  OPD_2
  ARR
  CNT_3
  STORE
  OPD_0
  CONST
  OPD_3
  LOAD
  OPD_0
  CONST
  OPD_0
  STORE_ARR
`)
}
