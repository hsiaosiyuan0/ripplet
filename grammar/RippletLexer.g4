lexer grammar RippletLexer;

options {
  superClass = RippletLexerBase;
}

If: 'if';
Then: 'then';
Else: 'else';
Match: 'match';
Fn: 'fn';
Repeat: 'repeat';
Until: 'until';
Typeof: 'typeof';
Is: 'is';
Not: 'not';
IsNot: 'isnt';
And: 'and';
Or: 'or';
Break: 'break';
Ok: 'ok';
Err: 'err';
Return: 'return';
Object: 'object';
This: 'this';

LineTerminator: [\r\n\u0085\u2028\u2029] -> channel(HIDDEN);
Whitespace: [\t\u000B\u000C\u0020\u00A0]+ -> skip;
Comment: '//' ~[\r\n\u0085\u2028\u2029]* -> channel(HIDDEN);

// https://github.com/antlr/antlr4/blob/master/doc/predicates.md#semantic-predicates
RegexpStart:
  { p.IsRegexCanStart(); }? '/' -> pushMode(RegexpMode);

BraceOpen: '{';
BraceClose: '}' { l.CloseBrace(); };
BracketOpen: '[';
BracketClose: ']';
ParenOpen: '(';
ParenClose: ')';
Comma: ',';
Colon: ':';
SemiColon: ';';
Dot: '.';
Plus: '+';
Minus: '-';
Power: '**';
Multiply: '*';
Divide: '/';
Modulus: '%';
Declare: ':=';
Assign: '=';
LessThan: '<';
MoreThan: '>';
Equals: '==';
NotEquals: '!=';
LessThanEquals: '<=';
GreaterThanEquals: '>=';
LambdaConnect: '=>';
Ellipsis: '...';
Discard: '_';

NilLiteral: 'nil';

BooLiteral: 'true' | 'false';

IntLiteral: [0-9]+;

HexLiteral: '0' [xX] HexDigit+;

RealLiteral:
  [0-9]* '.' [0-9]+ ExponentPart?
  | [0-9]+ ExponentPart
  ;

Identifier: IdentifierStart IdentifierPart*;

StringOpen: '"' -> pushMode(StringMode);

mode RegexpMode;

RegexpComment: '//' ~[\r\n\u0085\u2028\u2029]*;

RegexpNewline:
  [\r\n] [\t\u000B\u000C\u0020\u00A0]+ -> channel(HIDDEN);

RegexpContent: RegexpChar+;

RexexpClose: '/' -> popMode;

mode StringMode;

StringInterpolataionStart:
  '{' { l.OpenBrace(); } -> pushMode(DEFAULT_MODE);

StringClose: '"' -> popMode;

StringQuoted: QuotedText;

fragment QuotedText: (EscapedCharacter | ~["\\{\u0085\u2028\u2029])+;

fragment HexDigit: [0-9a-fA-F];

fragment ExponentPart: [eE] [+-]? [0-9]+;

fragment IdentifierStart:
  [\p{L}]
  | [$_]
  | '\\' UnicodeEscapeSequence
  ;

fragment IdentifierPart:
  IdentifierStart
  | [\p{Mn}]
  | [\p{Nd}]
  | [\p{Pc}]
  ;

fragment EscapedCharacter:
  '\\' (
    | '{'
    | SimpleEscapeSequence
    | HexEscapeSequence
    | UnicodeEscapeSequence
  );

fragment SimpleEscapeSequence: ['"\\0abfnrtv];
fragment HexEscapeSequence:
  'x' HexDigit
  | 'x' HexDigit HexDigit
  | 'x' HexDigit HexDigit HexDigit
  | 'x' HexDigit HexDigit HexDigit HexDigit
  ;

fragment UnicodeEscapeSequence:
  'u' HexDigit HexDigit HexDigit HexDigit;

fragment RegexpChar:
  ~[\r\n\u2028\u2029\\/[]
  | RegexpEscapeSequence
  | '[' RegexpChar* ']'
  ;

fragment RegexpEscapeSequence: '\\' ~[\r\n\u2028\u2029];
