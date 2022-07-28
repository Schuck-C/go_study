package main

import (
	"fmt"
	"io"
	"os"
)

func demo1() {
	//var a string
	////a的内部pair
	////pair<statictype:string, value: "aceld">
	//a := "aceld"
	//
	////pair<type:string, value:"aceld">
	//var allType interface{}
	//allType = a
	//
	//str, _ := allType.(string)
	//fmt.Println(str)
}

func demo2() {
	//tty: pair<type:*os.File, value:"/dev/tty"文件描述符>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	//r: pair<type:  , value:>
	var r io.Reader
	//r: pair<type:*os.File, value:"/dev/tty"文件描述符>
	r = tty

	//w: pair<type:  , value:>
	var w io.Writer
	//w: pair<type:*os.File, value:"/dev/tty"文件描述符>
	w = r.(io.Writer)

	w.Write([]byte("HELLO THIS is A TEST!!!\n"))
}
