package main

import "fmt"

/*
	* 大写开头即public，小写开头即private
	* 所有定义的字段都必须使用到，否则编译报错

	* 关键字:

break    default      func    interface    select
case     defer        go      map          struct
chan     else         goto    package      switch
const    fallthrough  if      range        type
continue for          import  return       var

    var和const 变量和常量申明
    package和import
    func 用于定义函数和方法
    return 用于从函数返回
    defer 用于类似析构函数
    go 用于并发
    select 用于选择不同类型的通讯
    interface 用于定义接口
    struct 用于定义抽象数据类型
    break、case、continue、for、fallthrough、else、if、switch、goto、default
    chan 用于channel通讯
    type 用于声明自定义类型
    map 用于声明map类型数据
    range 用于读取slice、map、channel数据

*/
func main() {
	fmt.Printf("hello, world\n")
	fmt.Printf("你好, 世界!")
}
