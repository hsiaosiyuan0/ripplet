parser grammar RippletParser;

options {
  tokenVocab = RippletLexer;
}

program: statement* EOF;

statement:
  exprStmt
  | ifStmt
  | repeatStmt
  | matchStmt
  | assignStmt
  | varDeclareStmt
  | objDeclareStmt
  | fnDeclareStmt
  | blockStmt
  ;

expression:
  expression '[' expression ']'                         # MemterIdxExpr
  | expression Dot Identifier                           # MemberDotExpr
  | Typeof expression                                   # TypeofExpr
  | expression Is expression                            # EqualityIsExpr
  | expression IsNot expression                         # EqualityIsNotExpr
  | Not expression                                      # NotExpr
  | identifer arguments                                 # CallExpr
  | formalParams LambdaConnect fnBody                   # FnExpr
  | <assoc=right> expression '**' expression            # PowerExpr
  | expression ('*' | Divide | '%') expression          # MulExpr
  | expression ('+' | '-') expression                   # AddExpr
  | expression And expression                           # LogicAndExpr
  | expression Or expression                            # LogicOrExpr
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

ifStmt: 'if' expression Then? statement (Else statement)?;

repeatStmt: Repeat statement Until expression;

arguments: '(' (argument (',' argument)* ','?)? ')';

argument: Ellipsis? (expression | Identifier);

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
  StringOpen (StringQuoted | stringInterp)* StringClose;

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
