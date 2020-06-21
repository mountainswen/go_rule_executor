package main

import (
	parser "antlr_test/V2parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
	"strconv"
)

const (
	RightValue = 1
	LeftValue = 2
)

type VisitorImpl struct {
	cst  map[string]*Value
	env map[string]*Value
	*parser.BaseV2ParserVisitor
}

func NewVisitor()*VisitorImpl{
	v := &VisitorImpl{}
	v.cst = make(map[string]*Value)
	v.env = make(map[string]*Value)

	v.cst["true"] = valueTrue
	v.cst["false"] = valueFalse
	return v
}

func (v *VisitorImpl)PrintEnv(){
	fmt.Printf("%+v\n",v.env)
}

func (v *VisitorImpl) VisitChild(node antlr.RuleNode) interface{} {
	return node.Accept(v)
}

func (v *VisitorImpl) VisitSourceFile(ctx *parser.SourceFileContext) interface{} {
	fmt.Println("visit source file")
	stmtList := ctx.StatementList()
	return v.VisitChild(stmtList)
}

func (v *VisitorImpl) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	fmt.Println("enter visit StatementList")
	for _, stmt := range ctx.AllStatement() {
		v.VisitChild(stmt) //stmt没有返回值
	}

	fmt.Println("exist visit StatementList")
	return nil
}

func (v *VisitorImpl) VisitStatement(ctx *parser.StatementContext) interface{} {
	fmt.Println("enter visit VisitStatement")

	if ctx.AssignStatement() != nil {
		return v.VisitChild(ctx.AssignStatement())
	}else if ctx.IfStmt() != nil {
		return v.VisitChild(ctx.IfStmt())
	}else if ctx.ExpressionStmt() != nil {

	}else if ctx.Declaration() != nil {

	}else {
		panic("undefine statement")
	}

	return nil
}

func (v *VisitorImpl) VisitAssignStatement(ctx *parser.AssignStatementContext) interface{} {
	fmt.Println("enter visit VisitAssignStatement")
	if len(ctx.AllExpressionList()) != 2 {
		fmt.Println("invalid assignment statement")
	}

	left := v.VisitChild(ctx.AllExpressionList()[0])
	right := v.VisitChild(ctx.AllExpressionList()[1])

	//判断left是否是左值
	lvalues, ok := left.([]*Value)
	if !ok {
		panic("iiii")
	}
	fmt.Printf("lvalue:%+v\n",lvalues)

	//求right的个个表达式的值
	rvalues, ok := right.([]*Value)
	if !ok {
		fmt.Println("invalid right value")
		return nil
	}
	fmt.Printf("rvalue:%+v\n",rvalues)
	if len(lvalues) != len(rvalues) {
		panic("rvalue.len != rvalue.len")
	}

	for i,lv := range lvalues {
		if lv.rlType != LeftValue {
			panic("need left value!")
		}
		v.env[lv.symbol] = rvalues[i]
	}


	return nil
}

func (v *VisitorImpl) VisitAssign_op(ctx *parser.Assign_opContext) interface{} {
	fmt.Println("enter visit VisitAssign_op")

	return nil
}

func (v *VisitorImpl) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	fmt.Println("enter visit VisitExpressionList")
	var values []*Value
	for _, expr := range ctx.AllExpression() {
		vs := v.VisitChild(expr).(*Value)
		values = append(values,vs)
	}

	return values
}

func (v *VisitorImpl) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	fmt.Println("enter visit VisitExpression")
	if len(ctx.AllExpression()) > 0 {
		left := v.VisitChild(ctx.Expression(0)).(*Value)
		right := v.VisitChild(ctx.Expression(1)).(*Value)

		if ctx.PLUS() != nil { //加法
			return left.Add(right)
		}
		if ctx.STAR() != nil {
			return left.Mul(right)
		}
		if ctx.MINUS() != nil {
			return left.Sub(right)
		}
		if ctx.DIV() != nil {
			return left.Div(right)
		}

		if ctx.EQUALS() != nil {
			return left.Equal(right)
		}
		if ctx.NOT_EQUALS() != nil {
			return left.NotEqual(right)
		}
		//bool left < right
		if ctx.LESS() != nil {
			return left.Less(right)
		}
		if ctx.LESS_OR_EQUALS() != nil {
			return left.LessEq(right)
		}

		if ctx.GREATER() != nil { //left > right
			return left.Greater(right)
		}
		if ctx.GREATER_OR_EQUALS() != nil {
			return left.GreaterEq(right)
		}

		if ctx.LOGICAL_AND() != nil { //&&
			return left.And(right)
		}
		if ctx.LOGICAL_OR() != nil { // ||
			return left.And(left)
		}
	}

	value := v.VisitChild(ctx.PrimaryExpr()).(*Value)
	return value
}

