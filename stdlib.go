package main

import (
	"fmt"
	"time"
)

//标准库
type function func(args ...*Object)*Object
type StdFunction struct {
	FuncName string
	F function
}

type Module struct {
	Name string
	Desc string
	Funcs map[string]StdFunction
}

var StandardModule = make(map[string]*Module)
func Register(name string,module *Module){
	StandardModule[name] = module
}

var fmtModule = &Module{
	Name: "io",
	Desc: "格式化标准库",
	Funcs: map[string]StdFunction{
		//sprintf
		"sprintf":{
			F: func(args ...*Object) *Object {
				if len(args) < 2 {

				}
				format,_ := args[0].vConst.(string)

				arg := []interface{}{}
				for _, a := range args[1:] {
					arg = append(arg,a.vConst)
				}
				rets := fmt.Sprintf(format,arg...)
				return &Object{
					vConst: rets,
					rlType: RightValue,
				}
			},
		},

		//printf
		"print":{
			F: func(args ...*Object) *Object {
				//format,_ := args[0].vConst.(string)

				arg := []interface{}{}
				for _, a := range args[0:] {
					arg = append(arg,a.vConst)
				}
				fmt.Println(arg...)
				return &Object{
					vConst: "",
					rlType: RightValue,
				}
			},
		},
	},
}

var timeModule = &Module{
	Name: "time",
	Desc: "系统时间模块",
	Funcs: map[string]StdFunction{
		//time.now
		"now":{
			F: func(args ...*Object) *Object {
				ts := time.Now().Unix()
				return &Object{
					vConst: ts,
					rlType: RightValue,
				}
			},
		},
	},
}
func init() {
	Register("io",fmtModule)
	Register("time",timeModule)
}