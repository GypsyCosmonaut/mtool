package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <directory-path>", os.Args[0])
	}

	dirPath := os.Args[1]

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	for _, entry := range entries {
		fullPath := filepath.Join(dirPath, entry.Name())
		fmt.Println(fullPath)
	}
}

