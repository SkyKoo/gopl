// 练习 1.4： 修改dup2，出现重复的行时打印文件名称。

// Dup2 prints the count and text of lines that appear more than once in the input.
// It reads from stdin or from a list of named files.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if counts[line] > 1 {
				fmt.Printf("%s\t%s\n", filename, line)
			}
		}
	}
}
