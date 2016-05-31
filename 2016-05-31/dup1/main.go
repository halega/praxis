// Dup1 finds and outputs duplicate lines with count of duplication.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	counts := make(map[string]int)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Println(count, "\t", line)
		}
	}
}
