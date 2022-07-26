package main

import "fmt"

//本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

//具体的类
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("颜色是", this.color, "的小猫 is sleeping..")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

//具体的类
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("颜色是", this.color, "的狗子 is sleeping..")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color=", animal.GetColor())
}

func main() {

	/*
		方法一
		var animal AnimalIF //接口数据类型，父类指针
		animal = &Cat{"大黄"}

		animal.Sleep() //调用Cat的Sleep()方法，多态的现象

		animal = &Dog{"小黑"}
		animal.Sleep() //调用Dog的Sleep()方法，多态的现象
	*/
	cat := &Cat{"黑色"}
	dog := &Dog{"黄色"}
	showAnimal(cat)
	showAnimal(dog)

}
