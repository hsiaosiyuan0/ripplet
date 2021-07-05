package asm

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/hsiaosiyuan0/ripplet/internal/grammar"
)

// FnPlot consist of the infomation collected and required
// for function instructions generation
type FnPlot struct {
	// scope id, used to retrieve its associate scope
	Scope uint

	Shape  *FnShape
	Parent *FnPlot

	// where this object is located at its `Parent.Subs`
	IdxInParent int
	Subs        []*FnPlot
}

func NewFnPlot(scope uint) *FnPlot {
	return &FnPlot{
		Scope:       scope,
		Shape:       NewFnShape(),
		IdxInParent: 0,
		Subs:        make([]*FnPlot, 0),
	}
}

type CodegenListener struct {
	*parser.BaseRippletParserListener

	scopeId   int
	curFnPlot *FnPlot
	symtab    *SymTab

	inVarDecLhs bool

	chunk *Chunk
}

func NewCodegenListener(symtab *SymTab) *CodegenListener {
	c := &CodegenListener{
		scopeId:   1,
		curFnPlot: NewFnPlot(0),
		symtab:    symtab,
		chunk:     NewChunk(),
	}
	return c
}

func (s *CodegenListener) curSymtab() *Scope {
	return s.symtab.Scopes[s.curFnPlot.Scope]
}

func (s *CodegenListener) Resolve(tree *parser.IProgramContext) *CodegenListener {
	antlr.ParseTreeWalkerDefault.Walk(s, *tree)
	return s
}

func (s *CodegenListener) Finalize() *Chunk {
	s.chunk.Fn = *s.curFnPlot.Shape
	return s.chunk
}

func (s *CodegenListener) emitOpcode(op Opcode) {
	s.curFnPlot.Shape.Instrs = append(s.curFnPlot.Shape.Instrs, int(op))
}

func (s *CodegenListener) emitIndex(i int) {
	s.curFnPlot.Shape.Instrs = append(s.curFnPlot.Shape.Instrs, i)
}

func (s *CodegenListener) EnterVarDeclareLhs(ctx *parser.VarDeclareLhsContext) {
	if ctx.Identifer() != nil {
		ctx.Identifer().EnterRule(s)
	}
}

func (s *CodegenListener) ExitVarDeclareStmt(ctx *parser.VarDeclareStmtContext) {
	s.inVarDecLhs = true
	ctx.VarDeclareLhs().EnterRule(s)
	s.inVarDecLhs = false
}

func (s *CodegenListener) EnterIdentifer(ctx *parser.IdentiferContext) {
	name := ctx.GetText()
	if s.inVarDecLhs {
		idx := s.curSymtab().LocalIdx(name)
		s.emitOpcode(STORE)
		s.emitIndex(idx)
	}
}

func (s *CodegenListener) EnterIntLiteral(ctx *parser.IntLiteralContext) {
	i, err := strconv.ParseInt(ctx.GetText(), 10, 53)
	if err != nil {
		panic(err)
	}
	ci := s.chunk.AddConstNum(float64(i))
	s.emitOpcode(CONST)
	s.emitIndex(ci)
}

func (s *CodegenListener) EnterHexLiteral(ctx *parser.HexLiteralContext) {
	i, err := strconv.ParseInt(ctx.GetText(), 16, 53)
	if err != nil {
		panic(err)
	}
	s.chunk.AddConstNum(float64(i))
}

func (s *CodegenListener) EnterRealLiteral(ctx *parser.RealLiteralContext) {
	i, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		panic(err)
	}
	s.chunk.AddConstNum(i)
}

func (s *CodegenListener) ExitAddExpr(c *parser.AddExprContext) {
	s.emitOpcode(ADD)
}
