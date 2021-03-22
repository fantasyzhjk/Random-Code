package main //运算符

import "fmt"

func main() {
	//特殊应用

	//not取反，原来是真就为假，原来是假就为真
	isMarried := true
	fmt.Println(isMarried)  //true
	fmt.Println(!isMarried) //false

	// 位运算，针对二进制数
	// 5的二进制表示：101
	// 2的二进制表示：010

	// &：按位与（两位均为1时才为1）
	fmt.Println(5 & 2) //101 & 010	1:0=0,0:1=0,1:0=0	000

	// |：按位或（两位中有一个为1时就为1）
	fmt.Println(5 | 2) //101 & 010	1:0=1,0:1=1,1:0=1	111

	// ^：按位异或（两位不一样则为1）
	fmt.Println(5 ^ 2) //101 & 010	1:0=1,0:1=1,1:0=1	111

	// <<：将二进制左移指定位置
	fmt.Println(5 << 1) //101 => 1010 = 10

	// >>：将二进制右移指定位置
	fmt.Println(5 >> 1) //101 => 010 = 2

	var m = int8(1)      //只能存8位[0000,0000]
	fmt.Println(m << 10) //超出变量范围
}
