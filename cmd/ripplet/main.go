package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/hsiaosiyuan0/ripplet/internal/grammar"
)

type TreeShapeListener struct {
	*parser.BaseRippletParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

// func (s *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	fmt.Println(ctx.GetText())
// }

func (s *TreeShapeListener) EnterMemterIdxExpr(ctx *parser.MemterIdxExprContext) {
	// fmt.Println(ctx.GetText())
}

func (s *TreeShapeListener) EnterVarDeclareStmt(ctx *parser.VarDeclareStmtContext) {
	// fmt.Println(ctx.GetText())
}

func (s *TreeShapeListener) EnterIfStmt(ctx *parser.IfStmtContext) {
	fmt.Println(ctx.GetText())
}

func (s *TreeShapeListener) EnterIdentifer(ctx *parser.IdentiferContext) {
	fmt.Println("id: ", ctx.GetText())
}

func main() {
	input := antlr.NewInputStream(`
	a := 1
	fn b() {
		c := 1
	
		d := () => {
			e := 1
		}
	}	
	`)
	lexer := parser.NewRippletLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewRippletParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Program()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

	// p := parser.NewJSONParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// p.BuildParseTrees = true
}
