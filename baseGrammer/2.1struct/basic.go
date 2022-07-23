package main

import "fmt"

//定义一个结构体
type Book struct {
	title  string
	author string
}

//结构体值传参
func PrintBook(book Book) {
	fmt.Println(book)
}

//结构体指针传参(即引用传参)
func setTitle(book *Book) {
	book.title = "little prince"
}

func main() {
	//var book Book
	book := Book{}
	book.title = "[Little Prince]"
	book.author = "Tony"
	PrintBook(book)
	setTitle(&book)
	PrintBook(book)
}
