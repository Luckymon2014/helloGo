package main

import "fmt"

func add1(a int) int {
	a = a + 1
	return a
}

// 入参为指针
func add2(a *int) int {
	*a = *a + 1
	return *a
}

func main() {
	x := 3

	fmt.Println("old x = ", x)

	// Go调用函数传参时，实际是传了这个参数的一个副本，实际这个参数的值并不会被修改
	x1 := add1(x)
	fmt.Println("x + 1 = ", x1)
	fmt.Println("new x = ", x, " after add1")

	// 传x值的实际内存地址
	x2 := add2(&x)
	fmt.Println("x + 1 = ", x2)
	fmt.Println("new x = ", x, " after add2")
}
