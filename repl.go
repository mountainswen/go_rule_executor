package main

import (
	"antlr_test/V2parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/c-bata/go-prompt"
	"os"
)

const PROMPT = ">>"
var env = MakeEnv()
func Handle(input string) {
	//scanner := bufio.NewScanner(in)
		if len(input) == 0 {
			return
		}

		if input == "exit" {
			os.Exit(0)
		}

		var v = NewVisitor()
		v.SetEnv(env)

	//env := evaluator.NewEnvironment()
	//for {
		is := antlr.NewInputStream(input)
		// Setup the input
		// Create the Lexer
		lexer := parser.NewV2ParserLexer(is)
		tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		// Create the Parser
		p := parser.NewV2ParserParser(tokens)

		//Start is rule name of Calc.g4
		p.SourceFile().Accept(v) //从头结点开始遍历
	//}
}

func Start(){
	p := prompt.New(Handle, completer, prompt.OptionPrefix(">>>> "))
	p.Run()
}

func completer(d prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
