package main //rune切片，byte类型

import (
	"fmt"
	"math"
)

func main() {
	//byte	uint8的别名
	//rune	int32的别名
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	s := "the世界"
	for i := 0; i < len(s); i++ { //len(s)求字符串的长度
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%v(%c) ", r, r)
	}

	//字符串修改
	fmt.Println()
	s1 := "白酱ovo"
	fmt.Println("Before:", s1)
	s2 := []rune(s1) //把字符串强制转换成rune切片:['白','酱','o','v','o']
	s2[0] = '黑'
	fmt.Println("afterRune:", s2)
	fmt.Println("After:", string(s2))
	//c := "文本"|字符串(string)
	//c := '文本'|字符rune(int32)
	//c := byte('H')|字符(uint8)

	//类型转换
	i := 10
	var f float64
	f = float64(i)
	fmt.Printf("i=%T\n", i)
	fmt.Printf("f=%T\n", f)

	//计算直角三角形的斜边长demo
	sqrtDemo()
}

func sqrtDemo() {
	var a, b = 6, 7
	var c float64
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = math.Sqrt(float64(a*a + b*b))
	fmt.Println(int(c))
}
