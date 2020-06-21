// Generated from V2Parser.g4 by ANTLR 4.7.

package parser // V2Parser

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 74, 231,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 6, 3, 66, 10, 3,
	13, 3, 14, 3, 67, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 74, 10, 4, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 85, 10, 7, 3, 7, 3, 7,
	5, 7, 89, 10, 7, 3, 8, 3, 8, 3, 8, 7, 8, 94, 10, 8, 12, 8, 14, 8, 97, 11,
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 5, 12,
	108, 10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 5, 14, 121, 10, 14, 3, 15, 3, 15, 5, 15, 125, 10, 15,
	3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 7, 16, 132, 10, 16, 12, 16, 14, 16,
	135, 11, 16, 3, 17, 3, 17, 3, 17, 5, 17, 140, 10, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 7, 17, 157, 10, 17, 12, 17, 14, 17, 160, 11, 17, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 170, 10, 18, 7, 18,
	172, 10, 18, 12, 18, 14, 18, 175, 11, 18, 3, 19, 3, 19, 3, 19, 5, 19, 180,
	10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 189, 10,
	20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	5, 23, 201, 10, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 215, 10, 27, 3, 28, 3, 28, 3, 29,
	3, 29, 5, 29, 221, 10, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 5,
	30, 229, 10, 30, 3, 30, 2, 4, 32, 34, 31, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
	56, 58, 2, 9, 4, 2, 49, 54, 56, 60, 4, 2, 50, 54, 59, 60, 4, 2, 49, 49,
	56, 58, 3, 2, 43, 48, 3, 2, 55, 58, 3, 2, 68, 69, 4, 2, 62, 64, 66, 67,
	2, 235, 2, 60, 3, 2, 2, 2, 4, 65, 3, 2, 2, 2, 6, 73, 3, 2, 2, 2, 8, 75,
	3, 2, 2, 2, 10, 77, 3, 2, 2, 2, 12, 80, 3, 2, 2, 2, 14, 90, 3, 2, 2, 2,
	16, 98, 3, 2, 2, 2, 18, 100, 3, 2, 2, 2, 20, 102, 3, 2, 2, 2, 22, 107,
	3, 2, 2, 2, 24, 111, 3, 2, 2, 2, 26, 113, 3, 2, 2, 2, 28, 122, 3, 2, 2,
	2, 30, 128, 3, 2, 2, 2, 32, 139, 3, 2, 2, 2, 34, 161, 3, 2, 2, 2, 36, 179,
	3, 2, 2, 2, 38, 188, 3, 2, 2, 2, 40, 190, 3, 2, 2, 2, 42, 192, 3, 2, 2,
	2, 44, 200, 3, 2, 2, 2, 46, 202, 3, 2, 2, 2, 48, 204, 3, 2, 2, 2, 50, 206,
	3, 2, 2, 2, 52, 210, 3, 2, 2, 2, 54, 216, 3, 2, 2, 2, 56, 218, 3, 2, 2,
	2, 58, 228, 3, 2, 2, 2, 60, 61, 5, 4, 3, 2, 61, 3, 3, 2, 2, 2, 62, 63,
	5, 6, 4, 2, 63, 64, 5, 58, 30, 2, 64, 66, 3, 2, 2, 2, 65, 62, 3, 2, 2,
	2, 66, 67, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 5, 3,
	2, 2, 2, 69, 74, 5, 8, 5, 2, 70, 74, 5, 20, 11, 2, 71, 74, 5, 24, 13, 2,
	72, 74, 5, 26, 14, 2, 73, 69, 3, 2, 2, 2, 73, 70, 3, 2, 2, 2, 73, 71, 3,
	2, 2, 2, 73, 72, 3, 2, 2, 2, 74, 7, 3, 2, 2, 2, 75, 76, 5, 10, 6, 2, 76,
	9, 3, 2, 2, 2, 77, 78, 7, 22, 2, 2, 78, 79, 5, 12, 7, 2, 79, 11, 3, 2,
	2, 2, 80, 88, 5, 14, 8, 2, 81, 84, 5, 16, 9, 2, 82, 83, 7, 32, 2, 2, 83,
	85, 5, 30, 16, 2, 84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 89, 3, 2,
	2, 2, 86, 87, 7, 32, 2, 2, 87, 89, 5, 30, 16, 2, 88, 81, 3, 2, 2, 2, 88,
	86, 3, 2, 2, 2, 89, 13, 3, 2, 2, 2, 90, 95, 7, 25, 2, 2, 91, 92, 7, 33,
	2, 2, 92, 94, 7, 25, 2, 2, 93, 91, 3, 2, 2, 2, 94, 97, 3, 2, 2, 2, 95,
	93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 15, 3, 2, 2, 2, 97, 95, 3, 2, 2,
	2, 98, 99, 5, 18, 10, 2, 99, 17, 3, 2, 2, 2, 100, 101, 7, 25, 2, 2, 101,
	19, 3, 2, 2, 2, 102, 103, 5, 30, 16, 2, 103, 104, 5, 22, 12, 2, 104, 105,
	5, 30, 16, 2, 105, 21, 3, 2, 2, 2, 106, 108, 9, 2, 2, 2, 107, 106, 3, 2,
	2, 2, 107, 108, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 110, 7, 32, 2, 2,
	110, 23, 3, 2, 2, 2, 111, 112, 5, 32, 17, 2, 112, 25, 3, 2, 2, 2, 113,
	114, 7, 15, 2, 2, 114, 115, 7, 26, 2, 2, 115, 116, 5, 32, 17, 2, 116, 117,
	7, 27, 2, 2, 117, 120, 5, 28, 15, 2, 118, 119, 7, 9, 2, 2, 119, 121, 5,
	28, 15, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 27, 3, 2, 2,
	2, 122, 124, 7, 28, 2, 2, 123, 125, 5, 4, 3, 2, 124, 123, 3, 2, 2, 2, 124,
	125, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 127, 7, 29, 2, 2, 127, 29,
	3, 2, 2, 2, 128, 133, 5, 32, 17, 2, 129, 130, 7, 33, 2, 2, 130, 132, 5,
	32, 17, 2, 131, 129, 3, 2, 2, 2, 132, 135, 3, 2, 2, 2, 133, 131, 3, 2,
	2, 2, 133, 134, 3, 2, 2, 2, 134, 31, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2,
	136, 137, 8, 17, 1, 2, 137, 140, 5, 34, 18, 2, 138, 140, 5, 36, 19, 2,
	139, 136, 3, 2, 2, 2, 139, 138, 3, 2, 2, 2, 140, 158, 3, 2, 2, 2, 141,
	142, 12, 7, 2, 2, 142, 143, 9, 3, 2, 2, 143, 157, 5, 32, 17, 8, 144, 145,
	12, 6, 2, 2, 145, 146, 9, 4, 2, 2, 146, 157, 5, 32, 17, 7, 147, 148, 12,
	5, 2, 2, 148, 149, 9, 5, 2, 2, 149, 157, 5, 32, 17, 6, 150, 151, 12, 4,
	2, 2, 151, 152, 7, 42, 2, 2, 152, 157, 5, 32, 17, 5, 153, 154, 12, 3, 2,
	2, 154, 155, 7, 41, 2, 2, 155, 157, 5, 32, 17, 4, 156, 141, 3, 2, 2, 2,
	156, 144, 3, 2, 2, 2, 156, 147, 3, 2, 2, 2, 156, 150, 3, 2, 2, 2, 156,
	153, 3, 2, 2, 2, 157, 160, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 158, 159,
	3, 2, 2, 2, 159, 33, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 161, 162, 8, 18,
	1, 2, 162, 163, 5, 38, 20, 2, 163, 173, 3, 2, 2, 2, 164, 169, 12, 3, 2,
	2, 165, 166, 7, 36, 2, 2, 166, 170, 7, 25, 2, 2, 167, 170, 5, 50, 26, 2,
	168, 170, 5, 56, 29, 2, 169, 165, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169,
	168, 3, 2, 2, 2, 170, 172, 3, 2, 2, 2, 171, 164, 3, 2, 2, 2, 172, 175,
	3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 35, 3, 2,
	2, 2, 175, 173, 3, 2, 2, 2, 176, 180, 5, 34, 18, 2, 177, 178, 9, 6, 2,
	2, 178, 180, 5, 32, 17, 2, 179, 176, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2,
	180, 37, 3, 2, 2, 2, 181, 189, 5, 40, 21, 2, 182, 189, 5, 48, 25, 2, 183,
	189, 5, 52, 27, 2, 184, 185, 7, 26, 2, 2, 185, 186, 5, 32, 17, 2, 186,
	187, 7, 27, 2, 2, 187, 189, 3, 2, 2, 2, 188, 181, 3, 2, 2, 2, 188, 182,
	3, 2, 2, 2, 188, 183, 3, 2, 2, 2, 188, 184, 3, 2, 2, 2, 189, 39, 3, 2,
	2, 2, 190, 191, 5, 44, 23, 2, 191, 41, 3, 2, 2, 2, 192, 193, 9, 7, 2, 2,
	193, 43, 3, 2, 2, 2, 194, 201, 7, 24, 2, 2, 195, 201, 5, 46, 24, 2, 196,
	201, 5, 42, 22, 2, 197, 201, 7, 65, 2, 2, 198, 201, 7, 66, 2, 2, 199, 201,
	7, 67, 2, 2, 200, 194, 3, 2, 2, 2, 200, 195, 3, 2, 2, 2, 200, 196, 3, 2,
	2, 2, 200, 197, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 200, 199, 3, 2, 2, 2,
	201, 45, 3, 2, 2, 2, 202, 203, 9, 8, 2, 2, 203, 47, 3, 2, 2, 2, 204, 205,
	7, 25, 2, 2, 205, 49, 3, 2, 2, 2, 206, 207, 7, 30, 2, 2, 207, 208, 5, 32,
	17, 2, 208, 209, 7, 31, 2, 2, 209, 51, 3, 2, 2, 2, 210, 211, 5, 54, 28,
	2, 211, 212, 7, 36, 2, 2, 212, 214, 7, 25, 2, 2, 213, 215, 5, 56, 29, 2,
	214, 213, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 53, 3, 2, 2, 2, 216, 217,
	5, 18, 10, 2, 217, 55, 3, 2, 2, 2, 218, 220, 7, 26, 2, 2, 219, 221, 5,
	30, 16, 2, 220, 219, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 222, 3, 2,
	2, 2, 222, 223, 7, 27, 2, 2, 223, 57, 3, 2, 2, 2, 224, 229, 7, 34, 2, 2,
	225, 229, 7, 2, 2, 3, 226, 229, 6, 30, 8, 2, 227, 229, 6, 30, 9, 2, 228,
	224, 3, 2, 2, 2, 228, 225, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 228, 227,
	3, 2, 2, 2, 229, 59, 3, 2, 2, 2, 22, 67, 73, 84, 88, 95, 107, 120, 124,
	133, 139, 156, 158, 169, 173, 179, 188, 200, 214, 220, 228,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'break'", "'default'", "'func'", "'interface'", "'case'", "'struct'",
	"'else'", "'goto'", "'package'", "'switch'", "'const'", "'fallthrough'",
	"'if'", "'range'", "'type'", "'continue'", "'for'", "'import'", "'return'",
	"'var'", "'list'", "'null'", "", "'('", "')'", "'{'", "'}'", "'['", "']'",
	"'='", "','", "';'", "':'", "'.'", "'++'", "'--'", "':='", "'...'", "'||'",
	"'&&'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'|'", "'/'", "'%'",
	"'<<'", "'>>'", "'&^'", "'!'", "'+'", "'-'", "'^'", "'*'", "'&'", "'<-'",
}
var symbolicNames = []string{
	"", "BREAK", "DEFAULT", "FUNC", "INTERFACE", "CASE", "STRUCT", "ELSE",
	"GOTO", "PACKAGE", "SWITCH", "CONST", "FALLTHROUGH", "IF", "RANGE", "TYPE",
	"CONTINUE", "FOR", "IMPORT", "RETURN", "VAR", "LIST", "NULL_LIT", "IDENTIFIER",
	"L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "ASSIGN",
	"COMMA", "SEMI", "COLON", "DOT", "PLUS_PLUS", "MINUS_MINUS", "DECLARE_ASSIGN",
	"ELLIPSIS", "LOGICAL_OR", "LOGICAL_AND", "EQUALS", "NOT_EQUALS", "LESS",
	"LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS", "OR", "DIV", "MOD", "LSHIFT",
	"RSHIFT", "BIT_CLEAR", "EXCLAMATION", "PLUS", "MINUS", "CARET", "STAR",
	"AMPERSAND", "RECEIVE", "DECIMAL_LIT", "OCTAL_LIT", "HEX_LIT", "FLOAT_LIT",
	"IMAGINARY_LIT", "RUNE_LIT", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT",
	"WS", "COMMENT", "TERMINATOR", "LINE_COMMENT", "NEWLINE",
}

