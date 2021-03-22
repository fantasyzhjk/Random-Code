package main

import (
	"fmt"
	"strconv"
)

var enKey = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
var enArray = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63}

func main() {
	base64Encode("Ow!==cv123")
}

func base64Decode(str string) string {
	for _, v := range str {
		for i, vv := range enKey {
			if vv == v {
				i++
			}
		}
	}
	return ""
}

func base64Encode(str string) string {
	rstr := []byte(str)
	strb := ""
	for _, v := range rstr {
		strb += convertToBin(int(v))
	}
	println(strb)
	// strlen := len(str)
	var strc [100]string
	for i, v := range strb {
		ii := 0
		strc[ii] += string(v)
		if i%6 == 0 {
			ii++
		}
	}
	for _, v := range strc {
		fmt.Printf("%s/n", v)
	}

	return ""
}

func convertToBin(num int) string {
	s := ""
	if num == 0 {
		return "0"
	}
	for ; num > 0; num /= 2 {
		lsb := num % 2
		s = strconv.Itoa(lsb) + s
	}
	return changeNumber(s)
}

func changeNumber(num string) string {
	i, _ := strconv.Atoi(num)
	return fmt.Sprintf("%08d", i)
}
