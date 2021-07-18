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

type RepeatInfo struct {
	start int
	brk   int
}

type CodegenVisitor struct {
	parser.BaseRippletParserVisitor

	curFnPlot *FnPlot
	symtab    *SymTab

	inVarDecLhs       bool
	inFormalParamList bool
	inAssignLhs       bool

	repeatStk []*RepeatInfo

	chunk *Chunk
}

func NewCodegenVisitor(symtab *SymTab) *CodegenVisitor {
	c := &CodegenVisitor{
		curFnPlot: NewFnPlot(1),
		symtab:    symtab,
		repeatStk: make([]*RepeatInfo, 0),
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
	case *parser.AddExprContext:
		return v.VisitAddExpr(val)
	case *parser.MulExprContext:
		return v.VisitMulExpr(val)
	case *parser.ArrayExprContext:
		return v.VisitArrayExpr(val)
	case *parser.SubscriptExprContext:
		return v.VisitSubscriptExpr(val)
	case *parser.IfStmtContext:
		return v.VisitIfStmt(val)
	case *parser.BlockStmtContext:
		return v.VisitBlockStmt(val)
	case *parser.RelationExprContext:
		return v.VisitRelationExpr(val)
	case *parser.EqualityExprContext:
		return v.VisitEqualityExpr(val)
	case *parser.LogicExprContext:
		return v.VisitLogicExpr(val)
	case *parser.ReturnStmtContext:
		return v.VisitReturnStmt(val)
	case *parser.RepeatStmtContext:
		return v.VisitRepeatStmt(val)
	case *parser.BreakStmtContext:
		return v.VisitBreakStmt(val)
	case *parser.NegativeExprContext:
		return v.VisitNegativeExpr(val)
	default:
		panic("visit unsupported context: " + reflect.TypeOf(val).String())
	}
}

func (v *CodegenVisitor) VisitNegativeExpr(ctx *parser.NegativeExprContext) interface{} {
	v.Visit(ctx.Expression())
	v.emitOpcode(NEG)
	return nil
}

// 01: OPCODE_X
// 02: OPCODE_X
// 03: OPCODE_X
// 04: TEST
// 05: JMP_F
func (v *CodegenVisitor) VisitRepeatStmt(ctx *parser.RepeatStmtContext) interface{} {
	info := &RepeatInfo{
		start: len(v.curFnPlot.Shape.Instrs),
		brk:   -1,
	}
	v.repeatStk = append(v.repeatStk, info)

	v.Visit(ctx.Statement())
	until := ctx.Expression()
	if until != nil {
		v.Visit(until)
		v.emitOpcode(TEST)

		v.emitOpcode(JMP_F)
		v.emitInt(0)
		jmpToStartMrk := len(v.curFnPlot.Shape.Instrs)

		v.curFnPlot.Shape.Instrs[jmpToStartMrk-1] = info.start - jmpToStartMrk
	} else {
		v.emitOpcode(JMP)
		v.emitInt(0)
		jmpToStartMrk := len(v.curFnPlot.Shape.Instrs)
		v.curFnPlot.Shape.Instrs[jmpToStartMrk-1] = info.start - jmpToStartMrk
	}

	v.resolveBrk()
	return nil
}

// 01: OPCODE_X
// 02: OPCODE_X
// 03: JMP  ─────────────────┐
// 04: OPCODE_X              │
// 05: TEST                  │
// 06: JMP_F                 │
// -------------             │
// 07: OPCODE_X   ◄──────────┘
func (v *CodegenVisitor) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	if len(v.curFnPlot.Shape.Instrs) == 0 {
		panic("dangling break")
	}
	info := v.repeatStk[len(v.repeatStk)-1]
	if info.brk != -1 {
		panic("redundant break")
	}

	v.emitOpcode(JMP)
	v.emitInt(0)
	info.brk = len(v.curFnPlot.Shape.Instrs)
	return nil
}

func (v *CodegenVisitor) resolveBrk() {
	stkLen := len(v.repeatStk)
	info, repeatStk := v.repeatStk[stkLen-1], v.repeatStk[:stkLen-1]

	if info.brk != -1 {
		v.curFnPlot.Shape.Instrs[info.brk-1] = len(v.curFnPlot.Shape.Instrs) - info.brk
	}
	v.repeatStk = repeatStk
}

func (v *CodegenVisitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	v.Visit(ctx.Expression())
	v.emitOpcode(RET)
	return nil
}

func (v *CodegenVisitor) VisitEqualityExpr(ctx *parser.EqualityExprContext) interface{} {
	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))

	if ctx.Is() != nil {
		v.emitOpcode(IS)
	} else if ctx.IsNot() != nil {
		v.emitOpcode(IS_NOT)
	}
	return nil
}

func (v *CodegenVisitor) VisitLogicExpr(ctx *parser.LogicExprContext) interface{} {
	// more compact opcodes - short-circuit
	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))

	if ctx.And() != nil {
		v.emitOpcode(AND)
	} else if ctx.Or() != nil {
		v.emitOpcode(OR)
	}
	return nil
}

func (v *CodegenVisitor) VisitRelationExpr(ctx *parser.RelationExprContext) interface{} {
	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))

	if ctx.GreaterThan() != nil {
		v.emitOpcode(GT)
	} else if ctx.GreaterThanEquals() != nil {
		v.emitOpcode(GE)
	} else if ctx.LessThan() != nil {
		v.emitOpcode(LT)
	} else if ctx.LessThanEquals() != nil {
		v.emitOpcode(LE)
	}

	return nil
}

func (v *CodegenVisitor) VisitBlockStmt(ctx *parser.BlockStmtContext) interface{} {
	for _, stmt := range ctx.AllStatement() {
		v.Visit(stmt)
	}
	return nil
}

