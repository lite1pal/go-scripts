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
