package main

import (
	"log"
	"os"
	_ "sample/matchers"
	"sample/search"
)

func init() {
	//日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
