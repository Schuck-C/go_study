package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//demo1()
	demo3()
}
func demo1() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)

	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
}

func demo2() {
	logFile, err := os.OpenFile("./1.6log/xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[小王子]")
	log.Println("这是一条很普通的日志。")
}

func demo3() {
	logger := log.New(os.Stdout, "<New1>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志。")
}

//总结
//Go内置的log库功能有限，//例如无法满足记录不同级别日志的情况，
//我们在实际的项目中根据自己的需要选择使用第三方的日志库，如logrus、zap等。
