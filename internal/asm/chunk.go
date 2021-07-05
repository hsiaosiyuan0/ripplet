package asm

import (
	"fmt"
	"strings"
)

type Position struct {
	Line   int
	Column int
}

type FnShape struct {
	Pos    Position
	Instrs []int
	Upvals []string
	Subs   []*FnShape
}

func NewFnShape() *FnShape {
	return &FnShape{
		Instrs: make([]int, 0),
		Upvals: make([]string, 0),
		Subs:   make([]*FnShape, 0),
	}
}

func (f *FnShape) Dump() string {
	var b strings.Builder
	for i := 0; i < len(f.Instrs); i++ {
		op := Opcode(f.Instrs[i])
		switch op {
		case CONST, LOAD, STORE:
			fmt.Fprint(&b, op.String()+"\n")
			i++
			fmt.Fprintf(&b, "INDEX_%d\n", f.Instrs[i])
		default:
			fmt.Fprintf(&b, "%s\n", op.String())
		}
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

func (c *Chunk) AddConstStr(str string) int {
	c.Consts = append(c.Consts, &Const{Typ: ConstStr, Val: str})
	return len(c.Consts)
}

func (c *Chunk) AddConstNum(num float64) int {
	c.Consts = append(c.Consts, &Const{Typ: ConstNum, Val: num})
	return len(c.Consts)
}
