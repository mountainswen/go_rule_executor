package main

import "fmt"

//非线程安全
//原则上不同执行流不应该共用一个environment
type Enviroment struct {
	frames []*Frame
	cur int
}

type Frame struct {
	local map[string]*Object
}

func MakeEnv()*Enviroment{
	return &Enviroment{
		cur: -1,
	}
}

func (e *Enviroment)SetSymbol(name string,object *Object){
	if e.cur == -1 {
		e.Extend()
	}

	frame := e.frames[e.cur]
	frame.local[name] = object
}

func (e *Enviroment)GetSymbol(name string)(*Object,bool){
	if e.cur == -1 {
		return nil,false
	}

	if !(e.cur < len(e.frames)){
		return nil,false
	}

    //从内层作用域往外层找
	for i := e.cur; i >=0; i-- {
		frame := e.frames[i]
		if frame == nil {
			continue
		}

		if symbol, ok := frame.local[name]; ok {
			return symbol,true
		}
	}

	return nil, false
}

func (e *Enviroment)Shrink(){
	e.cur--
}

func (e *Enviroment)Extend(){
	e.cur++

	newFrame := &Frame{local: make(map[string]*Object)}
	e.frames = append(e.frames,newFrame)
}

func (e *Enviroment)String()string{
	str := ""
	for _, frame := range e.frames {
		str += "===============================\n"
		str += fmt.Sprintf("%+v\n",frame.local)
	}

	return str
}