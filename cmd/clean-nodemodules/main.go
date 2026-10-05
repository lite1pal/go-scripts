package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

/*
	Potential improvements:

	- add args support to specify projects directory via cli
	- display sum of deleted folders sizes
	- check modification time of all files inside folder; currently only parent folder mod time is compared with cutoff time
*/

const dirWithProjects = "Work/Projects"

var foldersToDelete = map[string]bool{"node_modules": true, ".next": true}

func main() {
	// get user home dir
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("get user home dir: %v", err)
	}

	// read dir entries
	projectsDirPath := filepath.Join(homeDir, dirWithProjects)
	entries, err := os.ReadDir(projectsDirPath)
	if err != nil {
		log.Fatalf("read dir: %v", err)
	}

	now := time.Now()
	start := now

	// declare cutoff time, e.g. 7 days ago
	cutoff := now.AddDate(0, 0, -7)

	pathsToDelete := make([]string, 0)

	var wg sync.WaitGroup
	var mu sync.Mutex

	// loop through the entries
	for _, entry := range entries {
		// skip files
		if !entry.IsDir() {
			continue
		}

		// get entry info with last modification time
		dirInfo, err := entry.Info()
		if err != nil {
			log.Printf("read %s: %v", entry.Name(), err)
			continue
		}

		// skip if modified before cutoff date
		if !dirInfo.ModTime().Before(cutoff) {
			continue
		}

		wg.Add(1)
		// start goroutine
		go func() {
			defer wg.Done()

			// walk the entry dir recursively
			err := filepath.WalkDir(filepath.Join(projectsDirPath, entry.Name()), func(path string, info fs.DirEntry, err error) error {
				if err != nil {
					log.Printf("walk %s: %v", path, err)
					return nil
				}

				// skip files
				if !info.IsDir() {
					return nil
				}

				// skip .git
				if info.Name() == ".git" {
					return filepath.SkipDir
				}

				// skip walking the founded folder
				if foldersToDelete[info.Name()] {
					// lock array append to avoid race condition
					mu.Lock()
					pathsToDelete = append(pathsToDelete, path)
					mu.Unlock()

					return filepath.SkipDir
				}

				return nil
			})

			if err != nil {
				log.Printf("walk %s: %v", entry.Name(), err)
			}
		}()
	}

	wg.Wait()

	// remove found paths
	for _, path := range pathsToDelete {
		// safety check #1: directory must be one we explicitly allow deleting
		if !foldersToDelete[filepath.Base(path)] {
			log.Printf("refusing to delete unexpected directory: %s", path)
			continue
		}

		// safety check #2: path must still be inside specified dir, e.g. Work/Projects
		rel, err := filepath.Rel(projectsDirPath, path)
		if err != nil {
			log.Printf("check path %s: %v", path, err)
			continue
		}

		if rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
			log.Printf("refusing to delete path outside projects directory: %s", path)
			continue
		}

		fmt.Printf("Deleting: %s\n", path)

		// remove all folders under the paths
		if err := os.RemoveAll(path); err != nil {
			log.Printf("remove %s: %v", path, err)
		}
	}

	if len(pathsToDelete) > 0 {
		fmt.Printf("Successfully deleted %d folders\n", len(pathsToDelete))
	} else {
		fmt.Print("Nothing to delete.\n")
	}

	fmt.Printf("time spent: %d ms\n", time.Since(start).Milliseconds())
}