var ruleNames = []string{
	"sourceFile", "statementList", "statement", "declaration", "varDecl", "varSpec",
	"identifierList", "type_", "typeName", "assignStatement", "assign_op",
	"expressionStmt", "ifStmt", "block", "expressionList", "expression", "primaryExpr",
	"unaryExpr", "operand", "literal", "string_", "basicLit", "integer", "operandName",
	"index", "methodExpr", "receiverType", "arguments", "eos",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type V2ParserParser struct {
	*antlr.BaseParser
}

func NewV2ParserParser(input antlr.TokenStream) *V2ParserParser {
	this := new(V2ParserParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "V2Parser.g4"

	return this
}

// V2ParserParser tokens.
const (
	V2ParserParserEOF                    = antlr.TokenEOF
	V2ParserParserBREAK                  = 1
	V2ParserParserDEFAULT                = 2
	V2ParserParserFUNC                   = 3
	V2ParserParserINTERFACE              = 4
	V2ParserParserCASE                   = 5
	V2ParserParserSTRUCT                 = 6
	V2ParserParserELSE                   = 7
	V2ParserParserGOTO                   = 8
	V2ParserParserPACKAGE                = 9
	V2ParserParserSWITCH                 = 10
	V2ParserParserCONST                  = 11
	V2ParserParserFALLTHROUGH            = 12
	V2ParserParserIF                     = 13
	V2ParserParserRANGE                  = 14
	V2ParserParserTYPE                   = 15
	V2ParserParserCONTINUE               = 16
	V2ParserParserFOR                    = 17
	V2ParserParserIMPORT                 = 18
	V2ParserParserRETURN                 = 19
	V2ParserParserVAR                    = 20
	V2ParserParserLIST                   = 21
	V2ParserParserNULL_LIT               = 22
	V2ParserParserIDENTIFIER             = 23
	V2ParserParserL_PAREN                = 24
	V2ParserParserR_PAREN                = 25
	V2ParserParserL_CURLY                = 26
	V2ParserParserR_CURLY                = 27
	V2ParserParserL_BRACKET              = 28
	V2ParserParserR_BRACKET              = 29
	V2ParserParserASSIGN                 = 30
	V2ParserParserCOMMA                  = 31
	V2ParserParserSEMI                   = 32
	V2ParserParserCOLON                  = 33
	V2ParserParserDOT                    = 34
	V2ParserParserPLUS_PLUS              = 35
	V2ParserParserMINUS_MINUS            = 36
	V2ParserParserDECLARE_ASSIGN         = 37
	V2ParserParserELLIPSIS               = 38
	V2ParserParserLOGICAL_OR             = 39
	V2ParserParserLOGICAL_AND            = 40
	V2ParserParserEQUALS                 = 41
	V2ParserParserNOT_EQUALS             = 42
	V2ParserParserLESS                   = 43
	V2ParserParserLESS_OR_EQUALS         = 44
	V2ParserParserGREATER                = 45
	V2ParserParserGREATER_OR_EQUALS      = 46
	V2ParserParserOR                     = 47
	V2ParserParserDIV                    = 48
	V2ParserParserMOD                    = 49
	V2ParserParserLSHIFT                 = 50
	V2ParserParserRSHIFT                 = 51
	V2ParserParserBIT_CLEAR              = 52
	V2ParserParserEXCLAMATION            = 53
	V2ParserParserPLUS                   = 54
	V2ParserParserMINUS                  = 55
	V2ParserParserCARET                  = 56
	V2ParserParserSTAR                   = 57
	V2ParserParserAMPERSAND              = 58
	V2ParserParserRECEIVE                = 59
	V2ParserParserDECIMAL_LIT            = 60
	V2ParserParserOCTAL_LIT              = 61
	V2ParserParserHEX_LIT                = 62
	V2ParserParserFLOAT_LIT              = 63
	V2ParserParserIMAGINARY_LIT          = 64
	V2ParserParserRUNE_LIT               = 65
	V2ParserParserRAW_STRING_LIT         = 66
	V2ParserParserINTERPRETED_STRING_LIT = 67
	V2ParserParserWS                     = 68
	V2ParserParserCOMMENT                = 69
	V2ParserParserTERMINATOR             = 70
	V2ParserParserLINE_COMMENT           = 71
	V2ParserParserNEWLINE                = 72
)

// V2ParserParser rules.
const (
	V2ParserParserRULE_sourceFile      = 0
	V2ParserParserRULE_statementList   = 1
	V2ParserParserRULE_statement       = 2
	V2ParserParserRULE_declaration     = 3
	V2ParserParserRULE_varDecl         = 4
	V2ParserParserRULE_varSpec         = 5
	V2ParserParserRULE_identifierList  = 6
	V2ParserParserRULE_type_           = 7
	V2ParserParserRULE_typeName        = 8
	V2ParserParserRULE_assignStatement = 9
	V2ParserParserRULE_assign_op       = 10
	V2ParserParserRULE_expressionStmt  = 11
	V2ParserParserRULE_ifStmt          = 12
	V2ParserParserRULE_block           = 13
	V2ParserParserRULE_expressionList  = 14
	V2ParserParserRULE_expression      = 15
	V2ParserParserRULE_primaryExpr     = 16
	V2ParserParserRULE_unaryExpr       = 17
	V2ParserParserRULE_operand         = 18
	V2ParserParserRULE_literal         = 19
	V2ParserParserRULE_string_         = 20
	V2ParserParserRULE_basicLit        = 21
	V2ParserParserRULE_integer         = 22
	V2ParserParserRULE_operandName     = 23
	V2ParserParserRULE_index           = 24
	V2ParserParserRULE_methodExpr      = 25
	V2ParserParserRULE_receiverType    = 26
	V2ParserParserRULE_arguments       = 27
	V2ParserParserRULE_eos             = 28
)

// ISourceFileContext is an interface to support dynamic dispatch.
type ISourceFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceFileContext differentiates from other interfaces.
	IsSourceFileContext()
}

type SourceFileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceFileContext() *SourceFileContext {
	var p = new(SourceFileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_sourceFile
	return p
}

func (*SourceFileContext) IsSourceFileContext() {}

func NewSourceFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceFileContext {
	var p = new(SourceFileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_sourceFile

	return p
}

func (s *SourceFileContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceFileContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *SourceFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceFileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitSourceFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) SourceFile() (localctx ISourceFileContext) {
	localctx = NewSourceFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, V2ParserParserRULE_sourceFile)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.StatementList()
	}

	return localctx
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *StatementListContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitStatementList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, V2ParserParserRULE_statementList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<V2ParserParserIF)|(1<<V2ParserParserVAR)|(1<<V2ParserParserNULL_LIT)|(1<<V2ParserParserIDENTIFIER)|(1<<V2ParserParserL_PAREN))) != 0) || (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(V2ParserParserEXCLAMATION-53))|(1<<(V2ParserParserPLUS-53))|(1<<(V2ParserParserMINUS-53))|(1<<(V2ParserParserCARET-53))|(1<<(V2ParserParserDECIMAL_LIT-53))|(1<<(V2ParserParserOCTAL_LIT-53))|(1<<(V2ParserParserHEX_LIT-53))|(1<<(V2ParserParserFLOAT_LIT-53))|(1<<(V2ParserParserIMAGINARY_LIT-53))|(1<<(V2ParserParserRUNE_LIT-53))|(1<<(V2ParserParserRAW_STRING_LIT-53))|(1<<(V2ParserParserINTERPRETED_STRING_LIT-53)))) != 0) {
		{
			p.SetState(60)
			p.Statement()
		}
		{
			p.SetState(61)
			p.Eos()
		}

		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) AssignStatement() IAssignStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignStatementContext)
}

