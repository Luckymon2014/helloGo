package main

func main() {

	// Go的while和for一个写法
	sum := 1
	for sum < 10 {
		sum += sum
		println(sum)
	}

	println("==============================")

	// 用range来遍历字典
	numbers := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, value := range numbers {
		println(key, ": ", value)
	}

}
