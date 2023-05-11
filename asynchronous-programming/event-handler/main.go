package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
)

// 定义一个事件处理器
func handleOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	// 创建一个命令
	cmd := exec.Command("ls", "-l")

	// 获取命令的输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// 注册事件处理器到输出管道上
	go handleOutput(stdout)

	// 启动命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// 等待命令结束
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
