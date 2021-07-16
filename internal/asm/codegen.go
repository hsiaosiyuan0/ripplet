package asm

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/hsiaosiyuan0/ripplet/internal/grammar"
)

// FnPlot consist of the infomation collected and required
// for the function instructions generation
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

func (f *FnPlot) ToShape() *FnShape {
	s := f.Shape
	for _, ss := range f.Subs {
		s.Subs = append(s.Subs, ss.ToShape())
	}
	return s
}

type CodegenVisitor struct {
	parser.BaseRippletParserVisitor

	scopeId   uint
	curFnPlot *FnPlot
	symtab    *SymTab

	inVarDecLhs       bool
	inFormalParamList bool
	inAssignLhs       bool

	chunk *Chunk
}

func NewCodegenVisitor(symtab *SymTab) *CodegenVisitor {
	c := &CodegenVisitor{
		scopeId:   1,
		curFnPlot: NewFnPlot(1),
		symtab:    symtab,
		chunk:     NewChunk(),
	}
	return c
}

func (v *CodegenVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.ProgramContext:
		return v.VisitProgram(val)
	case *parser.VarDeclareStmtContext:
		return v.VisitVarDeclareStmt(val)
	case *parser.VarDeclareLhsContext:
		return v.VisitVarDeclareLhs(val)
	case *parser.StatementContext:
		return v.VisitStatement(val)
	case *parser.ExprStmtContext:
		return v.VisitExprStmt(val)
	case *parser.LiteralExprContext:
		return v.VisitLiteralExpr(val)
	case *parser.LiteralContext:
		return v.VisitLiteral(val)
	case *parser.NumberLiteralContext:
		return v.VisitNumberLiteral(val)
	case *parser.FnExprContext:
		return v.VisitFnExpr(val)
	case *parser.IdentifierExprContext:
		return v.VisitIdentifierExpr(val)
	case *parser.AssignStmtContext:
		return v.VisitAssignStmt(val)
	case *parser.CallExprContext:
		return v.VisitCallExpr(val)
	case *parser.ArgumentContext:
		return v.VisitArgument(val)
	default:
		panic("visit unsupported context: " + reflect.TypeOf(val).String())
	}
}

func (v *CodegenVisitor) VisitIdentifierExpr(ctx *parser.IdentifierExprContext) interface{} {
	name := ctx.GetText()
	if v.curSymtab().HasLocal(name) {
		v.emitOpcode(LOAD)
		v.emitInt(v.curSymtab().LocalIdx(name))
	} else if v.curSymtab().HasBinding(name) {
		v.emitOpcode(LOAD_UP)
		v.emitString(name)
	} else if v.symtab.HasExternal(name) {
		v.emitOpcode(LOAD_EXT)
		v.emitString(name)
	}
	return nil
}

func (v *CodegenVisitor) VisitFormalParamList(ctx *parser.FormalParamListContext) interface{} {
	v.inFormalParamList = true
	for _, id := range ctx.AllIdentifer() {
		v.VisitIdentifer(id.(*parser.IdentiferContext))
	}
	v.inFormalParamList = false
	return nil
}

func (v *CodegenVisitor) VisitFormalParams(ctx *parser.FormalParamsContext) interface{} {
	if ctx.FormalParamList() != nil {
		v.VisitFormalParamList(ctx.FormalParamList().(*parser.FormalParamListContext))
	}
	return nil
}

func (v *CodegenVisitor) enterFn() {
	fn := NewFnPlot(v.scopeId + 1)
	fn.IdxInParent = len(v.curFnPlot.Subs)
	v.curFnPlot.Subs = append(v.curFnPlot.Subs, fn)
	fn.Parent = v.curFnPlot
	v.curFnPlot = fn
}

func (v *CodegenVisitor) exitFn() {
	v.curFnPlot.Shape.LocalCnt = v.curSymtab().LocalCnt()
	v.curFnPlot = v.curFnPlot.Parent
}

func (v *CodegenVisitor) VisitFnBody(ctx *parser.FnBodyContext) interface{} {
	for _, stmt := range ctx.AllStatement() {
		v.Visit(stmt)
	}
	return nil
}

func (v *CodegenVisitor) VisitFnExpr(ctx *parser.FnExprContext) interface{} {
	v.enterFn()
	idx := v.curFnPlot.IdxInParent

	v.VisitFormalParams(ctx.FormalParams().(*parser.FormalParamsContext))
	v.VisitFnBody(ctx.FnBody().(*parser.FnBodyContext))

	v.exitFn()
	v.emitOpcode(CLOSURE)
	v.emitInt(idx)
	return nil
}

func (v *CodegenVisitor) VisitNumberLiteral(ctx *parser.NumberLiteralContext) interface{} {
	if ctx.IntLiteral() != nil {
		v.VisitIntLiteral(ctx.IntLiteral().(*parser.IntLiteralContext))
	}

	if ctx.HexLiteral() != nil {
		v.VisitHexLiteral(ctx.HexLiteral().(*parser.HexLiteralContext))
	}
	return nil
}

func (v *CodegenVisitor) VisitStringInterpExpr(ctx *parser.StringInterpExprContext) interface{} {
	return v.Visit(ctx.Expression())
}

func (v *CodegenVisitor) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if val, ok := child.(*parser.StringQuotedContext); ok {
			ci := v.chunk.AddConstStr(val.GetText())
			v.emitOpcode(CONST)
			v.emitInt(ci)
			continue
		}
		if val, ok := child.(*parser.StringInterpContext); ok {
			v.VisitStringInterpExpr(val.StringInterpExpr().(*parser.StringInterpExprContext))
			continue
		}
	}
	v.emitOpcode(CONCAT)
	v.emitInt(len(ctx.AllStringInterp()) + len(ctx.AllStringQuoted()))
	return nil
}

