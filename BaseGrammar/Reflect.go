package main

import "reflect"

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	println("type: ", v.Type())
	println("kind is float64: ", v.Kind() == reflect.Float64)
	println("value: ", v.Float())
	// 反射的值必须可修改
	// 错误写法: v.SetFloat(7.1)
	p := reflect.ValueOf(&x).Elem()
	p.SetFloat(7.1)
	println("new value: ", x)
}
