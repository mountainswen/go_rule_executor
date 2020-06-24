// Generated from V2Parser.g4 by ANTLR 4.7.

package parser // V2Parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by V2ParserParser.
type V2ParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by V2ParserParser#sourceFile.
	VisitSourceFile(ctx *SourceFileContext) interface{}

	// Visit a parse tree produced by V2ParserParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by V2ParserParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by V2ParserParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by V2ParserParser#importStmt.
	VisitImportStmt(ctx *ImportStmtContext) interface{}

	// Visit a parse tree produced by V2ParserParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by V2ParserParser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by V2ParserParser#varSpec.
	VisitVarSpec(ctx *VarSpecContext) interface{}

	// Visit a parse tree produced by V2ParserParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by V2ParserParser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by V2ParserParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by V2ParserParser#assignStatement.
	VisitAssignStatement(ctx *AssignStatementContext) interface{}

	// Visit a parse tree produced by V2ParserParser#assign_op.
	VisitAssign_op(ctx *Assign_opContext) interface{}

	// Visit a parse tree produced by V2ParserParser#expressionStmt.
	VisitExpressionStmt(ctx *ExpressionStmtContext) interface{}

	// Visit a parse tree produced by V2ParserParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by V2ParserParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by V2ParserParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by V2ParserParser#keyValue.
	VisitKeyValue(ctx *KeyValueContext) interface{}

	// Visit a parse tree produced by V2ParserParser#keyValues.
	VisitKeyValues(ctx *KeyValuesContext) interface{}

	// Visit a parse tree produced by V2ParserParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by V2ParserParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by V2ParserParser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by V2ParserParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by V2ParserParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by V2ParserParser#string_.
	VisitString_(ctx *String_Context) interface{}

	// Visit a parse tree produced by V2ParserParser#basicLit.
	VisitBasicLit(ctx *BasicLitContext) interface{}

	// Visit a parse tree produced by V2ParserParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by V2ParserParser#operandName.
	VisitOperandName(ctx *OperandNameContext) interface{}

	// Visit a parse tree produced by V2ParserParser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by V2ParserParser#methodExpr.
	VisitMethodExpr(ctx *MethodExprContext) interface{}

	// Visit a parse tree produced by V2ParserParser#receiverType.
	VisitReceiverType(ctx *ReceiverTypeContext) interface{}

	// Visit a parse tree produced by V2ParserParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by V2ParserParser#eos.
	VisitEos(ctx *EosContext) interface{}
}
