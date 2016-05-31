// Copyright 2016 Stas Kalashnikov <stas@halega.com>. All rights reserved.

// Experiments with fmt.Printf function and its verbs.
package main

import "fmt"

func main() {
	fmt.Printf("%d\n", 5)
	fmt.Printf("%s\n", 5)
	fmt.Printf("%c\n", 5)
	fmt.Printf("%T\n", 5)
	fmt.Printf("%s\n", "hello")
	fmt.Printf("%T\n", "hello")
	fmt.Printf("%q\n", "hello")
	fmt.Printf("%q\n", 5)
	fmt.Printf("%q\n", '5')
	fmt.Printf("%U\n", '5')
}
