package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	// 创建trace文件夹
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 启动trace
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	// 正常业务
	fmt.Println("hello GMP")

	// 停止trace
	trace.Stop()
}
