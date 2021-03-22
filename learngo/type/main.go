package main //自定义类型和类型别名
import "fmt"

type (
	myInt   int64   // 自定义类型
	yourInt = int64 // 类型别名
)

// golang的rune和byte都为类型别名

func main() {
	var n myInt
	n = 1
	fmt.Printf("%T\n", n)
	var m yourInt
	m = 1
	fmt.Printf("%T\n", m)
}
