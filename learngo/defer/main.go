package main //defer
import "fmt"

//defer执行时机：
//1.返回值赋值
//2.defer语句
//3.返回返回值

func main() {
	deferDemo()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

func deferDemo() { // defer多用于函数结束之前释放资源（文件句柄，数据库连接，socket连接等）
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿") //运用defer吧其后面的语句延迟到函数即将返回时再执行
	defer fmt.Println("呵呵呵") //倒序执行（顺序：呼呼呼 >> 呵呵呵 >> 嘿嘿嘿）
	defer fmt.Println("呼呼呼")
	fmt.Println("over")

}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x //先把x=5赋给返回值，然后再执行x++，故返回5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //先把5赋给返回值x，然后再执行x++，故返回6
}

func f3() (x int) {
	x = 5
	defer func() {
		x++
	}()
	return x //先把x=5赋给返回值y，然后再执行x++，故返回5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 //先把5赋给返回值x，然后再执行函数中的x++（改变的是副本），故返回5
}
