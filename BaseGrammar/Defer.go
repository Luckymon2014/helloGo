package main

// defer语句: 当函数执行到最后时，defer语句会逆序执行，然后再返回
// 适用于打开资源的操作，不用频繁的使用close()
func main() {
	//file.Open("file")
	//defer file.Close()
	// ...

	// defer是逆序执行，即先进后出
	for i := 0; i < 5; i++ {
		defer println(i)
	}
}
