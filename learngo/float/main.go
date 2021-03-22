package main //浮点型变量

import "fmt"

func main() {
	//math.MaxFloat32 //float32最大值
	f1 := 1.4241
	fmt.Printf("%T\n", f1) //默认golang中的小数都是float64
	f2 := float32(3.14)    //表示声明float32类型
	fmt.Printf("%T\n", f2)
	//f1 = f2
	//两种浮点类型不能互相赋值
}
