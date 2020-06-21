package main

import (
	"antlr_test/parserV1"
	"github.com/emirpasic/gods/stacks/linkedliststack"
	"fmt"
	"strconv"
)

type ListenImpl struct {
	parserV1.BaseCalcListener
	stack *linkedliststack.Stack
}

func NewListenImpl()*ListenImpl{
	var listen = &ListenImpl{}
	listen.stack = linkedliststack.New()

	return listen
}

func (s *ListenImpl)ExitNumber(ctx *parserV1.NumberContext) {
	n,_ := strconv.Atoi(ctx.NUMBER().GetText())
	s.stack.Push(n)
	s.PrintStack()
}

func (s *ListenImpl) ExitAddSub(ctx *parserV1.AddSubContext) {
	right,_ := s.stack.Pop()
	left,_ := s.stack.Pop()

	ret := right.(int) + left.(int)
	s.stack.Push(ret)
	s.PrintStack()
}

func (s *ListenImpl) ExitMulDiv(ctx *parserV1.MulDivContext) {
	right,_ := s.stack.Pop()
	left,_ := s.stack.Pop()

	ret := right.(int) * left.(int)
	s.stack.Push(ret)
	s.PrintStack()
}

func (s *ListenImpl)PrintStack(){
	fmt.Println(s.stack.String())
}