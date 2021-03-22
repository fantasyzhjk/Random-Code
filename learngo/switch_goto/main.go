package main //switch,goto
import "fmt"

func main() {
	//switch：简化大量的判断
	n := 5
	//if写法
	if n == 1 {
		fmt.Println("大拇指")
	} else if n == 2 {
		fmt.Println("食指")
	} else if n == 3 {
		fmt.Println("中指")
	} else if n == 4 {
		fmt.Println("无名指")
	} else if n == 5 {
		fmt.Println("小拇指")
	} else {
		fmt.Println("无效")
	}
	//switch简化代码
	switch n := 3; n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效")
	}
	//变种
	switch {
	case n%2 == 2:
		fmt.Println("n为偶数")
	case n%2 == 1:
		fmt.Println("n为奇数")
	default:
		fmt.Println("n为0")
	}
	breakDemo()
}

func breakDemo() {
	//if语句
	flag := false
	for i := 0; ; i++ {
		for j := 'A'; j < 'K'; j++ {
			if j == 'C' && i == 5 { //当j为'C'且i为5时跳出内层for循环
				flag = true
				break
			}
			fmt.Printf("%v-%c\t", i, j)
		}
		if flag { //当flag为true时跳出外层for循环
			break
		}
	}
	fmt.Println()
	fmt.Println()
	//goto语句
	for i := 0; ; i++ {
		for j := 'A'; j < 'K'; j++ {
			if j == 'C' && i == 5 {
				goto breakPoint //跳到具体标签位置
			}
			fmt.Printf("%v-%c\t", i, j)
		}
	}
breakPoint: //指定标签
}
