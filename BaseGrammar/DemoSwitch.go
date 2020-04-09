package main

func main() {

	integer := 2
	switch integer {
	case 1:
		println("integer is 1")
	case 2:
		println("integer is 2")
		// Go的switch默认自带break，使用fallthrough强制执行后面的一个case
		fallthrough
	case 3:
		if integer == 3 {
			println("integer is 3")
		} else {
			println("integer is not 3 but fallthrough to here")
		}
	default:
		println("default case")
	}

}
