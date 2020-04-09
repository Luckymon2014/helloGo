package main

import "math"

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// method: func (r ReceiverType) funcName(parameters) (results)
// ReceiverType表示接收这个函数的Struct是谁

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

// method操作的也是副本，使用指针作为Receiver修改实际值
func (c *Circle) setRadius(r float64) {
	c.radius = r // 实际是*c，Go自动帮忙转换了
}

func main() {
	println(Rectangle{3, 4}.area())
	c := Circle{1}
	println(c.area())
	c.setRadius(2) // 实际是(&c)，Go自动帮忙转换了
	println(c.area())
}
