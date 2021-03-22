package main //if/for语句

import "fmt"

func main() {
	age := 9
	if age >= 18 {
		fmt.Println("澳门首家线上赌场开业了！")
	} else {
		fmt.Println("该写作业了")
	}
	//else if
	if nianLin := 40; nianLin >= 35 { //nianLin变量此时只在if判断语句中，无法从外部调用，可以减少程序内存占用。
		fmt.Println("人到中年")
	} else if nianLin >= 18 {
		fmt.Println("还在青年")
	} else {
		fmt.Println("好好上学！")
	}
	//fmt.Println("在此处找不到'nianLin'", nianLin)

	for i := 0; i <= 9; i++ { //基本格式（打印0-9）
		fmt.Println(i)
	}

	var i = 5
	for ; i < 10; i++ { //变种一
		fmt.Println(i)
	}

	var i1 = 5
	for i1 < 10 { //变种二
		fmt.Println(i1)
		i1++
	}

	// for { //变种三（无限循环）
	// 	fmt.Println(i)
	// }

	s := "你好,世界！"
	for i, v := range s { //变种四（range循环）
		fmt.Printf("%d %c\n", i, v)
	}
}
