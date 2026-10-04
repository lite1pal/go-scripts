package main

import (
	"fmt"
	"github.com/lite1pal/go-scripts/fileutils"
	"log"
)

var imgExtensions []string = []string{".png", ".jpg", ".webp", ".jpeg", ".gif", ".heif"}

const dir = "Downloads"
const folder = "images"

func main() {
	moved, err := fileutils.MoveFilesByExtension(dir, folder, imgExtensions)

	if err != nil {
		log.Fatalf("move %s: %v", folder, err)
	}

	fmt.Printf("Successfully moved %d %s!", moved, folder)
}
