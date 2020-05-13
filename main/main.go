package main

import (
	"fmt"
	"os"
)

func main() {
	// 准备工作
	beforRun()

	//

	// 服务初始化
	serviceInit()

	// 运行服务
	if err := Cmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
