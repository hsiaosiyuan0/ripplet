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
  | Identifier arguments                                # CallExpr
  | '(' formalParamList ')' LambdaConnect fnBody        # FnExpr
  | <assoc=right> expression '**' expression            # PowerExpr
  | expression ('*' | Divide | '%') expression          # MulExpr
  | expression ('+' | '-')                              # AddExpr
  | expression And expression                           # LogicAndExpr
  | expression Or expression                            # LogicOrExpr
	| Ok expression                                       # OkExpr
	| Err expression                                      # ErrExpr
	| This																								# ThisExpr
  | Identifier                                          # IdentifierExpr
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

objMethod: Identifier '(' formalParamList ')' blockStmt;

assignStmt: Identifier '=' statement;

varDeclareStmt: Identifier ':=' statement;

exprStmt: expression;

matchStmt: Match expression '{' mathClauses '}';

mathClauses: mathClause (',' mathClause)*;

mathClause: matchClauseTest LambdaConnect statement;

matchClauseTest: matchClauseTestVal | matchClauseTestRes | Discard;

matchClauseTestVal: matchClauseVal (Or matchClauseVal)*;

matchClauseTestRes: (Ok | Err) '(' Identifier ')';

matchClauseVal: Identifier | numberLiteral;

fnDeclareStmt: Fn Identifier '(' formalParamList? ')' fnBody;

formalParamList:
  Identifier (',' Identifier)* (',' restParamArg)?
  | restParamArg
  ;

fnBody: '{' statement* '}';

restParamArg: Ellipsis Identifier;

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

numberLiteral: IntLiteral | HexLiteral | RealLiteral;

stringLiteral:
  StringOpen (
    StringQuoted
    | (StringInterpolataionStart expression BraceClose)
  )* StringClose;

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
