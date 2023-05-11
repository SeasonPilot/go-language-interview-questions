package main

import (
	"log"
	"os/exec"
)

// 定义一个回调函数
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 创建一个命令
	cmd := exec.Command("ls", "-l")

	// 运行命令，并传递回调函数
	cmd.Run(handleError) //fixme: 这块代码有问题,这块主要为了举例说明回调函数是作为参数传递给另一个函数
}
