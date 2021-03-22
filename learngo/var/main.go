package main //变量声明

import "fmt"

var ( //批量声明
	name    string //default=""
	age     int    //default=0
	isRight bool   //default=false
)

//函数外的每个语句必须以关键字开始（var，func...）
func main() {
	name = "真实理想"
	age = 16
	isRight = true
	fmt.Println(name)          //输出并换行
	fmt.Printf("年龄=%d\n", age) //参考C
	fmt.Print(isRight)         //就输出
	t1 := "哈哈哈自动识别"            //简短变量声明（用的最多）不能用在函数外
	fmt.Println(t1)
	var t2 string = "233" //声明变量时同时赋值
	fmt.Println(t2)
	var t3 = 2333 //类型推导（自动识别变量类型）
	fmt.Println(t3)
	//匿名变量“_”不占命名空间，多用于占位，表示忽略此值
}
