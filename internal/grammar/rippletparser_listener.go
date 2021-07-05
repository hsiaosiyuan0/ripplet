// Code generated from ./grammar/RippletParser.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // RippletParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RippletParserListener is a complete listener for a parse tree produced by RippletParser.
type RippletParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterMulExpr is called when entering the MulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterOkExpr is called when entering the OkExpr production.
	EnterOkExpr(c *OkExprContext)

	// EnterErrExpr is called when entering the ErrExpr production.
	EnterErrExpr(c *ErrExprContext)

	// EnterAddExpr is called when entering the AddExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterLogicAndExpr is called when entering the LogicAndExpr production.
	EnterLogicAndExpr(c *LogicAndExprContext)

	// EnterVoidExpr is called when entering the VoidExpr production.
	EnterVoidExpr(c *VoidExprContext)

	// EnterPowerExpr is called when entering the PowerExpr production.
	EnterPowerExpr(c *PowerExprContext)

	// EnterArrayExpr is called when entering the ArrayExpr production.
	EnterArrayExpr(c *ArrayExprContext)

	// EnterEqualityIsNotExpr is called when entering the EqualityIsNotExpr production.
	EnterEqualityIsNotExpr(c *EqualityIsNotExprContext)

	// EnterMemterIdxExpr is called when entering the MemterIdxExpr production.
	EnterMemterIdxExpr(c *MemterIdxExprContext)

	// EnterMemberDotExpr is called when entering the MemberDotExpr production.
	EnterMemberDotExpr(c *MemberDotExprContext)

	// EnterTypeofExpr is called when entering the TypeofExpr production.
	EnterTypeofExpr(c *TypeofExprContext)

	// EnterObjectExpr is called when entering the ObjectExpr production.
	EnterObjectExpr(c *ObjectExprContext)

	// EnterIdentifierExpr is called when entering the IdentifierExpr production.
	EnterIdentifierExpr(c *IdentifierExprContext)

	// EnterLogicOrExpr is called when entering the LogicOrExpr production.
	EnterLogicOrExpr(c *LogicOrExprContext)

	// EnterLiteralExpr is called when entering the LiteralExpr production.
	EnterLiteralExpr(c *LiteralExprContext)

	// EnterCallExpr is called when entering the CallExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterEqualityIsExpr is called when entering the EqualityIsExpr production.
	EnterEqualityIsExpr(c *EqualityIsExprContext)

	// EnterThisExpr is called when entering the ThisExpr production.
	EnterThisExpr(c *ThisExprContext)

	// EnterFnExpr is called when entering the FnExpr production.
	EnterFnExpr(c *FnExprContext)

	// EnterObjDeclareStmt is called when entering the objDeclareStmt production.
	EnterObjDeclareStmt(c *ObjDeclareStmtContext)

	// EnterObjProps is called when entering the objProps production.
	EnterObjProps(c *ObjPropsContext)

	// EnterObjProp is called when entering the objProp production.
	EnterObjProp(c *ObjPropContext)

	// EnterObjPropInit is called when entering the objPropInit production.
	EnterObjPropInit(c *ObjPropInitContext)

	// EnterObjMethods is called when entering the objMethods production.
	EnterObjMethods(c *ObjMethodsContext)

	// EnterObjMethod is called when entering the objMethod production.
	EnterObjMethod(c *ObjMethodContext)

	// EnterAssignStmt is called when entering the assignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterVarDeclareStmt is called when entering the varDeclareStmt production.
	EnterVarDeclareStmt(c *VarDeclareStmtContext)

	// EnterVarDeclareLhs is called when entering the varDeclareLhs production.
	EnterVarDeclareLhs(c *VarDeclareLhsContext)

	// EnterExprStmt is called when entering the exprStmt production.
	EnterExprStmt(c *ExprStmtContext)

	// EnterMatchStmt is called when entering the matchStmt production.
	EnterMatchStmt(c *MatchStmtContext)

	// EnterMathClauses is called when entering the mathClauses production.
	EnterMathClauses(c *MathClausesContext)

	// EnterMathClause is called when entering the mathClause production.
	EnterMathClause(c *MathClauseContext)

	// EnterMatchClauseTest is called when entering the matchClauseTest production.
	EnterMatchClauseTest(c *MatchClauseTestContext)

	// EnterMatchClauseTestVal is called when entering the matchClauseTestVal production.
	EnterMatchClauseTestVal(c *MatchClauseTestValContext)

	// EnterMatchClauseTestUnwrap is called when entering the matchClauseTestUnwrap production.
	EnterMatchClauseTestUnwrap(c *MatchClauseTestUnwrapContext)

	// EnterMatchClauseVal is called when entering the matchClauseVal production.
	EnterMatchClauseVal(c *MatchClauseValContext)

	// EnterFnDeclareStmt is called when entering the fnDeclareStmt production.
	EnterFnDeclareStmt(c *FnDeclareStmtContext)

	// EnterFnName is called when entering the fnName production.
	EnterFnName(c *FnNameContext)

	// EnterIdentifer is called when entering the identifer production.
	EnterIdentifer(c *IdentiferContext)

	// EnterFormalParams is called when entering the formalParams production.
	EnterFormalParams(c *FormalParamsContext)

	// EnterFormalParamList is called when entering the formalParamList production.
	EnterFormalParamList(c *FormalParamListContext)

	// EnterFnBody is called when entering the fnBody production.
	EnterFnBody(c *FnBodyContext)

	// EnterRestParamArg is called when entering the restParamArg production.
	EnterRestParamArg(c *RestParamArgContext)

	// EnterBlockStmt is called when entering the blockStmt production.
	EnterBlockStmt(c *BlockStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterRepeatStmt is called when entering the repeatStmt production.
	EnterRepeatStmt(c *RepeatStmtContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterRegexpLiteral is called when entering the regexpLiteral production.
	EnterRegexpLiteral(c *RegexpLiteralContext)

	// EnterBoolLiteral is called when entering the boolLiteral production.
	EnterBoolLiteral(c *BoolLiteralContext)

	// EnterNilLiteral is called when entering the nilLiteral production.
	EnterNilLiteral(c *NilLiteralContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterIntLiteral is called when entering the intLiteral production.
	EnterIntLiteral(c *IntLiteralContext)

	// EnterHexLiteral is called when entering the hexLiteral production.
	EnterHexLiteral(c *HexLiteralContext)

	// EnterRealLiteral is called when entering the realLiteral production.
	EnterRealLiteral(c *RealLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterStringQuoted is called when entering the stringQuoted production.
	EnterStringQuoted(c *StringQuotedContext)

	// EnterStringInterp is called when entering the stringInterp production.
	EnterStringInterp(c *StringInterpContext)

	// EnterStringInterpExpr is called when entering the stringInterpExpr production.
	EnterStringInterpExpr(c *StringInterpExprContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterObjectLiteral is called when entering the objectLiteral production.
	EnterObjectLiteral(c *ObjectLiteralContext)

	// EnterPropExpr is called when entering the PropExpr production.
	EnterPropExpr(c *PropExprContext)

	// EnterComputedPropExpr is called when entering the ComputedPropExpr production.
	EnterComputedPropExpr(c *ComputedPropExprContext)

	// EnterPropName is called when entering the propName production.
	EnterPropName(c *PropNameContext)

	// EnterReservedWord is called when entering the reservedWord production.
	EnterReservedWord(c *ReservedWordContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitMulExpr is called when exiting the MulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitOkExpr is called when exiting the OkExpr production.
	ExitOkExpr(c *OkExprContext)

	// ExitErrExpr is called when exiting the ErrExpr production.
	ExitErrExpr(c *ErrExprContext)

	// ExitAddExpr is called when exiting the AddExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitLogicAndExpr is called when exiting the LogicAndExpr production.
	ExitLogicAndExpr(c *LogicAndExprContext)

	// ExitVoidExpr is called when exiting the VoidExpr production.
	ExitVoidExpr(c *VoidExprContext)

	// ExitPowerExpr is called when exiting the PowerExpr production.
	ExitPowerExpr(c *PowerExprContext)

	// ExitArrayExpr is called when exiting the ArrayExpr production.
	ExitArrayExpr(c *ArrayExprContext)

	// ExitEqualityIsNotExpr is called when exiting the EqualityIsNotExpr production.
	ExitEqualityIsNotExpr(c *EqualityIsNotExprContext)

	// ExitMemterIdxExpr is called when exiting the MemterIdxExpr production.
	ExitMemterIdxExpr(c *MemterIdxExprContext)

	// ExitMemberDotExpr is called when exiting the MemberDotExpr production.
	ExitMemberDotExpr(c *MemberDotExprContext)

	// ExitTypeofExpr is called when exiting the TypeofExpr production.
	ExitTypeofExpr(c *TypeofExprContext)

	// ExitObjectExpr is called when exiting the ObjectExpr production.
	ExitObjectExpr(c *ObjectExprContext)

	// ExitIdentifierExpr is called when exiting the IdentifierExpr production.
	ExitIdentifierExpr(c *IdentifierExprContext)

	// ExitLogicOrExpr is called when exiting the LogicOrExpr production.
	ExitLogicOrExpr(c *LogicOrExprContext)

	// ExitLiteralExpr is called when exiting the LiteralExpr production.
	ExitLiteralExpr(c *LiteralExprContext)

	// ExitCallExpr is called when exiting the CallExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitEqualityIsExpr is called when exiting the EqualityIsExpr production.
	ExitEqualityIsExpr(c *EqualityIsExprContext)

	// ExitThisExpr is called when exiting the ThisExpr production.
	ExitThisExpr(c *ThisExprContext)

	// ExitFnExpr is called when exiting the FnExpr production.
	ExitFnExpr(c *FnExprContext)

	// ExitObjDeclareStmt is called when exiting the objDeclareStmt production.
	ExitObjDeclareStmt(c *ObjDeclareStmtContext)

	// ExitObjProps is called when exiting the objProps production.
	ExitObjProps(c *ObjPropsContext)

	// ExitObjProp is called when exiting the objProp production.
	ExitObjProp(c *ObjPropContext)

	// ExitObjPropInit is called when exiting the objPropInit production.
	ExitObjPropInit(c *ObjPropInitContext)

	// ExitObjMethods is called when exiting the objMethods production.
	ExitObjMethods(c *ObjMethodsContext)

	// ExitObjMethod is called when exiting the objMethod production.
	ExitObjMethod(c *ObjMethodContext)

	// ExitAssignStmt is called when exiting the assignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitVarDeclareStmt is called when exiting the varDeclareStmt production.
	ExitVarDeclareStmt(c *VarDeclareStmtContext)

	// ExitVarDeclareLhs is called when exiting the varDeclareLhs production.
	ExitVarDeclareLhs(c *VarDeclareLhsContext)

	// ExitExprStmt is called when exiting the exprStmt production.
	ExitExprStmt(c *ExprStmtContext)

	// ExitMatchStmt is called when exiting the matchStmt production.
	ExitMatchStmt(c *MatchStmtContext)

	// ExitMathClauses is called when exiting the mathClauses production.
	ExitMathClauses(c *MathClausesContext)

	// ExitMathClause is called when exiting the mathClause production.
	ExitMathClause(c *MathClauseContext)

	// ExitMatchClauseTest is called when exiting the matchClauseTest production.
	ExitMatchClauseTest(c *MatchClauseTestContext)

	// ExitMatchClauseTestVal is called when exiting the matchClauseTestVal production.
	ExitMatchClauseTestVal(c *MatchClauseTestValContext)

	// ExitMatchClauseTestUnwrap is called when exiting the matchClauseTestUnwrap production.
	ExitMatchClauseTestUnwrap(c *MatchClauseTestUnwrapContext)

	// ExitMatchClauseVal is called when exiting the matchClauseVal production.
	ExitMatchClauseVal(c *MatchClauseValContext)

	// ExitFnDeclareStmt is called when exiting the fnDeclareStmt production.
	ExitFnDeclareStmt(c *FnDeclareStmtContext)

	// ExitFnName is called when exiting the fnName production.
	ExitFnName(c *FnNameContext)

	// ExitIdentifer is called when exiting the identifer production.
	ExitIdentifer(c *IdentiferContext)

	// ExitFormalParams is called when exiting the formalParams production.
	ExitFormalParams(c *FormalParamsContext)

	// ExitFormalParamList is called when exiting the formalParamList production.
	ExitFormalParamList(c *FormalParamListContext)

	// ExitFnBody is called when exiting the fnBody production.
	ExitFnBody(c *FnBodyContext)

	// ExitRestParamArg is called when exiting the restParamArg production.
	ExitRestParamArg(c *RestParamArgContext)

	// ExitBlockStmt is called when exiting the blockStmt production.
	ExitBlockStmt(c *BlockStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitRepeatStmt is called when exiting the repeatStmt production.
	ExitRepeatStmt(c *RepeatStmtContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitRegexpLiteral is called when exiting the regexpLiteral production.
	ExitRegexpLiteral(c *RegexpLiteralContext)

	// ExitBoolLiteral is called when exiting the boolLiteral production.
	ExitBoolLiteral(c *BoolLiteralContext)

	// ExitNilLiteral is called when exiting the nilLiteral production.
	ExitNilLiteral(c *NilLiteralContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitIntLiteral is called when exiting the intLiteral production.
	ExitIntLiteral(c *IntLiteralContext)

	// ExitHexLiteral is called when exiting the hexLiteral production.
	ExitHexLiteral(c *HexLiteralContext)

	// ExitRealLiteral is called when exiting the realLiteral production.
	ExitRealLiteral(c *RealLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitStringQuoted is called when exiting the stringQuoted production.
	ExitStringQuoted(c *StringQuotedContext)

	// ExitStringInterp is called when exiting the stringInterp production.
	ExitStringInterp(c *StringInterpContext)

	// ExitStringInterpExpr is called when exiting the stringInterpExpr production.
	ExitStringInterpExpr(c *StringInterpExprContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitObjectLiteral is called when exiting the objectLiteral production.
	ExitObjectLiteral(c *ObjectLiteralContext)

	// ExitPropExpr is called when exiting the PropExpr production.
	ExitPropExpr(c *PropExprContext)

	// ExitComputedPropExpr is called when exiting the ComputedPropExpr production.
	ExitComputedPropExpr(c *ComputedPropExprContext)

	// ExitPropName is called when exiting the propName production.
	ExitPropName(c *PropNameContext)

	// ExitReservedWord is called when exiting the reservedWord production.
	ExitReservedWord(c *ReservedWordContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)
}
