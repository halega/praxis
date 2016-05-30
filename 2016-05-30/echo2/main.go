// time:

// Echos program's arguments using os package
package main

import "os"

func main() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += arg + sep
		sep = " "
	}
	s += "\n"
	os.Stdout.WriteString(s)
}
