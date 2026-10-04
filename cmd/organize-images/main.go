package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
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

	err = os.MkdirAll(filepath.Join(dirPath, "images"), os.ModePerm)

	if err != nil {
		log.Fatalf("create images folder: %v", err)
	}

	moved := 0

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		file := entry

		fileExt := strings.ToLower(filepath.Ext(file.Name()))

		imgExts := []string{".png", ".jpg", ".webp", ".jpeg", ".gif", ".heif"}

		if slices.Contains(imgExts, fileExt) {
			fmt.Println(file.Name())
			sourcePath := filepath.Join(dirPath, file.Name())

			destPath := filepath.Join(dirPath, "images", file.Name())

			err := os.Rename(sourcePath, destPath)

			if err != nil {
				log.Fatalf("move image to folder: %v", err)
			}

			moved++
		}
	}

	fmt.Printf("Successfully moved %d images!", moved)
}
