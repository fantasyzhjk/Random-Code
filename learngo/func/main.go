package main //函数
import "fmt"

//函数的定义
//func 函数名(参数)(返回值){
//    函数体
//}

//函数能让你的代码更简洁清晰

func main() {
	fmt.Println(funcTion(1, 2)) //调用函数
	funcTion1()
	funcTion2(1, 2)
	funcTion3()
	_, s := funcTion4()
	fmt.Println(s)
	funcTion5("abcd", 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	f6(funcTion3) //函数类型必须匹配
	fmt.Printf("%T\n", f7())
}

//命名返回值就相当于在函数中已经声明了一个变量作为返回值
func funcTion(x int, y int) (ret int) {
	ret = x + y
	return //这里return后可省略
}

//函数变种（无参数无返回值）
func funcTion1() {
	fmt.Println("Hello！")
}

//函数变种（有参数无返回值）
func funcTion2(x int, y int) { //参数类型简写(x int, y int) >> (x, y int)
	fmt.Println(x + y)
}

//函数变种（无参数有返回值）//参数命名可选
func funcTion3() int {
	ret := 3
	return ret
}

//多个返回值
func funcTion4() (int, string) {
	return 114514, "一九一九八一零"
}

//可变参数，必须放在函数参数的最后
func funcTion5(x string, y ...int) {
	fmt.Println(x, y) //这里的y为切片
}

//golang中没有默认参数的概念
// 函数也可以作为参数的类型
func f6(x func() int) {
	fmt.Println(x())
}

// 函数也可以作为返回值
func f7() func(int, int) int {
	return ff
}
func ff(x, y int) int {
	return x + y
}
