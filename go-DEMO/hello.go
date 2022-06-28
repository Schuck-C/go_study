package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

type person struct {
	name string
	city string
	age  int8
}

func main() {
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "七米"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}
}
