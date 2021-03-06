// Echo prints its command-line arguments.
//
package main

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkEchoJoin(b *testing.B) {
	var s, seq string
	for i := 1; i < len(os.Args); i++ {
		s += seq + os.Args[i]
		seq = " "
	}
	fmt.Println(s)
}
