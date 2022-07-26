package main

import "fmt"

//结构体封装
/*
go中根据首字母的大小写来确定可以访问的权限。
无论是方法名、常量、变量名还是结构体的名称，
如果首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用。
如果类名首字母大写，表示被其他包导入时，该类可以被访问
*/
type Hero struct {
	//如果类的属性首字母大写，表示该属性对外可以访问，否则只能内部访问
	Name  string
	Ad    string
	Level int8
	kl    int
}

func (this Hero) GetName() string {
	fmt.Println("Name = ", this.Name)
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func encapsulation() {
	hero := Hero{Name: "小黑", Ad: "555", Level: 100, kl: 5}
	hero.GetName()
	hero.SetName("小红")
	hero.GetName()
	fmt.Println(hero.kl)
}
