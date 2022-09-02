package main

import (
	_ "./matchers"
	//_ "github.com/goinaction/code/chapter2/sample/matchers"
	//"github.com/goinaction/code/chapter2/sample/search"
	"go-in-action-master/chapter2/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
