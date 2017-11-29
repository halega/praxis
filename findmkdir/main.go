package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filepath.Walk(curDir, find)
}

func find(path string, info os.FileInfo, err error) error {
	if err != nil { return nil }
	if info.IsDir() { return nil }

	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	findStr := []byte("os.Mkdir")
	scanner := bufio.NewScanner(file)
	line := 1
	for scanner.Scan() {
		if bytes.Contains(scanner.Bytes(), findStr) {
			fmt.Printf("Found in %s at line %d\n", file.Name(), line)
		}
		line++
	}
	return nil
}
