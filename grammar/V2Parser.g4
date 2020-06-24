grammar V2Parser;
import V2Lexer;

sourceFile
    :statementList
    ;

statementList
    :(statement eos)+
    ;

statement
    :declaration
    |importStmt
    |assignStatement
    |expressionStmt
    |ifStmt
    |forStmt
    ;

forStmt
    : FOR L_PAREN assignStatement SEMI expression SEMI assignStatement R_PAREN  block  //循环语句
    ;

importStmt
    :IMPORT L_PAREN IDENTIFIER R_PAREN
    ;

//声明语句
declaration
    :varDecl
    ;

varDecl
    : VAR varSpec
    ;

varSpec
    : identifierList (type_ (ASSIGN expressionList)? | ASSIGN expressionList)
    ;

identifierList
    : IDENTIFIER (COMMA IDENTIFIER)*
    ;

type_
    : typeName
    ;

typeName
    : IDENTIFIER
    ;

//赋值语句
assignStatement
    :expressionList (assign_op | DECLARE_ASSIGN) expressionList
    ;

assign_op
    : (PLUS | MINUS | OR | CARET | STAR | DIV | MOD | LSHIFT | RSHIFT | AMPERSAND | BIT_CLEAR)? ASSIGN
    ;

expressionStmt
    : expression
    ;

ifStmt
    : IF  L_PAREN expression R_PAREN block (ELSE  block)?
    ;

block
    : L_CURLY statementList? R_CURLY
    ;

expressionList
    :expression (COMMA expression)*
    ;

keyValue
    :(integer|string_) COLON expression
    ;

keyValues
    : keyValue (COMMA keyValue)*
    ;

//表达式
expression
    : primaryExpr
    | unaryExpr
    | expression (STAR | DIV | MOD | LSHIFT | RSHIFT | AMPERSAND | BIT_CLEAR) expression
    | expression (PLUS | MINUS | OR | CARET) expression
    | expression (EQUALS | NOT_EQUALS | LESS | LESS_OR_EQUALS | GREATER | GREATER_OR_EQUALS) expression
    | expression LOGICAL_AND expression
    | expression LOGICAL_OR expression
    ;

primaryExpr
    : operand
    | primaryExpr ( DOT IDENTIFIER
                  | index
                  | arguments)
    ;

unaryExpr
    : primaryExpr
    | (PLUS | MINUS | EXCLAMATION | CARET) expression
    ;

operand
    : literal
    | operandName
    | methodExpr
    | L_PAREN expression R_PAREN
    | L_CURLY (expressionList | keyValues) R_CURLY
    ;

literal
    : basicLit
    ;

string_
    : RAW_STRING_LIT
    | INTERPRETED_STRING_LIT
    ;

basicLit
    : NULL_LIT
    | integer
    | string_
    | FLOAT_LIT
    | IMAGINARY_LIT
    | RUNE_LIT
    ;

integer
    : DECIMAL_LIT
    | OCTAL_LIT
    | HEX_LIT
    | IMAGINARY_LIT
    | RUNE_LIT
    ;

operandName
    : IDENTIFIER
    ;

index
    : L_BRACKET expression R_BRACKET
    ;

methodExpr
    : receiverType DOT IDENTIFIER
    ;

receiverType
    : typeName
    ;

arguments
    : L_PAREN expressionList? R_PAREN
    ;

eos
    : SEMI
    | EOF
    | {lineTerminatorAhead()}?
    | {checkPreviousTokenText("}")}?
    ;

