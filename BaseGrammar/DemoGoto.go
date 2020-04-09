package main

func main() {
	i := 0
	// Goto需要指明标签
	Here:
		if i < 10 {
			print(i)
			i++
			goto Here
		}
}
