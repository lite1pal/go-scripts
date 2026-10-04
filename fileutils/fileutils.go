package fileutils

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func MoveFilesByExtension(dir string, folder string, extensions []string) (int, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return 0, fmt.Errorf("get user home dir: %v", err)
	}

	dirPath := filepath.Join(homeDir, dir)

	entries, err := os.ReadDir(dirPath)

	if err != nil {
		return 0, fmt.Errorf("read folder: %v", err)
	}

	err = os.MkdirAll(filepath.Join(dirPath, folder), os.ModePerm)

	if err != nil {
		return 0, fmt.Errorf("create '%s' folder: %v", folder, err)
	}

	moved := 0

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		file := entry

		fileExt := strings.ToLower(filepath.Ext(file.Name()))

		if slices.Contains(extensions, fileExt) {
			fmt.Println(file.Name())
			sourcePath := filepath.Join(dirPath, file.Name())
			destPath := filepath.Join(dirPath, folder, file.Name())

			err := os.Rename(sourcePath, destPath)

			if err != nil {
				return 0, fmt.Errorf("move '%s' to folder: %v", file.Name(), err)
			}

			moved++
		}
	}

	return moved, nil
}

func DeleteFilesByExtension(dir string, extensions []string) (int, int64, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return 0, 0, fmt.Errorf("get user home dir: %w", err)
	}

	dirPath := filepath.Join(homeDir, dir)

	entries, err := os.ReadDir(dirPath)

	if err != nil {
		return 0, 0, fmt.Errorf("get entries from dir: %w", err)
	}

	var deleted int
	var size int64

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		file := entry

		fileExt := strings.ToLower(filepath.Ext(file.Name()))

		if slices.Contains(extensions, fileExt) {
			fmt.Printf("%s ...👋\n", file.Name())

			err := os.Remove(filepath.Join(dirPath, file.Name()))

			if err != nil {
				fmt.Printf("⚠️ Failed to delete: %s, %w\n", file.Name(), err)
				continue
			}

			deleted++

			fileInfo, err := file.Info()

			if err == nil {
				size = size + fileInfo.Size()
			}
		}
	}

	return deleted, size, nil
}

var sizes = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}

func FormatFileSize(s float64, base float64) string {
	unitsLimit := len(sizes)
	i := 0
	for s >= base && i < unitsLimit {
		s = s / base
		i++
	}

	f := "%.0f%s"
	if i > 1 {
		f = "%.2f%s"
	}

	return fmt.Sprintf(f, s, sizes[i])
}
