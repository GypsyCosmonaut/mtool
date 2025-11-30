package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	showAll := flag.Bool("a", false, "show hidden files")
	showModTime := flag.Bool("m", false, "show modification time")
	flag.Parse()

	// Determine directory path
	var dirPath string

	// If user passed a directory argument, use it
	if flag.NArg() > 0 {
		dirPath = flag.Arg(0)
	} else {
		// Otherwise use current working directory
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get working directory: %v", err)
		}
		dirPath = wd
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	for _, entry := range entries {
		name := entry.Name()

		// Skip hidden files if -a is not set
		if !*showAll && strings.HasPrefix(name, ".") {
			continue
		}

		fullPath := filepath.Join(dirPath, name)

		// If -m, fetch file info for mod time
		if *showModTime {
			info, err := entry.Info()
			if err != nil {
				fmt.Printf("%s (error reading info)\n", fullPath)
				continue
			}
			mod := info.ModTime().Format(time.RFC3339)
			fmt.Printf("%s  %s\n", fullPath, mod)
		} else {
			fmt.Println(fullPath)
		}
	}
}

