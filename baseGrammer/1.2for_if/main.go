package main

import "fmt"

func main() {
	forDemo()
	fmt.Printf("")
	fmt.Printf((""))
	fmt.Printf("\n")
}

func forDemo() {
	for i := 0; i < 10; i++ {
		switchDemo1(i)
		if i == 0 {
			fmt.Println("000000")
		} else if i == 1 {
			fmt.Println("11111111")
		}
	}
}
func switchDemo1(finger int) {
	//finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")

	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	case 6, 7:
		fmt.Println("666777")
	default:
		//fmt.Println("无效的输入！")
	}
}
