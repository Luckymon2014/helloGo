package main

// Go没有异常机制，不能抛出异常
/*
Panic
是一个内建函数，可以中断原有的控制流程，进入一个panic状态中。
当函数F调用panic，函数F的执行被中断，但是F中的延迟函数会正常执行，然后F返回到调用它的地方。
在调用的地方，F的行为就像调用了panic。
这一过程继续向上，直到发生panic的goroutine中所有调用的函数返回，此时程序退出。
panic可以直接调用panic产生。也可以由运行时错误产生，例如访问越界的数组。

Recover
是一个内建的函数，可以让进入panic状态的goroutine恢复过来。
recover仅在延迟函数中有效。
在正常的执行过程中，调用recover会返回nil，并且没有其它任何效果。
如果当前的goroutine陷入panic状态，调用recover可以捕获到panic的输入值，并且恢复正常的执行。
*/

// 检查入参函数执行时是否发生panic
func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
	return
}

func testPanic() {
	test := ""
	if test == "" {
		panic("test is empty")
	}
}

func main() {
	println(throwsPanic(testPanic))
}
