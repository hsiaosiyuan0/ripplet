// Code generated from ./grammar/RippletParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // RippletParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseRippletParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRippletParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitMulExpr(ctx *MulExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitOkExpr(ctx *OkExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitSubscriptExpr(ctx *SubscriptExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitErrExpr(ctx *ErrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitLogicExpr(ctx *LogicExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitAddExpr(ctx *AddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitRelationExpr(ctx *RelationExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitVoidExpr(ctx *VoidExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitPowerExpr(ctx *PowerExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitArrayExpr(ctx *ArrayExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitEqualityExpr(ctx *EqualityExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitMemberDotExpr(ctx *MemberDotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitTypeofExpr(ctx *TypeofExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitObjectExpr(ctx *ObjectExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitIdentifierExpr(ctx *IdentifierExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitLiteralExpr(ctx *LiteralExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitThisExpr(ctx *ThisExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitFnExpr(ctx *FnExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitObjDeclareStmt(ctx *ObjDeclareStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitObjProps(ctx *ObjPropsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitObjProp(ctx *ObjPropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitObjPropInit(ctx *ObjPropInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitObjMethods(ctx *ObjMethodsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitObjMethod(ctx *ObjMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitAssignStmt(ctx *AssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitVarDeclareStmt(ctx *VarDeclareStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitVarDeclareLhs(ctx *VarDeclareLhsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitExprStmt(ctx *ExprStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitMatchStmt(ctx *MatchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitMathClauses(ctx *MathClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitMathClause(ctx *MathClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitMatchClauseTest(ctx *MatchClauseTestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitMatchClauseTestVal(ctx *MatchClauseTestValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitMatchClauseTestUnwrap(ctx *MatchClauseTestUnwrapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitMatchClauseVal(ctx *MatchClauseValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitFnDeclareStmt(ctx *FnDeclareStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitFnName(ctx *FnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitIdentifer(ctx *IdentiferContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitFormalParams(ctx *FormalParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitFormalParamList(ctx *FormalParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitFnBody(ctx *FnBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitRestParamArg(ctx *RestParamArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitBlockStmt(ctx *BlockStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitRepeatStmt(ctx *RepeatStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitRegexpLiteral(ctx *RegexpLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitBoolLiteral(ctx *BoolLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitNilLiteral(ctx *NilLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitNumberLiteral(ctx *NumberLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitHexLiteral(ctx *HexLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitRealLiteral(ctx *RealLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitStringQuoted(ctx *StringQuotedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitStringInterp(ctx *StringInterpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitStringInterpExpr(ctx *StringInterpExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitObjectLiteral(ctx *ObjectLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitPropExpr(ctx *PropExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitComputedPropExpr(ctx *ComputedPropExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitPropName(ctx *PropNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitReservedWord(ctx *ReservedWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRippletParserVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}