// 01: TEST
// 02: JMP_F  ─────────────────┐
// 03: OFFSET                  │
// 04: OPCODE_X                │
// 05: OPCODE_X                │
// 06: JMP   ──────────────────┼────────┐
// 07: OFFSET                  │        │
// 08: OPCODE_X  ◄─────────────┘        │
// 09: OPCODE_X                         │
// -------------                        │
// 10: OPCODE_X  ◄──────────────────────┘
func (v *CodegenVisitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	v.Visit(ctx.Expression())
	v.emitOpcode(TEST)

	v.emitOpcode(JMP_F)
	// a placeholder to resolve later
	v.emitInt(0)
	jmpOverTrueMrk := len(v.curFnPlot.Shape.Instrs)

	// emit true branch
	v.Visit(ctx.Statement(0))
	v.emitOpcode(JMP)
	v.emitInt(0)
	jmpOverFalseMrk := len(v.curFnPlot.Shape.Instrs)

	// resolve offset1
	v.curFnPlot.Shape.Instrs[jmpOverTrueMrk-1] = jmpOverFalseMrk - jmpOverTrueMrk

	// emit false branch
	elseBranch := ctx.Statement(1)
	if elseBranch != nil {
		v.Visit(elseBranch)

		// resolve offset2
		afterElseInstrLen := len(v.curFnPlot.Shape.Instrs)
		v.curFnPlot.Shape.Instrs[jmpOverFalseMrk-1] = afterElseInstrLen - jmpOverFalseMrk
	}

	return nil
}

func (v *CodegenVisitor) VisitSubscriptExpr(ctx *parser.SubscriptExprContext) interface{} {
	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))
	v.emitOpcode(SUBSCRIPT)
	return nil
}

func (v *CodegenVisitor) VisitArrayLiteral(ctx *parser.ArrayLiteralContext) interface{} {
	exprs := ctx.AllExpression()
	for _, expr := range exprs {
		v.Visit(expr)
	}
	v.emitOpcode(ARR)
	v.emitInt(len(exprs))
	return nil
}

func (v *CodegenVisitor) VisitArrayExpr(ctx *parser.ArrayExprContext) interface{} {
	v.VisitArrayLiteral(ctx.ArrayLiteral().(*parser.ArrayLiteralContext))
	return nil
}

func (v *CodegenVisitor) VisitMulExpr(ctx *parser.MulExprContext) interface{} {
	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))

	if ctx.Multiply() != nil {
		v.emitOpcode(MUL)
	} else if ctx.Divide() != nil {
		v.emitOpcode(DIV)
	} else if ctx.Modulus() != nil {
		v.emitOpcode(MOD)
	}
	return nil
}

func (v *CodegenVisitor) VisitAddExpr(ctx *parser.AddExprContext) interface{} {
	v.Visit(ctx.Expression(0))
	v.Visit(ctx.Expression(1))

	if ctx.Plus() != nil {
		v.emitOpcode(ADD)
	} else if ctx.Minus() != nil {
		v.emitOpcode(SUB)
	}
	return nil
}

// ensure the upval is in the upval-chain otherwise propagate up through FnPlots
// to setup the upval
func (v *CodegenVisitor) resolveUpChain(fn *FnPlot, name string) {
	if fn == nil {
		return
	}

	shape := fn.Shape
	if _, ok := shape.Upvals[name]; ok {
		return
	}

	var loc UpLoc
	scope := v.symtab.Scopes[fn.Scope]
	if pi := scope.ParentLocalIdx(name); pi != -1 {
		loc = UpLoc{
			Typ: UP_LOC_LOCAL,
			At:  pi,
		}
	} else {
		loc = UpLoc{
			Typ: UP_LOC_UP,
			At:  name,
		}
	}
	shape.Upvals[name] = loc

	v.resolveUpChain(fn.Parent, name)
}

func (v *CodegenVisitor) VisitIdentifierExpr(ctx *parser.IdentifierExprContext) interface{} {
	name := ctx.GetText()
	if v.curSymtab().HasLocal(name) {
		v.emitOpcode(LOAD)
		v.emitInt(v.curSymtab().LocalIdx(name))
	} else if v.curSymtab().HasBinding(name) {
		v.resolveUpChain(v.curFnPlot, name)

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
	fn := NewFnPlot(v.curFnPlot.Scope + 1)
	fn.IdxInParent = len(v.curFnPlot.Subs)
	v.curFnPlot.Subs = append(v.curFnPlot.Subs, fn)
	fn.Parent = v.curFnPlot
	v.curFnPlot = fn
}

func (v *CodegenVisitor) exitFn() {
	v.emitOpcode(RET)

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

	if ctx.RealLiteral() != nil {
		v.VisitRealLiteral(ctx.RealLiteral().(*parser.RealLiteralContext))
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

	if ctx.BoolLiteral() != nil {
		return v.VisitBoolLiteral(ctx.BoolLiteral().(*parser.BoolLiteralContext))
	}
	return nil
}

func (v *CodegenVisitor) VisitBoolLiteral(ctx *parser.BoolLiteralContext) interface{} {
	raw := ctx.GetText()
	if raw == "true" {
		v.emitOpcode(BOOL_T)
	} else if raw == "false" {
		v.emitOpcode(BOOL_F)
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
		return nil
	}

	if v.inFormalParamList {
		// TODO: varargs
		v.emitOpcode(LOAD_ARG)
		v.emitInt(v.curSymtab().LocalIdx(name))
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

	ci := v.chunk.AddConstNum(float64(i))
	v.emitOpcode(CONST)
	v.emitInt(ci)
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
