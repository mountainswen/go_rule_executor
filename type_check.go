package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Object struct {
	rlType int //值类型: 左值、右值
	vConst interface{} //如果是左值，则从当前环境中取值，如果是右值，则从这里取值
	symbol string
	vType int //数据类型
}

//一些常量池
var (
	valueTrue = &Object{
		rlType:RightValue,
		vConst: true,
	}

	valueFalse = &Object{
		rlType:RightValue,
		vConst: false,
	}
)

func (v *Object)String() string {
	return fmt.Sprintf("rtype:%d,vconst:%+v,vtype:%d",v.rlType,v.vConst,v.vType)
}

func (v *Object)Add(v1 *Object)*Object {
	if v.vConst == nil {
		return &Object{
			vConst: "undefine variable:"+v.symbol,
		}
	}
	if v1.vConst == nil {
		return &Object{
			vConst: "undefine variable:"+v1.symbol,
		}
	}

	value := &Object{}
	value.rlType = RightValue
	switch reflect.TypeOf(v.vConst).Kind() {
	case reflect.String:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
			value.vConst = v.vConst.(string) + v1.vConst.(string)
			return value
		case reflect.Int64:

		case reflect.Float64:

		}
	case reflect.Int64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
		case reflect.Int64:
			value.vConst = v.vConst.(int64) + v1.vConst.(int64)
			return value
		case reflect.Float64:
			value.vConst = float64(v.vConst.(int64)) + v1.vConst.(float64)
			return value
		}

	case reflect.Float64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:

		case reflect.Int64:
			value.vConst = v.vConst.(float64) + float64(v1.vConst.(int64))
			return value

		case reflect.Float64:
			value.vConst = v.vConst.(float64) + v1.vConst.(float64)
			return value
		}
	}

	return value
}

func (v *Object)Sub(v1 *Object)*Object {
	value := &Object{}
	value.rlType = RightValue
	switch reflect.TypeOf(v.vConst).Kind() {
	case reflect.String:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
		case reflect.Int64:
		case reflect.Float64:

		}
	case reflect.Int64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
		case reflect.Int64:
			value.vConst = v.vConst.(int64) - v1.vConst.(int64)
			return value
		case reflect.Float64:
			value.vConst = float64(v.vConst.(int64)) - v1.vConst.(float64)
			return value
		}

	case reflect.Float64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:

		case reflect.Int64:
			value.vConst = v.vConst.(float64) - float64(v1.vConst.(int64))
			return value

		case reflect.Float64:
			value.vConst = v.vConst.(float64) - v1.vConst.(float64)
			return value
		}
	}

	return value
}

func (v *Object)Mul(v1 *Object)*Object {
	value := &Object{}
	value.rlType = RightValue
	switch reflect.TypeOf(v.vConst).Kind() {
	case reflect.String:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
		case reflect.Int64:
		case reflect.Float64:

		}
	case reflect.Int64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
		case reflect.Int64:
			value.vConst = v.vConst.(int64) * v1.vConst.(int64)
			return value
		case reflect.Float64:
			value.vConst = float64(v.vConst.(int64)) * v1.vConst.(float64)
			return value
		}

	case reflect.Float64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:

		case reflect.Int64:
			value.vConst = v.vConst.(float64) * float64(v1.vConst.(int64))
			return value

		case reflect.Float64:
			value.vConst = v.vConst.(float64) * v1.vConst.(float64)
			return value
		}
	}

	return value
}

func (v *Object)Div(v1 *Object)*Object {
	value := &Object{}
	value.rlType = RightValue
	switch reflect.TypeOf(v.vConst).Kind() {
	case reflect.String:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:

		case reflect.Int64:

		case reflect.Float64:

		}
	case reflect.Int64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
		case reflect.Int64:
			value.vConst = v.vConst.(int64) / v1.vConst.(int64)
			return value
		case reflect.Float64:
			value.vConst = float64(v.vConst.(int64)) / v1.vConst.(float64)
			return value
		}

	case reflect.Float64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:

		case reflect.Int64:
			value.vConst = v.vConst.(float64) / float64(v1.vConst.(int64))
			return value

		case reflect.Float64:
			value.vConst = v.vConst.(float64) / v1.vConst.(float64)
			return value
		}
	}

	return value
}

