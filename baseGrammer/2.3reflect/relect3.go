//6.通过反射调用方法
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Id   int
	Name string
}

func (s Student) EchoName(name string) {
	fmt.Println("我的名字是：", name)
}

func main3() {
	s := Student{Id: 1, Name: "咖啡色的羊驼"}

	v := reflect.ValueOf(s)

	// 获取方法控制权
	// 官方解释：返回v的名为name的方法的已绑定（到v的持有值的）状态的函数形式的Value封装
	mv := v.MethodByName("EchoName")
	// 拼凑参数
	args := []reflect.Value{reflect.ValueOf("咖啡色的羊驼")}

	// 调用函数
	mv.Call(args)
}
