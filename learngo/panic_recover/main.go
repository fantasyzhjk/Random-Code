package main //panic,recover
import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// recover()必须搭配defer使用。
// defer一定要在可能引发panic的语句之前定义。

func clear() {
	clear := make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	v, isOk := clear[runtime.GOOS]
	if isOk {
		v()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(") //程序崩溃
	}
}

func a() {
	fmt.Println("a")
}

func b() {
	defer func() {
		err := recover() //尝试恢复程序
		fmt.Println(err)
		fmt.Println("清除数据")
	}()
	panic("奔溃啦！")    //程序崩溃
	fmt.Println("b") //此语句执行不到
}

func c() {
	fmt.Println("c")
	panic("奔溃啦！")
}

func main() {
	clear()
	a()
	b()
	c()
}
