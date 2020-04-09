package main

// 自定义类型可以作用于基础类型，相当于起了个别名
type ages int

// 用struct声明一个自定义的类型
type Person struct {
	name string
	age ages // 这里用了自定义类型，其实就是int
	phone string
}

type Student struct {
	Person // 匿名字段，默认继承Person的所有字段
	int //基础类型也可以用作匿名字段
	phone string // 如果有冲突字段，优先取外层的
	grade string
}

// method也可以继承
func (p Person) sayHello() {
	println(p.name, "say hello")
}

// 同样，优先使用外层的method，即Override
func (s Student) sayHello() {
	println(s.name, "say hello, I'm a student")
}

func main() {

	// 使用struct
	var p1 Person
	p1.name = "Shane"
	p1.age = 28
	p2 := Person{"Tom", 25, ""}
	p3 := new(Person) // new出来的是*person指针
	p3.name, p3.age = "Jack", 26
	println(p1.name, "'s age is", p1.age)
	println(p2.name, "'s age is", p2.age)
	println(p3.name, "'s age is", p3.age)

	println("==============================")

	s1 := Student{Person{"Mark", 27, "111"}, 0, "222", "Senior"}
	// 使用s1.Person.name或者s1.name均可
	println(s1.Person.name, "'s age is", s1.age, "and grade is", s1.grade)
	println(s1.phone, s1.Person.phone)

	println("==============================")
	s1.Person.sayHello()
	s1.sayHello()
}
