package asm

import (
	"fmt"
	"log"
	"strings"

	parser "github.com/hsiaosiyuan0/ripplet/internal/grammar"
)

type Position struct {
	Line   int
	Column int
}

type UpLocType int

const (
	UP_LOC_LOCAL UpLocType = iota
	UP_LOC_UP
)

type UpLoc struct {
	Typ UpLocType
	At  interface{}
}

type FnShape struct {
	Pos      Position
	LocalCnt int
	Instrs   []int
	Upvals   map[string]UpLoc
	Subs     []*FnShape
}

func NewFnShape() *FnShape {
	return &FnShape{
		Instrs: make([]int, 0),
		Upvals: make(map[string]UpLoc),
		Subs:   make([]*FnShape, 0),
	}
}

func (f *FnShape) Dump(chunk *Chunk) string {
	var b strings.Builder
	for i := 0; i < len(f.Instrs); i++ {
		op := Opcode(f.Instrs[i])
		switch op {
		case CONST, LOAD, STORE, CLOSURE, CONCAT, JMP, JMP_F, NEG:
			fmt.Fprint(&b, op.String()+"\n")
			i++
			fmt.Fprintf(&b, "OPD_%d\n", f.Instrs[i])
		case CALL, ARR:
			fmt.Fprint(&b, op.String()+"\n")
			i++
			fmt.Fprintf(&b, "CNT_%d\n", f.Instrs[i])
		case LOAD_UP, STORE_UP, LOAD_EXT:
			fmt.Fprint(&b, op.String()+"\n")
			i++
			fmt.Fprintf(&b, "NAME_%s\n", chunk.ConstStr(f.Instrs[i]))
		default:
			fmt.Fprintf(&b, "%s\n", op.String())
		}
	}

	for i, sub := range f.Subs {
		fmt.Fprintf(&b, "\nSub: %d\n", i)
		fmt.Fprintf(&b, "%s", sub.Dump(chunk))
	}
	return b.String()
}

type ConstType = uint8

const (
	ConstStr ConstType = iota
	ConstNum
)

type Const struct {
	Typ ConstType
	Val interface{}
}

type Chunk struct {
	Sig    string
	Ver    uint32
	Consts []*Const
	Fn     FnShape
}

func NewChunk() *Chunk {
	return &Chunk{
		Sig:    "ripplet",
		Ver:    0,
		Consts: make([]*Const, 0),
	}
}

func (c *Chunk) ConstStr(i int) string {
	return c.Consts[i].Val.(string)
}

func (c *Chunk) GetConst(i int) *Const {
	return c.Consts[i]
}

func (c *Chunk) IdxOfConstStr(str string) int {
	for i, c := range c.Consts {
		if c.Typ == ConstStr && c.Val == str {
			return i
		}
	}
	return -1
}

func (c *Chunk) IdxOfConstNum(num float64) int {
	for i, c := range c.Consts {
		if c.Typ == ConstNum && c.Val == num {
			return i
		}
	}
	return -1
}

func (c *Chunk) AddConstStr(str string) int {
	if i := c.IdxOfConstStr(str); i != -1 {
		return i
	}

	i := len(c.Consts)
	c.Consts = append(c.Consts, &Const{Typ: ConstStr, Val: str})
	return i
}

func (c *Chunk) AddConstNum(num float64) int {
	if i := c.IdxOfConstNum(num); i != -1 {
		return i
	}

	i := len(c.Consts)
	c.Consts = append(c.Consts, &Const{Typ: ConstNum, Val: num})
	return i
}

func (c *Chunk) Dump() string {
	return c.Fn.Dump(c)
}

func Compile(code string, externals []string) *Chunk {
	tree := parser.Parse(code)

	s := NewSymTabListener(externals)
	symtab, err := s.Resolve(&tree)
	if err != nil {
		log.Fatal(err)
	}

	c := NewCodegenVisitor(symtab)
	c.Visit(tree)
	return c.Finalize()
}
