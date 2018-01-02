package main

import "fmt"

func main() {
	var drawCats = func(howManyTimes int) {
		for i := 0; i < howManyTimes; i++ {
			fmt.Printf("%d =^.^=\n", i)
		}
	}
	drawCats(100)
}
