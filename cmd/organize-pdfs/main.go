package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Reading Downloads folder...")

	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatalf("get user home dir: %v", err)
	}

	dirPath := filepath.Join(homeDir, "Downloads")

	entries, err := os.ReadDir(dirPath)

	if err != nil {
		log.Fatalf("read folder: %v", err)
	}

	// create a folder for pdfs
	err = os.MkdirAll(filepath.Join(dirPath, "pdfs"), os.ModePerm)

	if err != nil {
		log.Fatalf("create pdfs folder: %v", err)
	}

	moved := 0

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		file := entry

		fileExt := filepath.Ext(file.Name())

		if fileExt == ".pdf" {
			fmt.Println(file.Name())

			sourcePath := filepath.Join(dirPath, file.Name())
			destPath := filepath.Join(dirPath, "pdfs", file.Name())

			err := os.Rename(sourcePath, destPath)

			if err != nil {
				log.Fatalf("move file to folder: %v", err)
			}

			moved++
		}
	}

	fmt.Printf("Successfully moved %d pdfs!", moved)
}
