package main

import "fmt"

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
	Name: "fmt",
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
		"printf":{

		},
	},
}

func init(){
	Register("fmt",fmtModule)
}