package main

import "fmt"
import "strconv"

// interface是一组抽象方法的集合，必须有其它类型实现
// 任意类型都自带一个空interface{}，因此空interface可以存储任意类型的值，类似C的void*类型
// interface也可以匿名继承

type Human struct {
	name  string
	age   int
	phone string
}

type Employee struct {
	Human   //匿名字段
	company string
	money   float32
}

//Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println(h.name, "is singing: La la la la...", lyrics)
}

//Employee重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Interface Men被Human和Employee实现
// 因为这两个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

// fmt定义了一个Stringer的interface：
//type Stringer interface {
//	String() string
//}
// 因此只要实现了String()方法，就可以作为入参被fmt调用
func (h Human) String() string {
	return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years -  ✆ " + h.phone + "❱"
}

func main() {
	sam := Human{"Sam", 36, "444-222-XXX"}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i
	var i Men

	i = tom
	i.SayHi()
	i.Sing("Born to be wild")

	fmt.Println("==============================")

	// 不同类型的元素，但是他们实现了interface同一个接口
	x := make([]Men, 2)
	x[0], x[1] = sam, tom
	for _, value := range x {
		value.SayHi()
	}

	fmt.Println("==============================")
	fmt.Println("Sam: ", sam)
	fmt.Println("Tom: ", tom.String())

	fmt.Println("==============================")
	// 获取空interface里存的值的类型
	type Element interface{}
	type List []Element
	list := make(List, 3)
	list[0] = 1
	list[1] = "hello"
	list[2] = sam
	for index, element := range list {
		// comma-ok
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and value is %d\n", index, value)
		}
		if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and value is %s\n", index, value)
		}
		if value, ok := element.(Human); ok {
			fmt.Println("list["+strconv.Itoa(index)+"] is a Human and value is", value)
		}
		// switch
		// element.(type)只能在switch中使用
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and value is %s\n", index, value)
		case Human:
			fmt.Println("list["+strconv.Itoa(index)+"] is a Human and value is", value)
		}
	}
}
