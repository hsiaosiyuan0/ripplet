package asm

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/hsiaosiyuan0/ripplet/internal/grammar"
)

type Scope struct {
	// an auto-increment number which is generated according
	// the depth-first walk over the entire AST
	Id uint

	Parent *Scope
	Subs   []*Scope

	// formal parameters will be added in both `Params` and `Bindings`
	Params   map[string]bool
	Bindings map[string]bool
}

func NewScope() *Scope {
	scope := &Scope{
		Id:       0,
		Subs:     make([]*Scope, 0),
		Params:   make(map[string]bool),
		Bindings: make(map[string]bool),
	}
	return scope
}

func (s *Scope) AddBinding(name string) {
	s.Bindings[name] = true
}

func (s *Scope) HasLocal(name string) bool {
	return s.Bindings[name]
}

func (s *Scope) HasBinding(name string) bool {
	scope := s
	for scope != nil {
		if scope.Bindings[name] {
			return true
		}
		scope = scope.Parent
	}
	return false
}

func (s *Scope) AddParam(name string) {
	s.Params[name] = true
}

func (s *Scope) HasParam(name string) bool {
	return s.Params[name]
}

type SymTab struct {
	Scopes map[uint]*Scope
	Root   *Scope
	Cur    *Scope
}

func NewSymTab() *SymTab {
	scope := NewScope()
	symtab := &SymTab{
		Scopes: make(map[uint]*Scope),
		Root:   scope,
		Cur:    scope,
	}
	symtab.Scopes[scope.Id] = scope
	return symtab
}

func (s *SymTab) EnterScope() {
	scope := NewScope()
	scope.Id = s.Cur.Id + 1
	s.Scopes[scope.Id] = scope

	scope.Parent = s.Cur
	s.Cur.Subs = append(s.Cur.Subs, scope)

	s.Cur = scope
}

func (s *SymTab) LeaveScope() {
	s.Cur = s.Cur.Parent
}

type SymTabListener struct {
	*parser.BaseRippletParserListener

	symtab *SymTab

	inVarDeclareLhs   bool
	inFnName          bool
	inFormalParamList bool
}

func NewSymTabListener() *SymTabListener {
	s := new(SymTabListener)
	s.symtab = NewSymTab()
	return s
}

func (s *SymTabListener) Resolve(tree *parser.IProgramContext) (symtab *SymTab, err error) {
	defer func() {
		if r := recover(); r != nil {
			if v, ok := r.(error); ok {
				err = v
			} else {
				panic(r)
			}
		}
	}()

	antlr.ParseTreeWalkerDefault.Walk(s, *tree)
	return s.symtab, nil
}

func (s *SymTabListener) EnterBlockStmt(ctx *parser.BlockStmtContext) {
	s.symtab.EnterScope()
}

func (s *SymTabListener) ExitBlockStmt(ctx *parser.BlockStmtContext) {
	s.symtab.LeaveScope()
}

func (s *SymTabListener) EnterFnName(ctx *parser.FnNameContext) {
	s.inFnName = true
}

func (s *SymTabListener) ExitFnName(ctx *parser.FnNameContext) {
	s.inFnName = false
}

func (s *SymTabListener) EnterFormalParams(ctx *parser.FormalParamsContext) {
	s.symtab.EnterScope()
	s.inFormalParamList = true
}

func (s *SymTabListener) ExitFormalParams(ctx *parser.FormalParamsContext) {
	s.inFormalParamList = false
}

func (s *SymTabListener) ExitFnBody(ctx *parser.FnBodyContext) {
	s.symtab.LeaveScope()
}

func (s *SymTabListener) EnterMathClauses(ctx *parser.MathClausesContext) {
	s.symtab.EnterScope()
}

func (s *SymTabListener) ExitMathClauses(ctx *parser.MathClausesContext) {
	s.symtab.LeaveScope()
}

func (s *SymTabListener) EnterVarDeclareLhs(ctx *parser.VarDeclareLhsContext) {
	s.inVarDeclareLhs = true
}

func (s *SymTabListener) ExitVarDeclareLhs(ctx *parser.VarDeclareLhsContext) {
	s.inVarDeclareLhs = false
}

func (s *SymTabListener) EnterIdentifer(ctx *parser.IdentiferContext) {
	name := ctx.GetText()

	if s.inVarDeclareLhs || s.inFnName {
		s.symtab.Cur.AddBinding(name)
	} else if s.inFormalParamList {
		s.symtab.Cur.AddParam(name)
		s.symtab.Cur.AddBinding(name)
	}

	if !s.symtab.Cur.HasBinding(name) {
		sym := ctx.GetToken(parser.RippletLexerIdentifier, 0).GetSymbol()
		sourceName := sym.GetInputStream().GetSourceName()
		line := sym.GetLine()
		col := sym.GetColumn()
		err := fmt.Errorf("%s undefined: %s at %d:%d", sourceName, name, line, col)
		panic(err)
	}
}
