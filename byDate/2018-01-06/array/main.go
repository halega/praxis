package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [...]string{"bananas", "apples", "cucumbers", "elderberries", "dates"}
	fmt.Println(arr)
	sort.Strings(arr[:])
	fmt.Println(arr)
}
