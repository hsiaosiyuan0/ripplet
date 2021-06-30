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

func main() {
	input := antlr.NewInputStream(`
	a := {
		ok b[1]
	}
	
	a = match a {
		ok(v) => v,
		_ => ()
	}
	
	b := if 1 then 2 else 3
	
	if 1 then {
	
	} else if s {
	
	} else {
	
	}
	
	repeat {
	
	} unit expr
	
	let a = () => {}
	
	fn f1() {
		c := match number {
			1 => true,
			ok(v) => {name: v},
			_ => false
		}
		c
	}
	
	object User {
		first_name = 1,
		last_name,
		age,
		bypass_data
	
		User(first_name) {
			self.first_name = first_name
		}
	
		full_name() {
			first_name + last_name
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