func (s *StatementContext) ExpressionStmt() IExpressionStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionStmtContext)
}

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, V2ParserParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.Declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.AssignStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)
			p.ExpressionStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(70)
			p.IfStmt()
		}

	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) VarDecl() IVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, V2ParserParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.VarDecl()
	}

	return localctx
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_varDecl
	return p
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) VAR() antlr.TerminalNode {
	return s.GetToken(V2ParserParserVAR, 0)
}

func (s *VarDeclContext) VarSpec() IVarSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarSpecContext)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, V2ParserParserRULE_varDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(V2ParserParserVAR)
	}
	{
		p.SetState(76)
		p.VarSpec()
	}

	return localctx
}

// IVarSpecContext is an interface to support dynamic dispatch.
type IVarSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarSpecContext differentiates from other interfaces.
	IsVarSpecContext()
}

type VarSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarSpecContext() *VarSpecContext {
	var p = new(VarSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_varSpec
	return p
}

func (*VarSpecContext) IsVarSpecContext() {}

func NewVarSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarSpecContext {
	var p = new(VarSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_varSpec

	return p
}

func (s *VarSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *VarSpecContext) IdentifierList() IIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *VarSpecContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VarSpecContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(V2ParserParserASSIGN, 0)
}

func (s *VarSpecContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *VarSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitVarSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) VarSpec() (localctx IVarSpecContext) {
	localctx = NewVarSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, V2ParserParserRULE_varSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.IdentifierList()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case V2ParserParserIDENTIFIER:
		{
			p.SetState(79)
			p.Type_()
		}
		p.SetState(82)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(80)
				p.Match(V2ParserParserASSIGN)
			}
			{
				p.SetState(81)
				p.ExpressionList()
			}

		}

	case V2ParserParserASSIGN:
		{
			p.SetState(84)
			p.Match(V2ParserParserASSIGN)
		}
		{
			p.SetState(85)
			p.ExpressionList()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_identifierList
	return p
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(V2ParserParserIDENTIFIER)
}

