package main //常量，定义的常量不能修改

import "fmt"

const pi = 3.1415926 //声明常量

const ( //批量声明常量
	time     = 0
	notFound = 404
)

const ( //批量声明常量时，如果某一行声明后没有赋值，默认就和上行一致
	n1 = 1
	n2
	n3
)

const ( //iota在const关键字出现时将被重置为0，const中"每新增一行常量声明"将使iota计数一次（iota+1），iota能简化定义，在枚举时很有用
	s1 = iota //0
	s2 = iota //1
	s3        //2
)

//几个常见的iota示例

const (
	a1 = iota //0
	a2        //1
	_         //_
	a3        //3
)

const ( //插队
	b1 = iota //0
	b2 = 100  //100
	b3        //100
	b4 = iota //3
)

const (
	c1, c2 = iota + 1, iota + 2 //c1 = 1 , c2 = 2
	c3, c4                      //c3 = 2 , c4 = 3
)

const ( //定义数量级
	_  = iota
	kB = 1 << (10 * iota) //"<<"左移符号，相当于1左移10（10*iota）位，为1024（二进制：10000000000）
	mB                    //同上，1左移20位
	gB
	tB
	pB
)

func main() {
	fmt.Printf("n1=%d	n2=%d	n3=%d\n", n1, n2, n3)
	fmt.Printf("s1=%d	s2=%d	s3=%d\n", s1, s2, s3)
	fmt.Printf("a1=%d	a2=%d	a3=%d\n", a1, a2, a3)
	fmt.Printf("b1=%d	b2=%d	b3=%d	b4=%d\n", b1, b2, b3, b4)
	fmt.Printf("c1=%d	c2=%d	c3=%d	c4=%d\n", c1, c2, c3, c4)
	fmt.Println(mB)
}
