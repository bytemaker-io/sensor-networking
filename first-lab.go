package main

import "fmt"

func main() {
	fmt.Println("嗨客网(www.haicoder.net)")
	defer func() {
		if info := recover(); info != nil {
			fmt.Println("触发了宕机, Info =", info)
		} else {
			fmt.Println("程序正常退出")
		}
	}()
	fmt.Println("Hai")
	fmt.Println("Coder")
	panic("fatal error")
	fmt.Println("Over")

}
