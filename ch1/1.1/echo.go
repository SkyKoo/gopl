// 练习 1.1： 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
// Echo prints its command-line arguments.
//
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}