func (s *IdentifierListContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserParserIDENTIFIER, i)
}

func (s *IdentifierListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(V2ParserParserCOMMA)
}

func (s *IdentifierListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserParserCOMMA, i)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, V2ParserParserRULE_identifierList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(V2ParserParserIDENTIFIER)
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == V2ParserParserCOMMA {
		{
			p.SetState(89)
			p.Match(V2ParserParserCOMMA)
		}
		{
			p.SetState(90)
			p.Match(V2ParserParserIDENTIFIER)
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitType_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, V2ParserParserRULE_type_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.TypeName()
	}

	return localctx
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_typeName
	return p
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(V2ParserParserIDENTIFIER, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, V2ParserParserRULE_typeName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(V2ParserParserIDENTIFIER)
	}

	return localctx
}

// IAssignStatementContext is an interface to support dynamic dispatch.
type IAssignStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignStatementContext differentiates from other interfaces.
	IsAssignStatementContext()
}

type AssignStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignStatementContext() *AssignStatementContext {
	var p = new(AssignStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_assignStatement
	return p
}

func (*AssignStatementContext) IsAssignStatementContext() {}

func NewAssignStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignStatementContext {
	var p = new(AssignStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_assignStatement

	return p
}

func (s *AssignStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignStatementContext) AllExpressionList() []IExpressionListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionListContext)(nil)).Elem())
	var tst = make([]IExpressionListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionListContext)
		}
	}

	return tst
}

