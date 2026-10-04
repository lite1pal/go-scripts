package main

import (
	"fmt"
	"github.com/lite1pal/go-scripts/fileutils"
	"log"
)

var pdfExtensions []string = []string{".pdf"}

const dir = "Downloads"
const folder = "pdfs"

func main() {
	moved, err := fileutils.MoveFilesByExtension(dir, folder, pdfExtensions)

	if err != nil {
		log.Fatalf("move %s: %v", folder, err)
	}

	fmt.Printf("Successfully moved %d %s!", moved, folder)
}
