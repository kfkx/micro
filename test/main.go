package main

import (
	"fmt"
)

var (
	sId          int
	sName        string
	sDescription string
	sWorkspace   string
)

func main() {

	fmt.Print("请输入服务ID (例如： 1024): ")
	fmt.Scanln(&sId)

	fmt.Print("请输入服务名称 (例如： core): ")
	fmt.Scanln(&sName)

	fmt.Print("请输入服务介绍 (例如： 核心服务): ")
	fmt.Scanln(&sDescription)

	fmt.Print("请输入服务地址 (例如： $GOPATH/src/demo): ")
	fmt.Scanln(&sWorkspace)

}
