package main //递归
import "fmt"

// 函数自己调用自己

// 计算n的阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

// 上台阶
func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

func main() {
	fmt.Println(f(5))
	fmt.Println(taijie(5))
}
