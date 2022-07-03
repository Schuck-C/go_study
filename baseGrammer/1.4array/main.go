package main

import "fmt"

//Print系列函数会将内容输出到系统的标准输出，
//区别在于Print函数直接输出内容，
//Printf函数支持格式化输出字符串，
//Println函数会在输出内容的结尾添加一个换行符。

func main() {
	demo6()
}

func demo0() {
	var a [3]int
	a[0] = 10
	fmt.Println(a[0])
	//var b [4]int
	//a = b //不可以这样做，因为此时a和b是不同的类型
}

//================================================================================
/*数组的初始化
数组的初始化也有很多方式。
*/

//方法一：初始化列表
//初始化数组时可以使用初始化列表来设置数组元素的值。

func demo1() {
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]
}

//方法二：使用...
//按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以使用...
//让编译器根据初始值的个数自行推断数组的长度，例如：

func demo2() {
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)                          //[0 0 0]
	fmt.Println(numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
}

//方法三：指定索引值
//我们还可以使用指定索引值的方式来初始化数组，例如:

func demo3() {
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int
}

//================================================================================================================================
//数组的遍历
//数组的遍历
//遍历数组a有以下两种方法：

func demo4() {
	var a = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}

//多维数组
//Go语言是支持多维数组的，我们这里以二维数组为例（数组中又嵌套数组）。

//二维数组的定义
func demo5() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆
}

//二维数组的遍历
func demo6() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range a {
		//fmt.Printf("%s\t", v1)
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}

/*
输出：

北京	上海
广州	深圳
成都	重庆
注意： 多维数组只有第一层可以使用...来让编译器推导数组长度。例如：

//支持的写法
a := [...][2]string{
{"北京", "上海"},
{"广州", "深圳"},
{"成都", "重庆"},
}
//不支持多维数组的内层使用...
b := [3][...]string{
{"北京", "上海"},
{"广州", "深圳"},
{"成都", "重庆"},
}
数组是值类型
数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
*/

func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}
func demo7() {
	a := [3]int{10, 20, 30}
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Println(a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
}

/*注意：

数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。
[n]*T表示指针数组，*[n]T表示数组指针 。
练习题
求数组[1, 3, 5, 7, 8]所有元素的和
*/
var (
	nums = [...]int{1, 3, 5, 7, 8}
)

func demo8() {
	res := 0
	for _, num := range nums {
		res += num
	}
	fmt.Println(res)
}

//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
