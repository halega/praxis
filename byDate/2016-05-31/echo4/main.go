// Copyright 2016 Stas Kalashnikov <stas@halega.com>. All rights reserved.
// time: 1h

// echo4 prints its command-line arguments with a space delimiter.
package main

import (
	"fmt"
	"os"
)

const sep = " "

func main() {
	fmt.Println("os.Args:", os.Args)
	if len(os.Args) > 1 {
		fmt.Print(os.Args[1])
		for _, arg := range os.Args[2:] {
			fmt.Print(sep, arg)
		}
	}
	fmt.Println()
}
