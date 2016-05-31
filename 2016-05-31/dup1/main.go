// From the book The Go Programming Language by Donovan and Kernighan.

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
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
