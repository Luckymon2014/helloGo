package main

import (
	"runtime"
	"time"
)

// 并发
// 协程: goroutine
// goroutine之间内存共享，不要通过共享来通信，要通过通信来共享
// 通过go关键字启动一个goroutine，其实就是一个普通函数

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		println(s)
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 在生产者的地方关闭channel
}

func fibonacci2(c, quit chan int) {
	x, y := 1, 1
	for {
		select { // 通过select选择channel，类似switch
		case c <- x:
			x, y = y, x+y
		case <-quit:
			println("quit")
			return
		case <-time.After(5 * time.Second):
			// 设置超时，防止阻塞
			println("timeout")
			break
		default:
			// channel阻塞时执行
			println("channel is blocked")
		}
	}
}

func main() {
	// 一个新的goroutine
	go say("world")
	// 当前的goroutine执行
	say("hello")

	println("==============================")
	a := []int{7, 2, 8, -9, 4, 0}
	// 通信机制channel
	// 必须通过make创建channel
	c1 := make(chan int)
	go sum(a[:len(a)/2], c1)
	go sum(a[len(a)/2:], c1)
	x, y := <-c1, <-c1 // receive from channel
	println(x, y, x+y)

	println("==============================")
	// Buffered Channel，创建一个可以存储n个缓冲元素的channel，防止阻塞，当数字位0时即无缓冲阻塞读写
	c2 := make(chan int, 2) // 当数字小于2时发生阻塞报错
	c2 <- 1
	c2 <- 2
	println(<-c2)
	println(<-c2)

	println("==============================")
	// 通过range操作缓存类型的channel
	c3 := make(chan int, 10)
	go fibonacci(cap(c3), c3)
	// for可以不断读取channel里的数据，直到channel被close
	for i := range c3 {
		println(i)
	}

	println("==============================")
	ok := <-c3 // 在消费的地方判断channel是否已关闭
	println(ok)
	// channel不需要频繁的close，除非确实没有数据需要发送了

	println("==============================")
	// 多个channel通过Select选择
	c4 := make(chan int, 10)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			println(<-c4)
			//println(<-c4) 测试阻塞
		}
		quit <- 0
	}()
	fibonacci2(c4, quit)

	println("==============================")
	// 退出当前执行的goroutine，但是defer函数还是会调用
	//runtime.Goexit()
	// 让出当前goroutine的执行权限，安排其他任务执行，并在下次从该位置恢复执行
	//runtime.Gosched()
	// CPU核数量
	println(runtime.NumCPU())
	// 当前任务总数
	println(runtime.NumGoroutine())
	// 设置可以并行计算的CPU核数的最大值，并返回之前的值
	println(runtime.GOMAXPROCS(1))
}
