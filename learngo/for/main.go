package main

import "fmt"

func main() {
	for i := 0; ; i++ {
		if i == 5 {
			break //当i=5时跳出for循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue //当i=5时继续下一次循环，跳过此次for循环
		}
		fmt.Println(i)
		//i++
	}
	fmt.Println("over")
}
