// example1.go
package main

import (
	"antlr_test/parserV1"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"antlr_test/V2parser"
)


var ifStmt = `
  aa = "mike" + " wen"
   d += 1.3 * 2.534 +1 
   d = true || false
   if (d) {
     f = 12+ 123.1
   }else {
    f = 1 + 2
	}

   ss = fmt.sprintf("hahah")
`

var methodcall = `
 ss = fmt.sprintf("hello:%s","mikewen")
`
func main() {
	var a = []rune("12313")
	fmt.Println(a[0])
	// Setup the input
	calc(methodcall)
}

func V1(){
	// Setup the input
	is := antlr.NewInputStream("1 + 2 * 3+4*12")

	// Create the Lexer
	lexer := parserV1.NewCalcLexer(is)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	listen := NewListenImpl()
	p := parserV1.NewCalcParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(listen, p.Start())

	listen.PrintStack()
}

func calc(input string) int {

	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewV2ParserLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewV2ParserParser(tokens)

	v := NewVisitor()

	//Start is rule name of Calc.g4
	p.SourceFile().Accept(v) //从头结点开始遍历
	v.PrintEnv()
	return 0
}