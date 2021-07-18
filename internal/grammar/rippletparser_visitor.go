// Code generated from ./grammar/RippletParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // RippletParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by RippletParser.
type RippletParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RippletParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by RippletParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by RippletParser#MulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by RippletParser#OkExpr.
	VisitOkExpr(ctx *OkExprContext) interface{}

	// Visit a parse tree produced by RippletParser#SubscriptExpr.
	VisitSubscriptExpr(ctx *SubscriptExprContext) interface{}

	// Visit a parse tree produced by RippletParser#ErrExpr.
	VisitErrExpr(ctx *ErrExprContext) interface{}

	// Visit a parse tree produced by RippletParser#LogicExpr.
	VisitLogicExpr(ctx *LogicExprContext) interface{}

	// Visit a parse tree produced by RippletParser#AddExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by RippletParser#RelationExpr.
	VisitRelationExpr(ctx *RelationExprContext) interface{}

	// Visit a parse tree produced by RippletParser#VoidExpr.
	VisitVoidExpr(ctx *VoidExprContext) interface{}

	// Visit a parse tree produced by RippletParser#PowerExpr.
	VisitPowerExpr(ctx *PowerExprContext) interface{}

	// Visit a parse tree produced by RippletParser#NegativeExpr.
	VisitNegativeExpr(ctx *NegativeExprContext) interface{}

	// Visit a parse tree produced by RippletParser#ArrayExpr.
	VisitArrayExpr(ctx *ArrayExprContext) interface{}

	// Visit a parse tree produced by RippletParser#EqualityExpr.
	VisitEqualityExpr(ctx *EqualityExprContext) interface{}

	// Visit a parse tree produced by RippletParser#MemberDotExpr.
	VisitMemberDotExpr(ctx *MemberDotExprContext) interface{}

	// Visit a parse tree produced by RippletParser#TypeofExpr.
	VisitTypeofExpr(ctx *TypeofExprContext) interface{}

	// Visit a parse tree produced by RippletParser#ObjectExpr.
	VisitObjectExpr(ctx *ObjectExprContext) interface{}

	// Visit a parse tree produced by RippletParser#IdentifierExpr.
	VisitIdentifierExpr(ctx *IdentifierExprContext) interface{}

	// Visit a parse tree produced by RippletParser#LiteralExpr.
	VisitLiteralExpr(ctx *LiteralExprContext) interface{}

	// Visit a parse tree produced by RippletParser#CallExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by RippletParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by RippletParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by RippletParser#ThisExpr.
	VisitThisExpr(ctx *ThisExprContext) interface{}

	// Visit a parse tree produced by RippletParser#FnExpr.
	VisitFnExpr(ctx *FnExprContext) interface{}

	// Visit a parse tree produced by RippletParser#breakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by RippletParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by RippletParser#objDeclareStmt.
	VisitObjDeclareStmt(ctx *ObjDeclareStmtContext) interface{}

	// Visit a parse tree produced by RippletParser#objProps.
	VisitObjProps(ctx *ObjPropsContext) interface{}

	// Visit a parse tree produced by RippletParser#objProp.
	VisitObjProp(ctx *ObjPropContext) interface{}

	// Visit a parse tree produced by RippletParser#objPropInit.
	VisitObjPropInit(ctx *ObjPropInitContext) interface{}

	// Visit a parse tree produced by RippletParser#objMethods.
	VisitObjMethods(ctx *ObjMethodsContext) interface{}

	// Visit a parse tree produced by RippletParser#objMethod.
	VisitObjMethod(ctx *ObjMethodContext) interface{}

	// Visit a parse tree produced by RippletParser#assignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by RippletParser#varDeclareStmt.
	VisitVarDeclareStmt(ctx *VarDeclareStmtContext) interface{}

	// Visit a parse tree produced by RippletParser#varDeclareLhs.
	VisitVarDeclareLhs(ctx *VarDeclareLhsContext) interface{}

	// Visit a parse tree produced by RippletParser#exprStmt.
	VisitExprStmt(ctx *ExprStmtContext) interface{}

	// Visit a parse tree produced by RippletParser#matchStmt.
	VisitMatchStmt(ctx *MatchStmtContext) interface{}

	// Visit a parse tree produced by RippletParser#mathClauses.
	VisitMathClauses(ctx *MathClausesContext) interface{}

	// Visit a parse tree produced by RippletParser#mathClause.
	VisitMathClause(ctx *MathClauseContext) interface{}

	// Visit a parse tree produced by RippletParser#matchClauseTest.
	VisitMatchClauseTest(ctx *MatchClauseTestContext) interface{}

	// Visit a parse tree produced by RippletParser#matchClauseTestVal.
	VisitMatchClauseTestVal(ctx *MatchClauseTestValContext) interface{}

	// Visit a parse tree produced by RippletParser#matchClauseTestUnwrap.
	VisitMatchClauseTestUnwrap(ctx *MatchClauseTestUnwrapContext) interface{}

	// Visit a parse tree produced by RippletParser#matchClauseVal.
	VisitMatchClauseVal(ctx *MatchClauseValContext) interface{}

	// Visit a parse tree produced by RippletParser#fnDeclareStmt.
	VisitFnDeclareStmt(ctx *FnDeclareStmtContext) interface{}

	// Visit a parse tree produced by RippletParser#fnName.
	VisitFnName(ctx *FnNameContext) interface{}

	// Visit a parse tree produced by RippletParser#identifer.
	VisitIdentifer(ctx *IdentiferContext) interface{}

	// Visit a parse tree produced by RippletParser#formalParams.
	VisitFormalParams(ctx *FormalParamsContext) interface{}

	// Visit a parse tree produced by RippletParser#formalParamList.
	VisitFormalParamList(ctx *FormalParamListContext) interface{}

	// Visit a parse tree produced by RippletParser#fnBody.
	VisitFnBody(ctx *FnBodyContext) interface{}

	// Visit a parse tree produced by RippletParser#restParamArg.
	VisitRestParamArg(ctx *RestParamArgContext) interface{}

	// Visit a parse tree produced by RippletParser#blockStmt.
	VisitBlockStmt(ctx *BlockStmtContext) interface{}

	// Visit a parse tree produced by RippletParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by RippletParser#repeatStmt.
	VisitRepeatStmt(ctx *RepeatStmtContext) interface{}

	// Visit a parse tree produced by RippletParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by RippletParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by RippletParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by RippletParser#regexpLiteral.
	VisitRegexpLiteral(ctx *RegexpLiteralContext) interface{}

	// Visit a parse tree produced by RippletParser#boolLiteral.
	VisitBoolLiteral(ctx *BoolLiteralContext) interface{}

	// Visit a parse tree produced by RippletParser#nilLiteral.
	VisitNilLiteral(ctx *NilLiteralContext) interface{}

	// Visit a parse tree produced by RippletParser#numberLiteral.
	VisitNumberLiteral(ctx *NumberLiteralContext) interface{}

	// Visit a parse tree produced by RippletParser#intLiteral.
	VisitIntLiteral(ctx *IntLiteralContext) interface{}

	// Visit a parse tree produced by RippletParser#hexLiteral.
	VisitHexLiteral(ctx *HexLiteralContext) interface{}

	// Visit a parse tree produced by RippletParser#realLiteral.
	VisitRealLiteral(ctx *RealLiteralContext) interface{}

	// Visit a parse tree produced by RippletParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by RippletParser#stringQuoted.
	VisitStringQuoted(ctx *StringQuotedContext) interface{}

	// Visit a parse tree produced by RippletParser#stringInterp.
	VisitStringInterp(ctx *StringInterpContext) interface{}

	// Visit a parse tree produced by RippletParser#stringInterpExpr.
	VisitStringInterpExpr(ctx *StringInterpExprContext) interface{}

	// Visit a parse tree produced by RippletParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by RippletParser#objectLiteral.
	VisitObjectLiteral(ctx *ObjectLiteralContext) interface{}

	// Visit a parse tree produced by RippletParser#PropExpr.
	VisitPropExpr(ctx *PropExprContext) interface{}

	// Visit a parse tree produced by RippletParser#ComputedPropExpr.
	VisitComputedPropExpr(ctx *ComputedPropExprContext) interface{}

	// Visit a parse tree produced by RippletParser#propName.
	VisitPropName(ctx *PropNameContext) interface{}

	// Visit a parse tree produced by RippletParser#reservedWord.
	VisitReservedWord(ctx *ReservedWordContext) interface{}

	// Visit a parse tree produced by RippletParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}
}
