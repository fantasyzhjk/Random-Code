package main //切片（slice）

import (
	"fmt"
	"sort"
)

func main() {
	//声明切片类型
	var s []int     //定义了一个存放int类型的切片
	var s1 []string //定义了一个存放string类型的切片
	fmt.Println(s, s1)
	fmt.Println(s == nil) //检测切片有没有值，此处输出true

	//切片初始化
	s = []int{1, 2, 3, 4}
	s1 = []string{"扬州", "常州", "上海"}
	fmt.Println(s, s1)
	fmt.Println(s == nil) //此处输出false
	//fmt.Println(s == s1)   //切片是引用类型，不支持直接比较，只能和nil比较

	//长度和容量
	//切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。
	fmt.Printf("len(s):%d cap(s):%d\n", len(s), cap(s))
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))

	//由数组得到切片
	a := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a[1:4] //基于数组a切割，左包含右不包含[3 5 7]
	s4 := a[:4]  //相当于从0切到4
	s5 := a[1:]  //相当于从1切到len(a)
	s6 := a[:]   //相当于从0切到len(a)
	fmt.Println(s3, s4, s5, s6)
	fmt.Printf("len(s4):%d cap(s4):%d\n", len(s4), cap(s4)) //切片的容量是指底层数组的容量
	fmt.Printf("len(s3):%d cap(s3):%d\n", len(s3), cap(s3)) //底层数组从切片的第一个元素到最后的元素数量

	//切片再切片
	s7 := s5[3:] //[9 11 13]
	fmt.Printf("len(s7):%d cap(s7):%d\n", len(s7), cap(s7))
	fmt.Printf("s5=%v s7=%v\n", s5, s7)

	//切片是引用类型，他们都指向了底层数组
	fmt.Println("before:", s7)
	a[5] = 233
	fmt.Println("after:", s7)

	//切片+++
	seliceExtra()

	//切片扩容，追加/删除元素；复制切片
	appendCopy()

	//作业
	homework()
}

func seliceExtra() {
	//使用make函数构造切片
	s := make([]int, 5, 10) //make(切片类型,长度,容量)	//如果不写容量就默认和长度一致
	fmt.Printf("v:%v len(s):%d cap(s):%d\n", s, len(s), cap(s))
	s = make([]int, 0, 10)
	fmt.Printf("v:%v len(s):%d cap(s):%d\n", s, len(s), cap(s))
	//切片就是一个框，框住了一块连续的内存，真正的数据都是保存在底层数组里的

	//切片不能直接比较
	//切片唯一合法的比较操作是和nil比较
	//一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0
	//但是我们不能说一个长度和容量都是0的切片一定是nil，例如下面的示例
	var s1 []int //len(s1)=0;cap(s1)=0;s1==nil
	fmt.Println(s1 == nil, len(s1) == 0)
	s2 := []int{} //len(s2)=0;cap(s2)=0;s2!=nil
	fmt.Println(s2 == nil, len(s2) == 0)
	s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	fmt.Println(s3 == nil, len(s3) == 0)
	//所以要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断。

	//切片的赋值
	s4 := []int{0, 1, 2, 3}
	s5 := s4 //这里指s4和s5都指向了同一个底层数组
	fmt.Println(s4, s5)
	s4[0] = 2333
	fmt.Println(s4, s5)

	//切片的遍历
	//索引遍历
	s6 := []int{0, 1, 2, 3, 4, 5}
	for i := 0; i < len(s6); i++ {
		fmt.Print(s6[i], " ")
	}
	fmt.Println()
	//for range循环
	for _, i := range s6 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func appendCopy() {
	//append()切片扩容
	s := []string{"水瓶", "Zh", "可爱"}
	fmt.Printf("s=%v\tlen(s)=%d\tcap(s)=%d\tptr:%p\n", s, len(s), cap(s), s)
	//s[3] = "牛逼"	//错误的写法，会导致编译错误，索引越界
	//调用append函数必须要用原来的切片变量接收返回值
	//再用append追加元素时，原来的底层数组放不下的时候，go底层就会换一个底层数组
	//必须用变量接收append返回值
	s = append(s, "！")
	fmt.Printf("s=%v\tlen(s)=%d\tcap(s)=%d\tptr:%p\n", s, len(s), cap(s), s)
	//追加多个元素
	s = append(s, "fcc", "也", "可爱")
	// 追加切片
	s1 := []string{"！"}
	s = append(s, s1...) //"..."表示将切片拆开
	fmt.Printf("s=%v\tlen(s)=%d\tcap(s)=%d\tptr:%p\n", s, len(s), cap(s), s)
	//切片的扩容策略
	//见$GOROOT/src/runtime/slice.go
	//从上面的代码可以看出以下内容
	//首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）
	//否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap）
	//否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
	//如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
	///切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样。

	//copy()复制切片
	a1 := []int{1, 2, 3}
	a2 := a1 //赋值
	var a3 = make([]int, 3, 3)
	copy(a3, a2) //使用copy()函数将切片a2中的元素复制到另外一个切片空间中（切片a3）
	fmt.Println(a1, a2, a3)
	a2[0] = 100 //由于切片是引用类型，所以a1和a2其实都指向了同一块内存地址。修改a2的同时a1的值也会发生变化。
	fmt.Println(a1, a2, a3)

	//从切片中删除元素
	a4 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s4 := a4[:]
	fmt.Printf("s4=%v\tlen(s4)=%d\tcap(s4)=%d\n", s4, len(s4), cap(s4))
	s4 = append(s4[:1], s4[2:]...) //将s4索引中的第二个元素（1）删掉,切片长度减少，切片容量不变
	fmt.Printf("s4=%v\tlen(s4)=%d\tcap(s4)=%d\n", s4, len(s4), cap(s4))
	fmt.Println("a4=", a4) //由此可以看出切片的真正数据都是保存在底层数组里，实际的更改在底层数组（将第二个元素删除后调整位置，空出来的地方用前一个的值填充）
}

func homework() {
	a := make([]int, 5, 10)
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	//输出：[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]

	//sort
	b := [...]int{3, 7, 8, 9, 1}
	sort.Ints(b[:]) //将int数组转换为切片类型后进行sort排序（默认升序）
	fmt.Println(b)
}