func (s *AssignStatementContext) ExpressionList(i int) IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *AssignStatementContext) Assign_op() IAssign_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssign_opContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssign_opContext)
}

func (s *AssignStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitAssignStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) AssignStatement() (localctx IAssignStatementContext) {
	localctx = NewAssignStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, V2ParserParserRULE_assignStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.ExpressionList()
	}
	{
		p.SetState(101)
		p.Assign_op()
	}
	{
		p.SetState(102)
		p.ExpressionList()
	}

	return localctx
}

// IAssign_opContext is an interface to support dynamic dispatch.
type IAssign_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssign_opContext differentiates from other interfaces.
	IsAssign_opContext()
}

type Assign_opContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_opContext() *Assign_opContext {
	var p = new(Assign_opContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_assign_op
	return p
}

func (*Assign_opContext) IsAssign_opContext() {}

func NewAssign_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_opContext {
	var p = new(Assign_opContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_assign_op

	return p
}

func (s *Assign_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_opContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(V2ParserParserASSIGN, 0)
}

func (s *Assign_opContext) PLUS() antlr.TerminalNode {
	return s.GetToken(V2ParserParserPLUS, 0)
}

func (s *Assign_opContext) MINUS() antlr.TerminalNode {
	return s.GetToken(V2ParserParserMINUS, 0)
}

func (s *Assign_opContext) OR() antlr.TerminalNode {
	return s.GetToken(V2ParserParserOR, 0)
}

func (s *Assign_opContext) CARET() antlr.TerminalNode {
	return s.GetToken(V2ParserParserCARET, 0)
}

func (s *Assign_opContext) STAR() antlr.TerminalNode {
	return s.GetToken(V2ParserParserSTAR, 0)
}

func (s *Assign_opContext) DIV() antlr.TerminalNode {
	return s.GetToken(V2ParserParserDIV, 0)
}

func (s *Assign_opContext) MOD() antlr.TerminalNode {
	return s.GetToken(V2ParserParserMOD, 0)
}

func (s *Assign_opContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserLSHIFT, 0)
}

func (s *Assign_opContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserRSHIFT, 0)
}

func (s *Assign_opContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(V2ParserParserAMPERSAND, 0)
}

func (s *Assign_opContext) BIT_CLEAR() antlr.TerminalNode {
	return s.GetToken(V2ParserParserBIT_CLEAR, 0)
}

func (s *Assign_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_opContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitAssign_op(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) Assign_op() (localctx IAssign_opContext) {
	localctx = NewAssign_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, V2ParserParserRULE_assign_op)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(V2ParserParserOR-47))|(1<<(V2ParserParserDIV-47))|(1<<(V2ParserParserMOD-47))|(1<<(V2ParserParserLSHIFT-47))|(1<<(V2ParserParserRSHIFT-47))|(1<<(V2ParserParserBIT_CLEAR-47))|(1<<(V2ParserParserPLUS-47))|(1<<(V2ParserParserMINUS-47))|(1<<(V2ParserParserCARET-47))|(1<<(V2ParserParserSTAR-47))|(1<<(V2ParserParserAMPERSAND-47)))) != 0 {
		p.SetState(104)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(V2ParserParserOR-47))|(1<<(V2ParserParserDIV-47))|(1<<(V2ParserParserMOD-47))|(1<<(V2ParserParserLSHIFT-47))|(1<<(V2ParserParserRSHIFT-47))|(1<<(V2ParserParserBIT_CLEAR-47))|(1<<(V2ParserParserPLUS-47))|(1<<(V2ParserParserMINUS-47))|(1<<(V2ParserParserCARET-47))|(1<<(V2ParserParserSTAR-47))|(1<<(V2ParserParserAMPERSAND-47)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}
	{
		p.SetState(107)
		p.Match(V2ParserParserASSIGN)
	}

	return localctx
}

// IExpressionStmtContext is an interface to support dynamic dispatch.
type IExpressionStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionStmtContext differentiates from other interfaces.
	IsExpressionStmtContext()
}

type ExpressionStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStmtContext() *ExpressionStmtContext {
	var p = new(ExpressionStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_expressionStmt
	return p
}

func (*ExpressionStmtContext) IsExpressionStmtContext() {}

func NewExpressionStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStmtContext {
	var p = new(ExpressionStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_expressionStmt

	return p
}

func (s *ExpressionStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitExpressionStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) ExpressionStmt() (localctx IExpressionStmtContext) {
	localctx = NewExpressionStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, V2ParserParserRULE_expressionStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.expression(0)
	}

	return localctx
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_ifStmt
	return p
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(V2ParserParserIF, 0)
}

func (s *IfStmtContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(V2ParserParserL_PAREN, 0)
}

func (s *IfStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(V2ParserParserR_PAREN, 0)
}

func (s *IfStmtContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(V2ParserParserELSE, 0)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, V2ParserParserRULE_ifStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(V2ParserParserIF)
	}
	{
		p.SetState(112)
		p.Match(V2ParserParserL_PAREN)
	}
	{
		p.SetState(113)
		p.expression(0)
	}
	{
		p.SetState(114)
		p.Match(V2ParserParserR_PAREN)
	}
	{
		p.SetState(115)
		p.Block()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(116)
			p.Match(V2ParserParserELSE)
		}
		{
			p.SetState(117)
			p.Block()
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(V2ParserParserL_CURLY, 0)
}

func (s *BlockContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(V2ParserParserR_CURLY, 0)
}

func (s *BlockContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, V2ParserParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(V2ParserParserL_CURLY)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<V2ParserParserIF)|(1<<V2ParserParserVAR)|(1<<V2ParserParserNULL_LIT)|(1<<V2ParserParserIDENTIFIER)|(1<<V2ParserParserL_PAREN))) != 0) || (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(V2ParserParserEXCLAMATION-53))|(1<<(V2ParserParserPLUS-53))|(1<<(V2ParserParserMINUS-53))|(1<<(V2ParserParserCARET-53))|(1<<(V2ParserParserDECIMAL_LIT-53))|(1<<(V2ParserParserOCTAL_LIT-53))|(1<<(V2ParserParserHEX_LIT-53))|(1<<(V2ParserParserFLOAT_LIT-53))|(1<<(V2ParserParserIMAGINARY_LIT-53))|(1<<(V2ParserParserRUNE_LIT-53))|(1<<(V2ParserParserRAW_STRING_LIT-53))|(1<<(V2ParserParserINTERPRETED_STRING_LIT-53)))) != 0) {
		{
			p.SetState(121)
			p.StatementList()
		}

	}
	{
		p.SetState(124)
		p.Match(V2ParserParserR_CURLY)
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(V2ParserParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(V2ParserParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, V2ParserParserRULE_expressionList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.expression(0)
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(127)
				p.Match(V2ParserParserCOMMA)
			}
			{
				p.SetState(128)
				p.expression(0)
			}

		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *ExpressionContext) UnaryExpr() IUnaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) STAR() antlr.TerminalNode {
	return s.GetToken(V2ParserParserSTAR, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(V2ParserParserDIV, 0)
}

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(V2ParserParserMOD, 0)
}

