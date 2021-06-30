// Code generated from ./grammar/RippletParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // RippletParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRippletParserListener is a complete listener for a parse tree produced by RippletParser.
type BaseRippletParserListener struct{}

var _ RippletParserListener = &BaseRippletParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRippletParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRippletParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRippletParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRippletParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseRippletParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseRippletParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseRippletParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseRippletParserListener) ExitStatement(ctx *StatementContext) {}

// EnterMulExpr is called when production MulExpr is entered.
func (s *BaseRippletParserListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production MulExpr is exited.
func (s *BaseRippletParserListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterOkExpr is called when production OkExpr is entered.
func (s *BaseRippletParserListener) EnterOkExpr(ctx *OkExprContext) {}

// ExitOkExpr is called when production OkExpr is exited.
func (s *BaseRippletParserListener) ExitOkExpr(ctx *OkExprContext) {}

// EnterErrExpr is called when production ErrExpr is entered.
func (s *BaseRippletParserListener) EnterErrExpr(ctx *ErrExprContext) {}

// ExitErrExpr is called when production ErrExpr is exited.
func (s *BaseRippletParserListener) ExitErrExpr(ctx *ErrExprContext) {}

// EnterAddExpr is called when production AddExpr is entered.
func (s *BaseRippletParserListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production AddExpr is exited.
func (s *BaseRippletParserListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterLogicAndExpr is called when production LogicAndExpr is entered.
func (s *BaseRippletParserListener) EnterLogicAndExpr(ctx *LogicAndExprContext) {}

// ExitLogicAndExpr is called when production LogicAndExpr is exited.
func (s *BaseRippletParserListener) ExitLogicAndExpr(ctx *LogicAndExprContext) {}

// EnterVoidExpr is called when production VoidExpr is entered.
func (s *BaseRippletParserListener) EnterVoidExpr(ctx *VoidExprContext) {}

// ExitVoidExpr is called when production VoidExpr is exited.
func (s *BaseRippletParserListener) ExitVoidExpr(ctx *VoidExprContext) {}

// EnterPowerExpr is called when production PowerExpr is entered.
func (s *BaseRippletParserListener) EnterPowerExpr(ctx *PowerExprContext) {}

// ExitPowerExpr is called when production PowerExpr is exited.
func (s *BaseRippletParserListener) ExitPowerExpr(ctx *PowerExprContext) {}

// EnterArrayExpr is called when production ArrayExpr is entered.
func (s *BaseRippletParserListener) EnterArrayExpr(ctx *ArrayExprContext) {}

// ExitArrayExpr is called when production ArrayExpr is exited.
func (s *BaseRippletParserListener) ExitArrayExpr(ctx *ArrayExprContext) {}

// EnterEqualityIsNotExpr is called when production EqualityIsNotExpr is entered.
func (s *BaseRippletParserListener) EnterEqualityIsNotExpr(ctx *EqualityIsNotExprContext) {}

// ExitEqualityIsNotExpr is called when production EqualityIsNotExpr is exited.
func (s *BaseRippletParserListener) ExitEqualityIsNotExpr(ctx *EqualityIsNotExprContext) {}

// EnterMemterIdxExpr is called when production MemterIdxExpr is entered.
func (s *BaseRippletParserListener) EnterMemterIdxExpr(ctx *MemterIdxExprContext) {}

// ExitMemterIdxExpr is called when production MemterIdxExpr is exited.
func (s *BaseRippletParserListener) ExitMemterIdxExpr(ctx *MemterIdxExprContext) {}

// EnterMemberDotExpr is called when production MemberDotExpr is entered.
func (s *BaseRippletParserListener) EnterMemberDotExpr(ctx *MemberDotExprContext) {}

// ExitMemberDotExpr is called when production MemberDotExpr is exited.
func (s *BaseRippletParserListener) ExitMemberDotExpr(ctx *MemberDotExprContext) {}

// EnterTypeofExpr is called when production TypeofExpr is entered.
func (s *BaseRippletParserListener) EnterTypeofExpr(ctx *TypeofExprContext) {}

// ExitTypeofExpr is called when production TypeofExpr is exited.
func (s *BaseRippletParserListener) ExitTypeofExpr(ctx *TypeofExprContext) {}

// EnterObjectExpr is called when production ObjectExpr is entered.
func (s *BaseRippletParserListener) EnterObjectExpr(ctx *ObjectExprContext) {}

// ExitObjectExpr is called when production ObjectExpr is exited.
func (s *BaseRippletParserListener) ExitObjectExpr(ctx *ObjectExprContext) {}

// EnterIdentifierExpr is called when production IdentifierExpr is entered.
func (s *BaseRippletParserListener) EnterIdentifierExpr(ctx *IdentifierExprContext) {}

// ExitIdentifierExpr is called when production IdentifierExpr is exited.
func (s *BaseRippletParserListener) ExitIdentifierExpr(ctx *IdentifierExprContext) {}

// EnterLogicOrExpr is called when production LogicOrExpr is entered.
func (s *BaseRippletParserListener) EnterLogicOrExpr(ctx *LogicOrExprContext) {}

// ExitLogicOrExpr is called when production LogicOrExpr is exited.
func (s *BaseRippletParserListener) ExitLogicOrExpr(ctx *LogicOrExprContext) {}

// EnterLiteralExpr is called when production LiteralExpr is entered.
func (s *BaseRippletParserListener) EnterLiteralExpr(ctx *LiteralExprContext) {}

// ExitLiteralExpr is called when production LiteralExpr is exited.
func (s *BaseRippletParserListener) ExitLiteralExpr(ctx *LiteralExprContext) {}

// EnterCallExpr is called when production CallExpr is entered.
func (s *BaseRippletParserListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production CallExpr is exited.
func (s *BaseRippletParserListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseRippletParserListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseRippletParserListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseRippletParserListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseRippletParserListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterEqualityIsExpr is called when production EqualityIsExpr is entered.
func (s *BaseRippletParserListener) EnterEqualityIsExpr(ctx *EqualityIsExprContext) {}

// ExitEqualityIsExpr is called when production EqualityIsExpr is exited.
func (s *BaseRippletParserListener) ExitEqualityIsExpr(ctx *EqualityIsExprContext) {}

// EnterThisExpr is called when production ThisExpr is entered.
func (s *BaseRippletParserListener) EnterThisExpr(ctx *ThisExprContext) {}

// ExitThisExpr is called when production ThisExpr is exited.
func (s *BaseRippletParserListener) ExitThisExpr(ctx *ThisExprContext) {}

// EnterFnExpr is called when production FnExpr is entered.
func (s *BaseRippletParserListener) EnterFnExpr(ctx *FnExprContext) {}

// ExitFnExpr is called when production FnExpr is exited.
func (s *BaseRippletParserListener) ExitFnExpr(ctx *FnExprContext) {}

// EnterObjDeclareStmt is called when production objDeclareStmt is entered.
func (s *BaseRippletParserListener) EnterObjDeclareStmt(ctx *ObjDeclareStmtContext) {}

// ExitObjDeclareStmt is called when production objDeclareStmt is exited.
func (s *BaseRippletParserListener) ExitObjDeclareStmt(ctx *ObjDeclareStmtContext) {}

// EnterObjProps is called when production objProps is entered.
func (s *BaseRippletParserListener) EnterObjProps(ctx *ObjPropsContext) {}

// ExitObjProps is called when production objProps is exited.
func (s *BaseRippletParserListener) ExitObjProps(ctx *ObjPropsContext) {}

// EnterObjProp is called when production objProp is entered.
func (s *BaseRippletParserListener) EnterObjProp(ctx *ObjPropContext) {}

// ExitObjProp is called when production objProp is exited.
func (s *BaseRippletParserListener) ExitObjProp(ctx *ObjPropContext) {}

// EnterObjPropInit is called when production objPropInit is entered.
func (s *BaseRippletParserListener) EnterObjPropInit(ctx *ObjPropInitContext) {}

// ExitObjPropInit is called when production objPropInit is exited.
func (s *BaseRippletParserListener) ExitObjPropInit(ctx *ObjPropInitContext) {}

// EnterObjMethods is called when production objMethods is entered.
func (s *BaseRippletParserListener) EnterObjMethods(ctx *ObjMethodsContext) {}

// ExitObjMethods is called when production objMethods is exited.
func (s *BaseRippletParserListener) ExitObjMethods(ctx *ObjMethodsContext) {}

// EnterObjMethod is called when production objMethod is entered.
func (s *BaseRippletParserListener) EnterObjMethod(ctx *ObjMethodContext) {}

// ExitObjMethod is called when production objMethod is exited.
func (s *BaseRippletParserListener) ExitObjMethod(ctx *ObjMethodContext) {}

// EnterAssignStmt is called when production assignStmt is entered.
func (s *BaseRippletParserListener) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production assignStmt is exited.
func (s *BaseRippletParserListener) ExitAssignStmt(ctx *AssignStmtContext) {}

// EnterVarDeclareStmt is called when production varDeclareStmt is entered.
func (s *BaseRippletParserListener) EnterVarDeclareStmt(ctx *VarDeclareStmtContext) {}

// ExitVarDeclareStmt is called when production varDeclareStmt is exited.
func (s *BaseRippletParserListener) ExitVarDeclareStmt(ctx *VarDeclareStmtContext) {}

// EnterVarDeclareLhs is called when production varDeclareLhs is entered.
func (s *BaseRippletParserListener) EnterVarDeclareLhs(ctx *VarDeclareLhsContext) {}

// ExitVarDeclareLhs is called when production varDeclareLhs is exited.
func (s *BaseRippletParserListener) ExitVarDeclareLhs(ctx *VarDeclareLhsContext) {}

// EnterExprStmt is called when production exprStmt is entered.
func (s *BaseRippletParserListener) EnterExprStmt(ctx *ExprStmtContext) {}

// ExitExprStmt is called when production exprStmt is exited.
func (s *BaseRippletParserListener) ExitExprStmt(ctx *ExprStmtContext) {}

// EnterMatchStmt is called when production matchStmt is entered.
func (s *BaseRippletParserListener) EnterMatchStmt(ctx *MatchStmtContext) {}

// ExitMatchStmt is called when production matchStmt is exited.
func (s *BaseRippletParserListener) ExitMatchStmt(ctx *MatchStmtContext) {}

// EnterMathClauses is called when production mathClauses is entered.
func (s *BaseRippletParserListener) EnterMathClauses(ctx *MathClausesContext) {}

// ExitMathClauses is called when production mathClauses is exited.
func (s *BaseRippletParserListener) ExitMathClauses(ctx *MathClausesContext) {}

// EnterMathClause is called when production mathClause is entered.
func (s *BaseRippletParserListener) EnterMathClause(ctx *MathClauseContext) {}

// ExitMathClause is called when production mathClause is exited.
func (s *BaseRippletParserListener) ExitMathClause(ctx *MathClauseContext) {}

// EnterMatchClauseTest is called when production matchClauseTest is entered.
func (s *BaseRippletParserListener) EnterMatchClauseTest(ctx *MatchClauseTestContext) {}

// ExitMatchClauseTest is called when production matchClauseTest is exited.
func (s *BaseRippletParserListener) ExitMatchClauseTest(ctx *MatchClauseTestContext) {}

// EnterMatchClauseTestVal is called when production matchClauseTestVal is entered.
func (s *BaseRippletParserListener) EnterMatchClauseTestVal(ctx *MatchClauseTestValContext) {}

// ExitMatchClauseTestVal is called when production matchClauseTestVal is exited.
func (s *BaseRippletParserListener) ExitMatchClauseTestVal(ctx *MatchClauseTestValContext) {}

// EnterMatchClauseTestUnwrap is called when production matchClauseTestUnwrap is entered.
func (s *BaseRippletParserListener) EnterMatchClauseTestUnwrap(ctx *MatchClauseTestUnwrapContext) {}

// ExitMatchClauseTestUnwrap is called when production matchClauseTestUnwrap is exited.
func (s *BaseRippletParserListener) ExitMatchClauseTestUnwrap(ctx *MatchClauseTestUnwrapContext) {}

// EnterMatchClauseVal is called when production matchClauseVal is entered.
func (s *BaseRippletParserListener) EnterMatchClauseVal(ctx *MatchClauseValContext) {}

// ExitMatchClauseVal is called when production matchClauseVal is exited.
func (s *BaseRippletParserListener) ExitMatchClauseVal(ctx *MatchClauseValContext) {}

// EnterFnDeclareStmt is called when production fnDeclareStmt is entered.
func (s *BaseRippletParserListener) EnterFnDeclareStmt(ctx *FnDeclareStmtContext) {}

// ExitFnDeclareStmt is called when production fnDeclareStmt is exited.
func (s *BaseRippletParserListener) ExitFnDeclareStmt(ctx *FnDeclareStmtContext) {}

// EnterFnName is called when production fnName is entered.
func (s *BaseRippletParserListener) EnterFnName(ctx *FnNameContext) {}

// ExitFnName is called when production fnName is exited.
func (s *BaseRippletParserListener) ExitFnName(ctx *FnNameContext) {}

// EnterIdentifer is called when production identifer is entered.
func (s *BaseRippletParserListener) EnterIdentifer(ctx *IdentiferContext) {}

// ExitIdentifer is called when production identifer is exited.
func (s *BaseRippletParserListener) ExitIdentifer(ctx *IdentiferContext) {}

// EnterFormalParams is called when production formalParams is entered.
func (s *BaseRippletParserListener) EnterFormalParams(ctx *FormalParamsContext) {}

// ExitFormalParams is called when production formalParams is exited.
func (s *BaseRippletParserListener) ExitFormalParams(ctx *FormalParamsContext) {}

// EnterFormalParamList is called when production formalParamList is entered.
func (s *BaseRippletParserListener) EnterFormalParamList(ctx *FormalParamListContext) {}

// ExitFormalParamList is called when production formalParamList is exited.
func (s *BaseRippletParserListener) ExitFormalParamList(ctx *FormalParamListContext) {}

// EnterFnBody is called when production fnBody is entered.
func (s *BaseRippletParserListener) EnterFnBody(ctx *FnBodyContext) {}

// ExitFnBody is called when production fnBody is exited.
func (s *BaseRippletParserListener) ExitFnBody(ctx *FnBodyContext) {}

// EnterRestParamArg is called when production restParamArg is entered.
func (s *BaseRippletParserListener) EnterRestParamArg(ctx *RestParamArgContext) {}

// ExitRestParamArg is called when production restParamArg is exited.
func (s *BaseRippletParserListener) ExitRestParamArg(ctx *RestParamArgContext) {}

// EnterBlockStmt is called when production blockStmt is entered.
func (s *BaseRippletParserListener) EnterBlockStmt(ctx *BlockStmtContext) {}

// ExitBlockStmt is called when production blockStmt is exited.
func (s *BaseRippletParserListener) ExitBlockStmt(ctx *BlockStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseRippletParserListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseRippletParserListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterRepeatStmt is called when production repeatStmt is entered.
func (s *BaseRippletParserListener) EnterRepeatStmt(ctx *RepeatStmtContext) {}

// ExitRepeatStmt is called when production repeatStmt is exited.
func (s *BaseRippletParserListener) ExitRepeatStmt(ctx *RepeatStmtContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseRippletParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseRippletParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseRippletParserListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseRippletParserListener) ExitArgument(ctx *ArgumentContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseRippletParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseRippletParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterRegexpLiteral is called when production regexpLiteral is entered.
func (s *BaseRippletParserListener) EnterRegexpLiteral(ctx *RegexpLiteralContext) {}

// ExitRegexpLiteral is called when production regexpLiteral is exited.
func (s *BaseRippletParserListener) ExitRegexpLiteral(ctx *RegexpLiteralContext) {}

// EnterBoolLiteral is called when production boolLiteral is entered.
func (s *BaseRippletParserListener) EnterBoolLiteral(ctx *BoolLiteralContext) {}

// ExitBoolLiteral is called when production boolLiteral is exited.
func (s *BaseRippletParserListener) ExitBoolLiteral(ctx *BoolLiteralContext) {}

// EnterNilLiteral is called when production nilLiteral is entered.
func (s *BaseRippletParserListener) EnterNilLiteral(ctx *NilLiteralContext) {}

// ExitNilLiteral is called when production nilLiteral is exited.
func (s *BaseRippletParserListener) ExitNilLiteral(ctx *NilLiteralContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseRippletParserListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseRippletParserListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseRippletParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseRippletParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseRippletParserListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseRippletParserListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterObjectLiteral is called when production objectLiteral is entered.
func (s *BaseRippletParserListener) EnterObjectLiteral(ctx *ObjectLiteralContext) {}

// ExitObjectLiteral is called when production objectLiteral is exited.
func (s *BaseRippletParserListener) ExitObjectLiteral(ctx *ObjectLiteralContext) {}

// EnterPropExpr is called when production PropExpr is entered.
func (s *BaseRippletParserListener) EnterPropExpr(ctx *PropExprContext) {}

// ExitPropExpr is called when production PropExpr is exited.
func (s *BaseRippletParserListener) ExitPropExpr(ctx *PropExprContext) {}

// EnterComputedPropExpr is called when production ComputedPropExpr is entered.
func (s *BaseRippletParserListener) EnterComputedPropExpr(ctx *ComputedPropExprContext) {}

// ExitComputedPropExpr is called when production ComputedPropExpr is exited.
func (s *BaseRippletParserListener) ExitComputedPropExpr(ctx *ComputedPropExprContext) {}

// EnterPropName is called when production propName is entered.
func (s *BaseRippletParserListener) EnterPropName(ctx *PropNameContext) {}

// ExitPropName is called when production propName is exited.
func (s *BaseRippletParserListener) ExitPropName(ctx *PropNameContext) {}

// EnterReservedWord is called when production reservedWord is entered.
func (s *BaseRippletParserListener) EnterReservedWord(ctx *ReservedWordContext) {}

// ExitReservedWord is called when production reservedWord is exited.
func (s *BaseRippletParserListener) ExitReservedWord(ctx *ReservedWordContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseRippletParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseRippletParserListener) ExitKeyword(ctx *KeywordContext) {}
