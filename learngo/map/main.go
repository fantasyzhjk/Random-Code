package main //map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
}

func main() {
	//m := map[string]bool	//不能这样用，必须初始化
	m := make(map[string]bool, 10)
	m["学习"] = true
	m["玩"] = false
	m["摸鱼"] = true
	fmt.Println(m)
	fmt.Println(m["玩"])
	fmt.Printf("%T\n", m)
	v, ok := m["睡觉"]     //'ok'为任意变量，返回布尔值
	fmt.Println(m["睡觉"]) //如果不存在这个key，返回对应类型的零值
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("睡你麻痹起来嗨！")
	}

	//map的遍历（for range）
	for k, v := range m {
		fmt.Println(k, v)
	}
	//只遍历k
	for k := range m {
		fmt.Println(k)
	}
	//只遍历v
	for _, v := range m {
		fmt.Println(v)
	}

	//删除map中某个元素
	//delete(map, key)
	delete(m, "摸鱼")
	delete(m, "睡觉") //删除不存在的key斌不会报错
	for k, v := range m {
		fmt.Println(k, v)
	}

	demo1()

	//map和切片组合
	s1 := make([]map[int]string, 3, 3) //声明一个元素类型为map的切片
	//要对内部的map做初始化操作
	s1[0] = make(map[int]string, 1)
	s1[0][114514] = "我 爱 你"
	fmt.Println(s1)
	for i, v := range s1 { //遍历
		fmt.Printf("i:%d v:%v\n", i, v)
	}
	s2 := make(map[string][]int, 3) //声明一个元素类型为切片的map
	s2["一班"] = []int{0, 100}
	fmt.Println(s2)

	homework()
}

func demo1() {
	var scoreMap = make(map[string]int, 100)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 100)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

func homework() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%v\n", s) //[1 2 3]
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%v\n", s)         //[1 3]
	fmt.Printf("%v\n", m["q1mi"]) //[1 3 3]

	str := "how do you do"
	retstr := strings.Split(str, " ") //按照定义分割字符串
	result := make(map[string]int, 3)
	for _, v := range retstr {
		result[v]++
	}
	fmt.Println(result)
}
