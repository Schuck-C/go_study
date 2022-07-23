package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Printf("a:%d \t ptr:%p \t type:%T\n", a, &a, a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%d \t b:%p \t type:%T\n", b, b, b)    // b:0xc00001a078 type:*int
	fmt.Println(&b)                                     // 0xc00000e018
}
