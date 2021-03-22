package main //指针
import "fmt"

func main() {
	// 1.'&'取地址
	n := 18
	p := &n
	fmt.Printf("%v %T\n", &n, p) //*int:int类型的指针
	// 2.'*'取地址对应的值
	m := *p
	fmt.Printf("type of m:%T\n", m)
	fmt.Println(m)
	// 取地址操作符'&'和取值操作符'*'是一对互补操作符，'&'取出地址，'*'根据地址取出地址指向的值。

	// new和make
	//var a *int	//这里只是声明了一个指针变量a但是没有初始化，并没有内存空间，无法进行赋值
	a := new(int)   //用new函数申请一个内存地址(返回*int)
	fmt.Println(*a) //使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。
	*a = 100        //通过内存地址找到变量并赋值
	fmt.Println(*a)

	//make只用于slice、map以及chan的内存创建,而且它返回是这三个类型本身,而不是他们的指针类型(返回int)
	e := make(map[string]int, 1)
	e["score"] = 100
	fmt.Println(e) //map[score:100]
}
