package main

import (
	"fmt"
	"reflect"
)

type Value struct {
	rlType int //值类型: 左值、右值
	vConst interface{} //如果是左值，则从当前环境中取值，如果是右值，则从这里取值
	symbol string
	vType int //数据类型
}

//一些常量池
var (
	valueTrue = &Value{
		rlType:RightValue,
		vConst: true,
	}

	valueFalse = &Value{
		rlType:RightValue,
		vConst: false,
	}
)

func (v *Value)String() string {
	return fmt.Sprintf("rtype:%d,vconst:%+v,vtype:%d",v.rlType,v.vConst,v.vType)
}

func (v *Value)Add(v1 *Value)*Value {
	value := &Value{}
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

func (v *Value)Sub(v1 *Value)*Value {
	value := &Value{}
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

func (v *Value)Mul(v1 *Value)*Value {
	value := &Value{}
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

func (v *Value)Div(v1 *Value)*Value {
	value := &Value{}
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
func (v *Value)Less(v1 *Value)*Value {
	value := &Value{}
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
func (v *Value)LessEq(v1 *Value)*Value {
	value := &Value{}
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
func (v *Value)Greater(v1 *Value)*Value {
	value := &Value{}
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
func (v *Value)GreaterEq(v1 *Value)*Value {
	value := &Value{}
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
func (v *Value)Equal(v1 *Value)*Value {
	value := &Value{}
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
func (v *Value)NotEqual(v1 *Value)*Value {
	value := &Value{}
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

func (v *Value)And(v1 *Value)*Value {
	value := &Value{}
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

func (v *Value)Or(v1 *Value)*Value {
	value := &Value{}
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