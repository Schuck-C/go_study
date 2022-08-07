//反射的含义：把变量的原型照出来

package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type : ", reflect.TypeOf(arg))
	fmt.Println("value : ", reflect.ValueOf(arg))
}

func main1() {
	var num float64 = 1.2345
	fmt.Println(reflect.ValueOf(52.5))

	reflectNum(num)
}