func (s *ExpressionContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserLSHIFT, 0)
}

func (s *ExpressionContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserRSHIFT, 0)
}

func (s *ExpressionContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(V2ParserParserAMPERSAND, 0)
}

func (s *ExpressionContext) BIT_CLEAR() antlr.TerminalNode {
	return s.GetToken(V2ParserParserBIT_CLEAR, 0)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(V2ParserParserPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(V2ParserParserMINUS, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(V2ParserParserOR, 0)
}

func (s *ExpressionContext) CARET() antlr.TerminalNode {
	return s.GetToken(V2ParserParserCARET, 0)
}

func (s *ExpressionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(V2ParserParserEQUALS, 0)
}

func (s *ExpressionContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(V2ParserParserNOT_EQUALS, 0)
}

func (s *ExpressionContext) LESS() antlr.TerminalNode {
	return s.GetToken(V2ParserParserLESS, 0)
}

func (s *ExpressionContext) LESS_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(V2ParserParserLESS_OR_EQUALS, 0)
}

func (s *ExpressionContext) GREATER() antlr.TerminalNode {
	return s.GetToken(V2ParserParserGREATER, 0)
}

func (s *ExpressionContext) GREATER_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(V2ParserParserGREATER_OR_EQUALS, 0)
}

func (s *ExpressionContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(V2ParserParserLOGICAL_AND, 0)
}

func (s *ExpressionContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(V2ParserParserLOGICAL_OR, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *V2ParserParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, V2ParserParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(135)
			p.primaryExpr(0)
		}

	case 2:
		{
			p.SetState(136)
			p.UnaryExpr()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(154)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, V2ParserParserRULE_expression)
				p.SetState(139)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				p.SetState(140)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(V2ParserParserDIV-48))|(1<<(V2ParserParserMOD-48))|(1<<(V2ParserParserLSHIFT-48))|(1<<(V2ParserParserRSHIFT-48))|(1<<(V2ParserParserBIT_CLEAR-48))|(1<<(V2ParserParserSTAR-48))|(1<<(V2ParserParserAMPERSAND-48)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(141)
					p.expression(6)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, V2ParserParserRULE_expression)
				p.SetState(142)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				p.SetState(143)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(V2ParserParserOR-47))|(1<<(V2ParserParserPLUS-47))|(1<<(V2ParserParserMINUS-47))|(1<<(V2ParserParserCARET-47)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(144)
					p.expression(5)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, V2ParserParserRULE_expression)
				p.SetState(145)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				p.SetState(146)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(V2ParserParserEQUALS-41))|(1<<(V2ParserParserNOT_EQUALS-41))|(1<<(V2ParserParserLESS-41))|(1<<(V2ParserParserLESS_OR_EQUALS-41))|(1<<(V2ParserParserGREATER-41))|(1<<(V2ParserParserGREATER_OR_EQUALS-41)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(147)
					p.expression(4)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, V2ParserParserRULE_expression)
				p.SetState(148)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(149)
					p.Match(V2ParserParserLOGICAL_AND)
				}
				{
					p.SetState(150)
					p.expression(3)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, V2ParserParserRULE_expression)
				p.SetState(151)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(152)
					p.Match(V2ParserParserLOGICAL_OR)
				}
				{
					p.SetState(153)
					p.expression(2)
				}

			}

		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_primaryExpr
	return p
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) Operand() IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PrimaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *PrimaryExprContext) DOT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserDOT, 0)
}

func (s *PrimaryExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(V2ParserParserIDENTIFIER, 0)
}

func (s *PrimaryExprContext) Index() IIndexContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *PrimaryExprContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	return p.primaryExpr(0)
}

func (p *V2ParserParser) primaryExpr(_p int) (localctx IPrimaryExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, V2ParserParserRULE_primaryExpr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Operand()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, V2ParserParserRULE_primaryExpr)
			p.SetState(162)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(167)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case V2ParserParserDOT:
				{
					p.SetState(163)
					p.Match(V2ParserParserDOT)
				}
				{
					p.SetState(164)
					p.Match(V2ParserParserIDENTIFIER)
				}

			case V2ParserParserL_BRACKET:
				{
					p.SetState(165)
					p.Index()
				}

			case V2ParserParserL_PAREN:
				{
					p.SetState(166)
					p.Arguments()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryExprContext is an interface to support dynamic dispatch.
type IUnaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryExprContext differentiates from other interfaces.
	IsUnaryExprContext()
}

type UnaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExprContext() *UnaryExprContext {
	var p = new(UnaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_unaryExpr
	return p
}

func (*UnaryExprContext) IsUnaryExprContext() {}

func NewUnaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExprContext {
	var p = new(UnaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_unaryExpr

	return p
}

func (s *UnaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *UnaryExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(V2ParserParserPLUS, 0)
}

func (s *UnaryExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(V2ParserParserMINUS, 0)
}

func (s *UnaryExprContext) EXCLAMATION() antlr.TerminalNode {
	return s.GetToken(V2ParserParserEXCLAMATION, 0)
}

func (s *UnaryExprContext) CARET() antlr.TerminalNode {
	return s.GetToken(V2ParserParserCARET, 0)
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) UnaryExpr() (localctx IUnaryExprContext) {
	localctx = NewUnaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, V2ParserParserRULE_unaryExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(177)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case V2ParserParserNULL_LIT, V2ParserParserIDENTIFIER, V2ParserParserL_PAREN, V2ParserParserDECIMAL_LIT, V2ParserParserOCTAL_LIT, V2ParserParserHEX_LIT, V2ParserParserFLOAT_LIT, V2ParserParserIMAGINARY_LIT, V2ParserParserRUNE_LIT, V2ParserParserRAW_STRING_LIT, V2ParserParserINTERPRETED_STRING_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(174)
			p.primaryExpr(0)
		}

	case V2ParserParserEXCLAMATION, V2ParserParserPLUS, V2ParserParserMINUS, V2ParserParserCARET:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(175)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(V2ParserParserEXCLAMATION-53))|(1<<(V2ParserParserPLUS-53))|(1<<(V2ParserParserMINUS-53))|(1<<(V2ParserParserCARET-53)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(176)
			p.expression(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *OperandContext) OperandName() IOperandNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandNameContext)
}

func (s *OperandContext) MethodExpr() IMethodExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodExprContext)
}

func (s *OperandContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(V2ParserParserL_PAREN, 0)
}

func (s *OperandContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OperandContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(V2ParserParserR_PAREN, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, V2ParserParserRULE_operand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)
			p.OperandName()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(181)
			p.MethodExpr()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(182)
			p.Match(V2ParserParserL_PAREN)
		}
		{
			p.SetState(183)
			p.expression(0)
		}
		{
			p.SetState(184)
			p.Match(V2ParserParserR_PAREN)
		}

	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) BasicLit() IBasicLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasicLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasicLitContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, V2ParserParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.BasicLit()
	}

	return localctx
}

// IString_Context is an interface to support dynamic dispatch.
type IString_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_Context differentiates from other interfaces.
	IsString_Context()
}

type String_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_Context() *String_Context {
	var p = new(String_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_string_
	return p
}

func (*String_Context) IsString_Context() {}

func NewString_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_Context {
	var p = new(String_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_string_

	return p
}

func (s *String_Context) GetParser() antlr.Parser { return s.parser }

func (s *String_Context) RAW_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserRAW_STRING_LIT, 0)
}

func (s *String_Context) INTERPRETED_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserINTERPRETED_STRING_LIT, 0)
}

func (s *String_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitString_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) String_() (localctx IString_Context) {
	localctx = NewString_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, V2ParserParserRULE_string_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(190)
	_la = p.GetTokenStream().LA(1)

	if !(_la == V2ParserParserRAW_STRING_LIT || _la == V2ParserParserINTERPRETED_STRING_LIT) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IBasicLitContext is an interface to support dynamic dispatch.
type IBasicLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasicLitContext differentiates from other interfaces.
	IsBasicLitContext()
}

type BasicLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasicLitContext() *BasicLitContext {
	var p = new(BasicLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_basicLit
	return p
}

func (*BasicLitContext) IsBasicLitContext() {}

func NewBasicLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicLitContext {
	var p = new(BasicLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_basicLit

	return p
}

func (s *BasicLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicLitContext) NULL_LIT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserNULL_LIT, 0)
}

func (s *BasicLitContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *BasicLitContext) String_() IString_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *BasicLitContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserFLOAT_LIT, 0)
}

func (s *BasicLitContext) IMAGINARY_LIT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserIMAGINARY_LIT, 0)
}

func (s *BasicLitContext) RUNE_LIT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserRUNE_LIT, 0)
}

func (s *BasicLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitBasicLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) BasicLit() (localctx IBasicLitContext) {
	localctx = NewBasicLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, V2ParserParserRULE_basicLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.Match(V2ParserParserNULL_LIT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.Integer()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(194)
			p.String_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(195)
			p.Match(V2ParserParserFLOAT_LIT)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(196)
			p.Match(V2ParserParserIMAGINARY_LIT)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(197)
			p.Match(V2ParserParserRUNE_LIT)
		}

	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) DECIMAL_LIT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserDECIMAL_LIT, 0)
}

func (s *IntegerContext) OCTAL_LIT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserOCTAL_LIT, 0)
}

func (s *IntegerContext) HEX_LIT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserHEX_LIT, 0)
}

func (s *IntegerContext) IMAGINARY_LIT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserIMAGINARY_LIT, 0)
}

