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
	l.braces = append(l.braces, 1)
}

func (l *RippletLexerBase) CloseBrace() {
	if len(l.braces) != 0 {
		braces := l.braces[:]
		l.PopMode()
		l.braces = braces[:len(braces)-1]
	}
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
