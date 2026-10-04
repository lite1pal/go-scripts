package main

import (
	"fmt"
	"github.com/lite1pal/go-scripts/fileutils"
	"log"
)

const dir = "Downloads"

var extensions []string = []string{".dmg"}

func main() {
	fmt.Println("‼️ Deleting dmgs...")

	deleted, size, err := fileutils.DeleteFilesByExtension("Downloads", extensions)

	if err != nil {
		log.Fatalf("delete dmgs: %v", err)
	}

	formattedSize := fileutils.FormatFileSize(float64(size), 1024.0)

	if deleted > 0 {
		fmt.Printf("Successfully deleted %d dmgs, freed up %s\n", deleted, formattedSize)
	} else {
		fmt.Print("Nothing to delete.\n")
	}
}
