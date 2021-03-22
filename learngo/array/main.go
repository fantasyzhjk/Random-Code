package main //数组
import "fmt"

//存放元素的容器
//必须指定存放的元素的类型和容量（长度）
//数组的长度是数组类型的一部分

func main() {
	var a [3]int  // [0,0,0]
	var b [5]int  // [0,0,0,0,0]
	var c [2]bool // [false,false]
	//a = b	//不可以这样做，此时a与b是两种类型，不能互通
	fmt.Printf("a=%T b=%T c=%T\n", a, b, c)

	//数组的初始化(如果不初始化，默认都为零指)
	fmt.Println(a, c)
	a = [3]int{1, 2, 5}
	c = [2]bool{true, false}
	fmt.Println(a, c)
	d := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //根据初始值自动推断数组长度
	fmt.Println(d)
	e := [4]int{0, 1} // [0,1,0,0]
	fmt.Println(e)
	f := [5]int{0: 1, 4: 5} // 根据索引初始化
	fmt.Println(f)

	//数组的遍历
	cities := [...]string{"北京", "上海", "扬州"} //索引：0-2 cities[0],cities[1],cities[2]
	//根据索引遍历
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
	//for range遍历
	for _, v := range cities {
		fmt.Println(v)
	}
	fmt.Println()

	//多维数组
	//注意： 多维数组只有第一层可以使用...来让编译器推导数组长度。例如：
	//支持的写法
	//a := [...][2]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}
	//不支持多维数组的内层使用...
	//b := [3][...]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}
	g1 := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(g1)
	g := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(g)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(g[2][1]) //支持索引取值:重庆
	//多维数组的遍历
	for _, v1 := range g {
		fmt.Print(v1)
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}
	//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	homework()
}

func homework() {
	i := 0
	s := [...]int{1, 3, 5, 7, 8}
	for _, num := range s {
		i = i + num
	}
	fmt.Println(i)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i]+s[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}
