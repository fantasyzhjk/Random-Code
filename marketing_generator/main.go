package main

import "fmt"

func main() {
	var subject, event, event2 string
	fmt.Println("营销号生成器")
	fmt.Print("主体：")
	fmt.Scanln(&subject)
	fmt.Print("事件：")
	fmt.Scanln(&event)
	fmt.Print("另一种说法：")
	fmt.Scanln(&event2)
	fmt.Printf(`%v%v是怎么回事呢？%v相信大家都很熟悉，但是%v%v是怎么回事呢，下面就让小编带大家一起了解吧。%v%v，其实就是%v，大家可能会很惊讶%v怎么会%v呢？但事实就是这样，小编也感到非常惊讶。这就是关于%v%v的事情了，大家有什么想法呢，欢迎在评论区告诉小编一起讨论哦！`, subject, event, subject, subject, event, subject, event, event2, subject, event, subject, event)
}
