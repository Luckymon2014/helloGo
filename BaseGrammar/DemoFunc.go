package main

import "fmt"

// Go的函数可以返回两个参数
// 返回值可以缺省命名，但不推荐
func SumAndProduct(A, B int) (add int, multiplied int) {
	add, multiplied = A+B, A*B
	return
}

// 变参, arg是一个int的slice
func myFunc(arg ...int) {
	for _, n := range arg {
		println("And the number is: ", n)
	}
}

// Go的函数可以当参数传递
type funcType func(int, int) (int, int) // 需要先声明一个函数类型
func addResult(x int, y int, f funcType) int {
	r1, r2 := f(x, y)
	return r1 + r2
}

func main() {
	x, y := 3, 4
	r1, r2 := SumAndProduct(x, y)
	println(x, " + ", y, " = ", r1)
	println(x, " * ", y, " = ", r2)

	println("==============================")

	myFunc(1, 2, 3)

	println("==============================")
	result := addResult(x, y, SumAndProduct)
	fmt.Printf("(%d + %d) + (%d * %d) = %d", x, y, x, y, result)
}
