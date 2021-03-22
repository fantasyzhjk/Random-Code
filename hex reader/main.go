package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileDir := "./lychee.jpg"
	buFX := make([]byte, 1000, 1000)
	file, _ := os.Open(fileDir)
	n, _ := io.ReadFull(file, buFX)
	// fmt.Println(n)
	buf := make([]byte, n, n)
	file, _ = os.Open(fileDir)
	_, err := io.ReadFull(file, buf)
	if err != nil {
		fmt.Println(n, err.Error())
	} else {
		var i int
		var v byte
		for i, v = range buf {
			fmt.Printf("%02X ", int(v))
			// fmt.Println(i % 16)
			if i%8 == 7 {
				fmt.Printf(" ")
			}
			if i%16 == 15 {
				for ii := i - 15; ii <= i; ii++ {
					if string(buf[ii]) == "\n" || string(buf[ii]) == "\r" {
						fmt.Print(" ")
						continue
					}
					fmt.Print(string(buf[ii]))
				}
				fmt.Println()
			}
		}
		if i < 15 {
			for ii := 0; ii < 15-i; ii++ {
				fmt.Printf("00 ")
				if ((i+ii)%8)+1 == 7 {
					fmt.Printf(" ")
				}
			}
			for ii := 0; ii <= i; ii++ {
				if string(buf[ii]) == "\n" || string(buf[ii]) == "\r" {
					fmt.Print(" ")
					continue
				}
				fmt.Print(string(buf[ii]))
			}
		}
		if i%16 < 15 && i > 15 {
			for ii := 0; ii < 15-i%16; ii++ {
				fmt.Printf("00 ")
				if ((i+ii)%8)+1 == 7 {
					fmt.Printf(" ")
				}
			}
			for ii := i - i%16; ii <= i; ii++ {
				if string(buf[ii]) == "\n" || string(buf[ii]) == "\r" {
					fmt.Print(" ")
					continue
				}
				fmt.Print(string(buf[ii]))
			}
		}
	}
}
