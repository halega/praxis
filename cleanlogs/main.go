package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tmpDir := filepath.Join(wd, "tmp")
	cleanLogs(tmpDir)
}

func cleanLogs(logDir string) {
	fmt.Printf("cleanLogs(\"%s\")\n", logDir)
	if logDir == "" {
		fmt.Println("cleanLogs: logDir is empty")
		return
	}
	if filepath.Base(logDir) == "W3CSV" {
		fmt.Println("cleanLogs: pass", logDir)
		return
	}

	files, err := ioutil.ReadDir(logDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
		if file.IsDir() {
			subDir := filepath.Join(logDir, file.Name())
			cleanLogs(subDir)
			continue
		}
		fn := filepath.Join(logDir, file.Name())
		fmt.Printf("Removing %s\n", fn)
	}
}

func removeFiles(dir string) error {
	return filepath.Walk(dir, removeFile)
}

func removeFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if info.IsDir() {
		return nil
	}
	fmt.Printf("Removing %s\n", path)
	return nil
}

func createDir(name string) error {
	_, err := os.Create(name)
	return err
}