func (v *CodegenVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.NumberLiteral() != nil {
		return v.VisitNumberLiteral(ctx.NumberLiteral().(*parser.NumberLiteralContext))
	}

	if ctx.StringLiteral() != nil {
		return v.VisitStringLiteral(ctx.StringLiteral().(*parser.StringLiteralContext))
	}
	return nil
}

func (v *CodegenVisitor) VisitLiteralExpr(ctx *parser.LiteralExprContext) interface{} {
	return v.VisitLiteral(ctx.Literal().(*parser.LiteralContext))
}

func (v *CodegenVisitor) VisitExprStmt(ctx *parser.ExprStmtContext) interface{} {
	return v.Visit(ctx.Expression())
}

func (v *CodegenVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	return nil
}

func (v *CodegenVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	switch val := node.(type) {
	case *parser.ProgramContext:
		for _, child := range val.AllStatement() {
			v.Visit(child)
		}
		return nil
	default:
		panic("visit children on unsupported context: " + reflect.TypeOf(val).String())
	}
}

func (v *CodegenVisitor) VisitIdentifer(ctx *parser.IdentiferContext) interface{} {
	name := ctx.GetText()
	if v.inVarDecLhs || (v.inAssignLhs && v.curSymtab().HasLocal(name)) {
		idx := v.curSymtab().LocalIdx(name)
		v.emitOpcode(STORE)
		v.emitInt(idx)
		return nil
	}

	if v.inAssignLhs {
		v.emitOpcode(STORE_UP)
		v.emitString(name)
	}
	return nil
}

func (v *CodegenVisitor) VisitVarDeclareLhs(ctx *parser.VarDeclareLhsContext) interface{} {
	v.inVarDecLhs = true
	v.VisitIdentifer(ctx.Identifer().(*parser.IdentiferContext))
	v.inVarDecLhs = false
	return nil
}

func (v *CodegenVisitor) VisitVarDeclareStmt(ctx *parser.VarDeclareStmtContext) interface{} {
	v.Visit(ctx.Statement())
	v.VisitVarDeclareLhs(ctx.VarDeclareLhs().(*parser.VarDeclareLhsContext))
	return nil
}

func (v *CodegenVisitor) VisitAssignStmt(ctx *parser.AssignStmtContext) interface{} {
	v.Visit(ctx.Statement())
	v.inAssignLhs = true
	v.VisitIdentifer(ctx.Identifer().(*parser.IdentiferContext))
	v.inAssignLhs = false
	return nil
}

func (v *CodegenVisitor) VisitArguments(ctx *parser.ArgumentsContext) int {
	for _, arg := range ctx.AllArgument() {
		v.VisitArgument(arg.(*parser.ArgumentContext))
	}
	return len(ctx.AllArgument())
}

func (v *CodegenVisitor) VisitArgument(ctx *parser.ArgumentContext) interface{} {
	v.Visit(ctx.Expression())
	return nil
}

func (v *CodegenVisitor) VisitCallExpr(ctx *parser.CallExprContext) interface{} {
	v.Visit(ctx.Expression())
	argsLen := v.VisitArguments(ctx.Arguments().(*parser.ArgumentsContext))
	v.emitOpcode(CALL)
	v.emitInt(argsLen + 1)
	return nil
}

func (v *CodegenVisitor) VisitIntLiteral(ctx *parser.IntLiteralContext) interface{} {
	i, err := strconv.ParseInt(ctx.GetText(), 10, 53)
	if err != nil {
		panic(err)
	}
	ci := v.chunk.AddConstNum(float64(i))
	v.emitOpcode(CONST)
	v.emitInt(ci)
	return nil
}

func (v *CodegenVisitor) VisitHexLiteral(ctx *parser.HexLiteralContext) interface{} {
	raw := ctx.GetText()
	raw = strings.Replace(raw, "0x", "", -1)

	i, err := strconv.ParseInt(raw, 16, 53)
	if err != nil {
		panic(err)
	}

	ci := v.chunk.AddConstNum(float64(i))
	v.emitOpcode(CONST)
	v.emitInt(ci)
	return nil
}

func (v *CodegenVisitor) VisitRealLiteral(ctx *parser.RealLiteralContext) interface{} {
	i, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		panic(err)
	}
	v.chunk.AddConstNum(i)
	return nil
}

func (v *CodegenVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	v.VisitChildren(ctx)
	v.curFnPlot.Shape.LocalCnt = v.curSymtab().LocalCnt()
	return nil
}

func (v *CodegenVisitor) curSymtab() *Scope {
	return v.symtab.Scopes[v.curFnPlot.Scope]
}

func (v *CodegenVisitor) Finalize() *Chunk {
	v.chunk.Fn = *v.curFnPlot.ToShape()
	return v.chunk
}

func (v *CodegenVisitor) emitOpcode(op Opcode) {
	v.curFnPlot.Shape.Instrs = append(v.curFnPlot.Shape.Instrs, int(op))
}

func (v *CodegenVisitor) emitInt(i int) {
	v.curFnPlot.Shape.Instrs = append(v.curFnPlot.Shape.Instrs, i)
}

func (v *CodegenVisitor) emitString(str string) {
	i := v.chunk.AddConstStr(str)
	v.emitInt(i)
}