func (s *IntegerContext) RUNE_LIT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserRUNE_LIT, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, V2ParserParserRULE_integer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(200)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-60)&-(0x1f+1)) == 0 && ((1<<uint((_la-60)))&((1<<(V2ParserParserDECIMAL_LIT-60))|(1<<(V2ParserParserOCTAL_LIT-60))|(1<<(V2ParserParserHEX_LIT-60))|(1<<(V2ParserParserIMAGINARY_LIT-60))|(1<<(V2ParserParserRUNE_LIT-60)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IOperandNameContext is an interface to support dynamic dispatch.
type IOperandNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandNameContext differentiates from other interfaces.
	IsOperandNameContext()
}

type OperandNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandNameContext() *OperandNameContext {
	var p = new(OperandNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_operandName
	return p
}

func (*OperandNameContext) IsOperandNameContext() {}

func NewOperandNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandNameContext {
	var p = new(OperandNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_operandName

	return p
}

func (s *OperandNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(V2ParserParserIDENTIFIER, 0)
}

func (s *OperandNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitOperandName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) OperandName() (localctx IOperandNameContext) {
	localctx = NewOperandNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, V2ParserParserRULE_operandName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(V2ParserParserIDENTIFIER)
	}

	return localctx
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_index
	return p
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(V2ParserParserL_BRACKET, 0)
}

func (s *IndexContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(V2ParserParserR_BRACKET, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, V2ParserParserRULE_index)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(V2ParserParserL_BRACKET)
	}
	{
		p.SetState(205)
		p.expression(0)
	}
	{
		p.SetState(206)
		p.Match(V2ParserParserR_BRACKET)
	}

	return localctx
}

// IMethodExprContext is an interface to support dynamic dispatch.
type IMethodExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodExprContext differentiates from other interfaces.
	IsMethodExprContext()
}

type MethodExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodExprContext() *MethodExprContext {
	var p = new(MethodExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_methodExpr
	return p
}

func (*MethodExprContext) IsMethodExprContext() {}

func NewMethodExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodExprContext {
	var p = new(MethodExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_methodExpr

	return p
}

func (s *MethodExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodExprContext) ReceiverType() IReceiverTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReceiverTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReceiverTypeContext)
}

func (s *MethodExprContext) DOT() antlr.TerminalNode {
	return s.GetToken(V2ParserParserDOT, 0)
}

func (s *MethodExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(V2ParserParserIDENTIFIER, 0)
}

func (s *MethodExprContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *MethodExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitMethodExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) MethodExpr() (localctx IMethodExprContext) {
	localctx = NewMethodExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, V2ParserParserRULE_methodExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.ReceiverType()
	}
	{
		p.SetState(209)
		p.Match(V2ParserParserDOT)
	}
	{
		p.SetState(210)
		p.Match(V2ParserParserIDENTIFIER)
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(211)
			p.Arguments()
		}

	}

	return localctx
}

// IReceiverTypeContext is an interface to support dynamic dispatch.
type IReceiverTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReceiverTypeContext differentiates from other interfaces.
	IsReceiverTypeContext()
}

type ReceiverTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReceiverTypeContext() *ReceiverTypeContext {
	var p = new(ReceiverTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_receiverType
	return p
}

func (*ReceiverTypeContext) IsReceiverTypeContext() {}

func NewReceiverTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReceiverTypeContext {
	var p = new(ReceiverTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_receiverType

	return p
}

func (s *ReceiverTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReceiverTypeContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *ReceiverTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReceiverTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReceiverTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitReceiverType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) ReceiverType() (localctx IReceiverTypeContext) {
	localctx = NewReceiverTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, V2ParserParserRULE_receiverType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.TypeName()
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(V2ParserParserL_PAREN, 0)
}

func (s *ArgumentsContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(V2ParserParserR_PAREN, 0)
}

func (s *ArgumentsContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, V2ParserParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(V2ParserParserL_PAREN)
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<V2ParserParserNULL_LIT)|(1<<V2ParserParserIDENTIFIER)|(1<<V2ParserParserL_PAREN))) != 0) || (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(V2ParserParserEXCLAMATION-53))|(1<<(V2ParserParserPLUS-53))|(1<<(V2ParserParserMINUS-53))|(1<<(V2ParserParserCARET-53))|(1<<(V2ParserParserDECIMAL_LIT-53))|(1<<(V2ParserParserOCTAL_LIT-53))|(1<<(V2ParserParserHEX_LIT-53))|(1<<(V2ParserParserFLOAT_LIT-53))|(1<<(V2ParserParserIMAGINARY_LIT-53))|(1<<(V2ParserParserRUNE_LIT-53))|(1<<(V2ParserParserRAW_STRING_LIT-53))|(1<<(V2ParserParserINTERPRETED_STRING_LIT-53)))) != 0) {
		{
			p.SetState(217)
			p.ExpressionList()
		}

	}
	{
		p.SetState(220)
		p.Match(V2ParserParserR_PAREN)
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = V2ParserParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = V2ParserParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) SEMI() antlr.TerminalNode {
	return s.GetToken(V2ParserParserSEMI, 0)
}

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(V2ParserParserEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case V2ParserVisitor:
		return t.VisitEos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *V2ParserParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, V2ParserParserRULE_eos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)
			p.Match(V2ParserParserSEMI)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(223)
			p.Match(V2ParserParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(224)

		if !(p.lineTerminatorAhead()) {
			panic(antlr.NewFailedPredicateException(p, "lineTerminatorAhead()", ""))
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(225)

		if !(p.checkPreviousTokenText("}")) {
			panic(antlr.NewFailedPredicateException(p, "checkPreviousTokenText(\"}\")", ""))
		}

	}

	return localctx
}

func (p *V2ParserParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 15:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 16:
		var t *PrimaryExprContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExprContext)
		}
		return p.PrimaryExpr_Sempred(t, predIndex)

	case 28:
		var t *EosContext = nil
		if localctx != nil {
			t = localctx.(*EosContext)
		}
		return p.Eos_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *V2ParserParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *V2ParserParser) PrimaryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *V2ParserParser) Eos_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.lineTerminatorAhead()

	case 7:
		return p.checkPreviousTokenText("}")

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func  (p *V2ParserParser)lineTerminatorAhead() bool {
	// Get the token ahead of the current index.
	possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
	ahead := p.GetTokenStream().Get(possibleIndexEosToken)

	if ahead.GetChannel() != antlr.LexerHidden {
		// We're only interested in tokens on the HIDDEN channel.
		return true
	}

	if ahead.GetTokenType() == V2ParserLexerTERMINATOR {
		// There is definitely a line terminator ahead.
		return true
	}

	if ahead.GetTokenType() == V2ParserLexerWS {
		// Get the token ahead of the current whitespaces.
		possibleIndexEosToken = p.GetCurrentToken().GetTokenIndex() - 2
		ahead = p.GetTokenStream().Get(possibleIndexEosToken)
	}

	// Get the token's text and type.
	text := ahead.GetText()
	_type := ahead.GetTokenType()

	// Check if the token is, or contains a line terminator.
	return (_type == V2ParserLexerCOMMA && (strings.Contains(text, "\r") || strings.Contains(text, "\n"))) ||
		(_type == V2ParserLexerTERMINATOR)
}

func (p *V2ParserParser) checkPreviousTokenText(text string) bool {
	stream := p.GetTokenStream()
	return stream.LT(1).GetText() == text
}