//v < v1 ?
func (v *Object)Less(v1 *Object)*Object {
	value := &Object{}
	value.rlType = RightValue
	switch reflect.TypeOf(v.vConst).Kind() {
	case reflect.String:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
			value.vConst = v.vConst.(string) <  v1.vConst.(string)
			return value

		case reflect.Int64:

		case reflect.Float64:

		}
	case reflect.Int64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
		case reflect.Int64:
			value.vConst = v.vConst.(int64) < v1.vConst.(int64)
			return value
		case reflect.Float64:
			value.vConst = float64(v.vConst.(int64)) < v1.vConst.(float64)
			return value
		}

	case reflect.Float64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:

		case reflect.Int64:
			value.vConst = v.vConst.(float64) < float64(v1.vConst.(int64))
			return value

		case reflect.Float64:
			value.vConst = v.vConst.(float64) < v1.vConst.(float64)
			return value
		default:

		}
	default:
	}
	return value
}

//v <= v1
func (v *Object)LessEq(v1 *Object)*Object {
	value := &Object{}
	value.rlType = RightValue
	switch reflect.TypeOf(v.vConst).Kind() {
	case reflect.String:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
			value.vConst = v.vConst.(string) <=  v1.vConst.(string)
			return value

		case reflect.Int64:

		case reflect.Float64:

		}
	case reflect.Int64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
		case reflect.Int64:
			value.vConst = v.vConst.(int64) <= v1.vConst.(int64)
			return value
		case reflect.Float64:
			value.vConst = float64(v.vConst.(int64)) <= v1.vConst.(float64)
			return value
		}

	case reflect.Float64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:

		case reflect.Int64:
			value.vConst = v.vConst.(float64) <= float64(v1.vConst.(int64))
			return value

		case reflect.Float64:
			value.vConst = v.vConst.(float64) <= v1.vConst.(float64)
			return value
		default:

		}
	default:
	}
	return value
}

//v > v1
func (v *Object)Greater(v1 *Object)*Object {
	value := &Object{}
	value.rlType = RightValue
	switch reflect.TypeOf(v.vConst).Kind() {
	case reflect.String:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
			value.vConst = v.vConst.(string) >  v1.vConst.(string)
			return value

		case reflect.Int64:

		case reflect.Float64:

		}
	case reflect.Int64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
		case reflect.Int64:
			value.vConst = v.vConst.(int64) > v1.vConst.(int64)
			return value
		case reflect.Float64:
			value.vConst = float64(v.vConst.(int64)) > v1.vConst.(float64)
			return value
		}

	case reflect.Float64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:

		case reflect.Int64:
			value.vConst = v.vConst.(float64) > float64(v1.vConst.(int64))
			return value

		case reflect.Float64:
			value.vConst = v.vConst.(float64) > v1.vConst.(float64)
			return value
		default:

		}
	default:
	}
	return value
}

//v >= v1
func (v *Object)GreaterEq(v1 *Object)*Object {
	value := &Object{}
	value.rlType = RightValue
	switch reflect.TypeOf(v.vConst).Kind() {
	case reflect.String:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
			value.vConst = v.vConst.(string) >=  v1.vConst.(string)
			return value

		case reflect.Int64:

		case reflect.Float64:

		}
	case reflect.Int64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
		case reflect.Int64:
			value.vConst = v.vConst.(int64) >= v1.vConst.(int64)
			return value
		case reflect.Float64:
			value.vConst = float64(v.vConst.(int64)) >= v1.vConst.(float64)
			return value
		}

	case reflect.Float64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:

		case reflect.Int64:
			value.vConst = v.vConst.(float64) >= float64(v1.vConst.(int64))
			return value

		case reflect.Float64:
			value.vConst = v.vConst.(float64) >= v1.vConst.(float64)
			return value
		default:

		}
	default:
	}
	return value
}

