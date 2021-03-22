package main //字符串的应用

import (
	"fmt"
	"strings"
)

func main() {
	Path := "\"D:\\Go\\src\\github.com\\fantayszhjk\\day1\\string\""
	fmt.Println(Path)
	//多行字符串
	s := `
		233
			quq	232
		323
	`
	fmt.Println(s)
	s1 := `D:\Go\src\github.com\fantayszhjk\day1\string` //反引号按原样输出
	fmt.Println(s1)
	//字符串相关操作
	fmt.Println(len(s1)) //输出该字符串长度
	//字符串拼接
	ss1 := "wow "
	ss2 := "u look fun"
	ss := ss1 + ss2
	fmt.Println(ss)
	ssr := fmt.Sprintf("%s%s", ss1, ss2)
	fmt.Println(ssr)
	//分割
	ret := strings.Split(s1, "\\") //按照定义分割字符串
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(ss1, "u")) //判断是否包含，输出true/false
	//前缀,后缀
	fmt.Println(strings.HasPrefix(ss1, "w")) //判断前缀
	fmt.Println(strings.HasSuffix(ss1, "w")) //判断后缀
	//判断位置
	fmt.Println(strings.Index(s1, "o"))
	fmt.Println(strings.LastIndex(s1, "o"))
	//拼接
	fmt.Println(strings.Join(ret, "+")) //用“+”连接字符串
}

// \r	回车符（返回行首）
// \n	换行符（直接跳到下一行的同列位置）
// \t	制表符
// \'	单引号
// \"	双引号
// \\	反斜杠
