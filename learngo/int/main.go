package main //整型变量

import "fmt"

func main() {
	i1 := 256              //十进制
	fmt.Printf("%d\n", i1) //十进制输出
	fmt.Printf("%b\n", i1) //二进制输出
	fmt.Printf("%o\n", i1) //八进制输出
	fmt.Printf("%x\n", i1) //十六进制输出
	i2 := 0777             //八进制（0开头）
	fmt.Printf("%d\n", i2)
	i3 := 0xFF //十六进制（0x开头）
	fmt.Printf("%d\n", i3)
	i4 := int8(9)
	fmt.Printf("%T\n", i4) //查看变量类型

}
