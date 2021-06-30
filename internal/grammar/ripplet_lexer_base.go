package parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type RippletLexerBase struct {
	*antlr.BaseLexer

	braces    []int
	lastToken antlr.Token
}

func (l *RippletLexerBase) NextToken() antlr.Token {
	next := l.BaseLexer.NextToken()
	if next.GetChannel() == antlr.TokenDefaultChannel {
		l.lastToken = next
	}
	return next
}

func (l *RippletLexerBase) OpenBrace() {
	if len(l.braces) != 0 {
		braces := l.braces[:]
		cnt := len(braces)
		last, braces := braces[cnt-1], braces[:cnt-1]
		if last == 0 {
			l.PopMode()
		} else {
			braces = append(braces, last)
		}
		l.braces = braces
	}
}

func (l *RippletLexerBase) CloseBrace() {
	l.braces = append(l.braces, 1)
}

func (l *RippletLexerBase) IsRegexCanStart() bool {
	if l.lastToken == nil {
		return true
	}
	switch l.lastToken.GetTokenType() {
	case RippletLexerIdentifier, RippletLexerNilLiteral, RippletLexerBooLiteral, RippletLexerIntLiteral, RippletLexerHexLiteral, RippletLexerRealLiteral, RippletLexerStringClose:
		return false
	default:
		return true
	}
}