func (v *VisitorImpl) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) interface{} {
	fmt.Println("enter visit VisitPrimaryExpr")

	if ctx.Operand() != nil {
		return v.VisitChild(ctx.Operand()).(*Value)
	}
	panic("unsupprt op")
	return nil
}

func (v *VisitorImpl) VisitOperand(ctx *parser.OperandContext) interface{} {
	fmt.Println("enter visit VisitOperand")
	if ctx.OperandName() != nil {
		return v.VisitChild(ctx.OperandName()).(*Value)
	}
	if ctx.Literal() != nil {
		return v.VisitChild(ctx.Literal()).(*Value)
	}
	if ctx.Expression() != nil {
		return v.VisitChild(ctx.Expression())
	}
	return nil
}

func (v *VisitorImpl) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	fmt.Println("enter visit VisitLiteral")
	return v.VisitChild(ctx.BasicLit()).(*Value)
}

func (v *VisitorImpl) VisitBasicLit(ctx *parser.BasicLitContext) interface{} {
	fmt.Println("enter visit VisitBasicLit")
	if ctx.Integer() != nil { //整型
		return v.VisitChild(ctx.Integer())
	}
	if ctx.FLOAT_LIT() != nil { //浮点数
		value := &Value{}
		fstr := ctx.FLOAT_LIT().GetText()
		f, _ := strconv.ParseFloat(fstr,64)
		value.rlType = 1
		value.vConst = f
		value.vType = 2
		return value
	}
	if ctx.String_() != nil {
		return v.VisitChild(ctx.String_())
	}

	return nil
}

func (v *VisitorImpl) VisitOperandName(ctx *parser.OperandNameContext) interface{} {
	fmt.Println("enter visit VisitOperandName")
	if ctx.IDENTIFIER() != nil {
		txt := ctx.GetText()
		value := &Value{}
		value.rlType = LeftValue
		value.symbol = txt
		value.vType = 2

		if val, ok := v.cst[txt];ok {
			return val
		}

		if val, ok := v.env[txt];ok {
			value.vConst = val.vConst
		}
		return value
	}

	return nil
}

func (v *VisitorImpl) VisitInteger(ctx *parser.IntegerContext) interface{} {
	fmt.Println("enter visit VisitInteger")
	txt := ctx.GetText()
	value := &Value{}
	value.rlType = 1
	value.vConst,_ = strconv.ParseInt(txt,10,64)
	value.vType = 2

	return value
}


func (v *VisitorImpl) VisitString_(ctx *parser.String_Context) interface{} {
	if ctx.RAW_STRING_LIT() != nil {
		fstr := ctx.RAW_STRING_LIT().GetText()
		value := &Value{}
		value.rlType = 1
		value.vConst = fstr
		value.vType = 2
		return value
	}

	if ctx.INTERPRETED_STRING_LIT() != nil {
		fstr := ctx.INTERPRETED_STRING_LIT().GetText()
		value := &Value{}
		value.rlType = 1
		value.vConst = fstr[1:len(fstr)-1] //trim掉双引号
		value.vType = 2
		return value
	}

	return nil
}


//if
func (v *VisitorImpl) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	cond := ctx.Expression()
	if cond == nil {
		panic("")
	}

	condition,_ := v.VisitChild(cond).(*Value)
	if reflect.TypeOf(condition.vConst).Kind() != reflect.Bool {
		panic("need bool")
	}

	val := condition.vConst.(bool)
	if val && len(ctx.AllBlock()) >0 {
		v.VisitChild(ctx.Block(0))
	}else if len(ctx.AllBlock()) > 1 {
		v.VisitChild(ctx.Block(1))
	}
	return nil
}

func (v *VisitorImpl) VisitBlock(ctx *parser.BlockContext) interface{} {
	stmtList := ctx.StatementList()
	return v.VisitChild(stmtList)
}
