package main

import (
	"fmt"
	"unicode"
)

func main() {
	var str string
	fmt.Scanln(&str)
	c, cnstr := hanZiJianCe(str)
	fmt.Println("汉字数量：", c, "	汉字为：", cnstr)
	chenfabiao()
	huiwenpanduan("上海自来水来自海上")
}

func hanZiJianCe(str string) (c int, cnstr string) {
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			cnstr = cnstr + string(r)
			c++
		}
	}
	if cnstr == "" {
		cnstr = "无汉字"
	}
	return
}

func chenfabiao() {
	for i := 1; i <= 9; i++ {
		for i1 := 1; i1 <= i; i1++ {
			fmt.Printf("%d*%d=%d	", i, i1, i*i1)
		}
		fmt.Println()
	}
}

func huiwenpanduan(s string) {
	//上海自来水来自海上
	//山西运煤车煤运西山
	//黄山落叶松叶落山黄
	r := make([]rune, 0, len(s))
	for _, c := range s {
		r = append(r, c)
	}
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("此句不为索引")
			return
		}
	}
	fmt.Println("此句为索引")
}
