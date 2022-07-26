package main

import "fmt"

//Human结构体
type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...", this.name)
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperHuman struct {
	Human //SuperHuman类继承了Human类的方法
	level int
}

//重定义父类的方法Eat
func (this *SuperHuman) Eat() {
	fmt.Println("SuperHuman.Eat()...", this.name)
}

//子类的新方法
func (this *SuperHuman) Fly() {
	fmt.Println("SuperHuman.Fly()...")
}

func inheritance() {
	human := Human{"张飒", "女"}
	human.Eat()
	s := SuperHuman{human, 10}
	s.Eat()
	fmt.Println(s.level)
	s.Fly()

}
