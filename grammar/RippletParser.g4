parser grammar RippletParser;

options {
  tokenVocab = RippletLexer;
}

program: statement* EOF;

statement:
  exprStmt
  | ifStmt
  | repeatStmt
	| breakStmt
  | matchStmt
  | assignStmt
  | varDeclareStmt
  | objDeclareStmt
  | fnDeclareStmt
  | blockStmt
	| returnStmt
  ;

expression:
  expression '[' expression ']'                         # SubscriptExpr
  | expression Dot Identifier                           # MemberDotExpr
  | Typeof expression                                   # TypeofExpr
  | Not expression                                      # NotExpr
  | expression arguments                                # CallExpr
  | formalParams LambdaConnect fnBody                   # FnExpr
  | <assoc=right> expression '**' expression            # PowerExpr
  | expression ('*' | Divide | '%') expression          # MulExpr
  | expression ('+' | '-') expression                   # AddExpr
	| expression ('>' | '<' | '>=' | '<=')  expression    # RelationExpr
	| expression (Is | IsNot) expression                  # EqualityExpr
  | expression (And | Or) expression                    # LogicExpr
	| Ok expression                                       # OkExpr
	| Err expression                                      # ErrExpr
	| This																								# ThisExpr
  | identifer                                           # IdentifierExpr
  | arrayLiteral                                        # ArrayExpr
  | objectLiteral                                       # ObjectExpr
  | literal                                             # LiteralExpr
  | '(' expression ')'                                  # ParenExpr
  | '(' ')'                                             # VoidExpr
  ;

breakStmt: 'break';

returnStmt: 'return' expression;

objDeclareStmt: Object '{' objProps? objMethods?  '}';

objProps: objProp (',' objProp)*;

objProp: Identifier ('=' objPropInit)?;

objPropInit: expression;

objMethods: objMethod+;

objMethod: Identifier formalParams blockStmt;

assignStmt: identifer '=' statement;

varDeclareStmt: varDeclareLhs ':=' statement;

varDeclareLhs: identifer;

exprStmt: expression;

matchStmt: Match expression '{' mathClauses '}';

mathClauses: mathClause (',' mathClause)*;

mathClause: matchClauseTest LambdaConnect statement;

matchClauseTest: matchClauseTestVal | matchClauseTestUnwrap | Discard;

matchClauseTestVal: matchClauseVal (Or matchClauseVal)*;

matchClauseTestUnwrap: (Ok | Err) '(' identifer ')';

matchClauseVal: identifer | numberLiteral;

fnDeclareStmt: Fn fnName formalParams fnBody;

fnName: identifer;

identifer: Identifier;

formalParams: '(' formalParamList? ')';

formalParamList:
  identifer (',' identifer)* (',' restParamArg)?
  | restParamArg
  ;

fnBody: '{' statement* '}';

restParamArg: Ellipsis identifer;

blockStmt: '{' statement* '}';

ifStmt: 'if' expression Then statement (Else statement)?;

repeatStmt: Repeat statement (Until expression)?;

arguments: '(' (argument (',' argument)* ','?)? ')';

argument: Ellipsis? expression;

literal:
  nilLiteral
  | boolLiteral
  | numberLiteral
  | regexpLiteral
  | stringLiteral
  ;

regexpLiteral:
  RegexpStart (RegexpContent | RegexpComment)+ RexexpClose Identifier*;

boolLiteral: BooLiteral;

nilLiteral: NilLiteral;

numberLiteral: intLiteral | hexLiteral | realLiteral;

intLiteral: IntLiteral;

hexLiteral: HexLiteral;

realLiteral: RealLiteral;

stringLiteral:
  StringOpen (stringQuoted | stringInterp)* StringClose;

stringQuoted: StringQuoted;

stringInterp: StringInterpStart stringInterpExpr BraceClose;

stringInterpExpr: expression;

arrayLiteral: '[' (expression (',' expression)*)* ']';

objectLiteral:
  '{' (prop (',' prop)*)? ','? '}';

prop:
  propName ':' expression                       # PropExpr
  | '[' expression ']' ':' expression           # ComputedPropExpr
  ;

propName: Identifier | stringLiteral | numberLiteral;

reservedWord: keyword | nilLiteral | boolLiteral;

keyword:
  If
  | Then
  | Else
  | Match
  | Fn
  | Repeat
  | Until
  | Is
  | IsNot
  | And
  | Or
  ;
