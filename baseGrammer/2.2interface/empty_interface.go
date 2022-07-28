package main

import "fmt"

//interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	//interface{}如何区分 此时引用的底层数据类型到底是什么？

	//给interface{}提供“类型断言的机制”
	value, ok := arg.(string)
	if !ok {
		fmt.Println("非string")
	} else {
		fmt.Println("this string is = ", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"刘德华"}
	myFunc(book)
	myFunc(100)
	myFunc("abc")
}