//v == v1
func (v *Object)Equal(v1 *Object)*Object {
	value := &Object{}
	value.rlType = RightValue
	switch reflect.TypeOf(v.vConst).Kind() {
	case reflect.String:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
			value.vConst = v.vConst.(string) ==  v1.vConst.(string)
			return value

		case reflect.Int64:

		case reflect.Float64:

		}
	case reflect.Int64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
		case reflect.Int64:
			value.vConst = v.vConst.(int64) == v1.vConst.(int64)
			return value
		case reflect.Float64:
			value.vConst = float64(v.vConst.(int64)) == v1.vConst.(float64)
			return value
		}

	case reflect.Float64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:

		case reflect.Int64:
			value.vConst = v.vConst.(float64) == float64(v1.vConst.(int64))
			return value

		case reflect.Float64:
			value.vConst = v.vConst.(float64) == v1.vConst.(float64)
			return value
		default:

	case reflect.Bool:
		switch reflect.TypeOf(v1.vConst).Kind() {
		case reflect.Bool:
			value.vConst = v.vConst.(bool) == v1.vConst.(bool)
			return value
		default:

		}
		}
	default:
	}
	return value
}

//v != v1
func (v *Object)NotEqual(v1 *Object)*Object {
	value := &Object{}
	value.rlType = RightValue
	switch reflect.TypeOf(v.vConst).Kind() {
	case reflect.String:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
			value.vConst = v.vConst.(string) !=  v1.vConst.(string)
			return value

		case reflect.Int64:

		case reflect.Float64:

		}
	case reflect.Int64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:
		case reflect.Int64:
			value.vConst = v.vConst.(int64) != v1.vConst.(int64)
			return value
		case reflect.Float64:
			value.vConst = float64(v.vConst.(int64)) != v1.vConst.(float64)
			return value
		}

	case reflect.Float64:
		switch  reflect.TypeOf(v1.vConst).Kind(){
		case reflect.String:

		case reflect.Int64:
			value.vConst = v.vConst.(float64) != float64(v1.vConst.(int64))
			return value

		case reflect.Float64:
			value.vConst = v.vConst.(float64) != v1.vConst.(float64)
			return value
		default:

		}
	case reflect.Bool:
		switch reflect.TypeOf(v1.vConst).Kind() {
		case reflect.Bool:
			value.vConst = v.vConst.(bool) != v1.vConst.(bool)
			return value

		default:
		}
	}
	return value
}

func (v *Object)And(v1 *Object)*Object {
	value := &Object{}
	value.rlType = RightValue
	switch reflect.TypeOf(v.vConst).Kind() {
	case reflect.String:
	case reflect.Int64:
	case reflect.Float64:

	case reflect.Bool:
		switch reflect.TypeOf(v1.vConst).Kind() {
		case reflect.Bool:
			value.vConst = v.vConst.(bool) && v1.vConst.(bool)
			return value
		default:
		}
	}
	return value
}

func (v *Object)Or(v1 *Object)*Object {
	value := &Object{}
	value.rlType = RightValue
	switch reflect.TypeOf(v.vConst).Kind() {
	case reflect.String:
	case reflect.Int64:
	case reflect.Float64:

	case reflect.Bool:
		switch reflect.TypeOf(v1.vConst).Kind() {
		case reflect.Bool:
			value.vConst = v.vConst.(bool) || v1.vConst.(bool)
			return value
		default:
		}
	}
	return value
}

func TriggerMethod(mmethod string,args interface{})*Object{
	var module string
	var function string
	arr := strings.Split(mmethod,".")
	if len(arr) != 2 {
		panic("invalid method call:"+mmethod)
	}

	module = arr[0]
	function = arr[1]

	realArgs, ok := args.([]*Object)
	if !ok {
		realArgs = []*Object{}
	}

	if mod, ok := StandardModule[module];ok {
		if f, ok := mod.Funcs[function]; ok {
			return f.F(realArgs...)
		}
	}
	//panic("undefine method")
	return &Object{
		vConst: "undefined method:"+mmethod,
	}
}

func (o1 *Object)Index(idx *Object)*Object {
	if o1arr,ok := o1.vConst.([]*Object);ok {
		if index, ok := idx.vConst.(int64); ok && len(o1arr) - 1 >= int(index) {
			return o1arr[int(index)] //不支持修改数组里的值
		} else if len(o1arr) - 1 < int(index){
			panic("index out of range")
		} else {
			panic("undefined index operator")
		}
	}else {
		panic("undefine index operator")
	}
	return nil
}