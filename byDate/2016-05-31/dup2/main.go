// From the book The Go Programming Language by Donovan and Kernighan.

// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			countLines(counts, f)
			f.Close()
		}
	} else {
		countLines(counts, os.Stdin)
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%q\n", count, line)
		}
	}
}

func countLines(counts map[string]int, file *os.File) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}
