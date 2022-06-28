package main //包名
import "fmt" //导入fmt库

type Person interface {
	Greet()
}
type person struct {
	name string
	age  int
}

func (p person) Greet() {
	fmt.Printf("Hi! My name is %s", p.name)
}

// NewPerson Here, NewPerson returns an interface, and not the person struct itself
func NewPerson(name string, age int) Person {
	return person{
		name: name,
		age:  age,
	}
}

func main() {
	// 结构体指针1
	var a Person
	a = NewPerson("wang", 6)
	a.Greet()
	fmt.Println()
	fmt.Println("hello world")
}
