package main

import (
	"fmt"
	"reflect"
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

var typeModule = &Module{
	Name: "std",
	Desc: "类型模块",
	Funcs: map[string]StdFunction{
		//time.now
		"len":{
			F: func(args ...*Object) *Object {
				if len(args) != 1 {
					panic("invalid type.len usage")
				}

				v := args[0].vConst
				var vv int64
				switch reflect.TypeOf(v).Kind() {
				case reflect.String:
					vv = int64(len(v.(string)))
				case reflect.Slice:
					vv = int64(len(v.([]*Object)))
				default:
					panic("unsupport type.len for type:%v")
				}

				return &Object{
					vConst: vv,
					rlType: RightValue,
				}
			},
		},
		"typeof":{
			F: func(args ...*Object) *Object {
				if len(args) != 1 {
					panic("invalid std.type usage")
				}

				v := args[0].vConst
				var vv string
				switch reflect.TypeOf(v).Kind() {
				case reflect.String:
					vv = "str"
				case reflect.Int64,reflect.Int:
					vv = "integer"
				case reflect.Slice:
					vv = "list"
				default:
					panic("undefined type")
				}

				return &Object{
					vConst: vv,
					rlType: RightValue,
				}
			},
		},
	},
}


func init() {
	Register("io",fmtModule)
	Register("time",timeModule)
	Register("std",typeModule)
}