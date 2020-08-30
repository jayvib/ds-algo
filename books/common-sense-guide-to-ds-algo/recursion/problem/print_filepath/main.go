package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func findDirectories(root string) error {
	// Step 1: Get the content of the directory
	// Step 2: Check each content of the directory
	// Step 3: If directory then traverse deeper
	// Step 4:(main case) Print the path of the files
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(root, file.Name())
		if file.IsDir() && !ignoreDirs[file.Name()] {
			err := findDirectories(filePath)
			if err != nil {
				return err
			}
		}
		fmt.Println(filePath)
	}
	return nil
}

var ignoreDirs = map[string]bool{
	".git":  true,
	".idea": true,
}

func main() {
	dirPath := "/home/jaysonv/go/src/ds-algo"
	if err := findDirectories(dirPath); err != nil {
		log.Fatal(err)
	}
}
