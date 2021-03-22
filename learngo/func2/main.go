package main // func2:匿名函数，闭包
import "fmt"

func main() {
	f0 := func(x, y string) { // 函数内部无法声明一个有名字的函数
		fmt.Println(x + y)
	}
	f0("Hello ", "World!")

	func(x, y string) { // 立即执行函数
		fmt.Println(x + y)
	}("Hello ", "World!")
	f := adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	ff := adder()
	fmt.Println(ff(40)) //40
	fmt.Println(ff(50)) //90
	f1(f3(f2, 1, 2))    //3
}

//闭包示例：
func f1(f func()) { //外部接口
	f()
}
func f2(x, y int) { //内部函数
	fmt.Println(x + y)
}
func f3(f func(x, y int), x, y int) func() { //闭包
	return func() {
		f(x, y)
	}
}

//闭包示例2 ：
func adder() func(int) int { // 闭包=函数+引用环境
	var x int                //外部应用环境
	return func(y int) int { // 内部函数
		x += y
		return x
	}
}
