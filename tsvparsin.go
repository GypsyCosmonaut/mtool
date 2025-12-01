package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

// randomName generates a random 5-letter lowercase name.
func randomName() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	name := make([]rune, 5)
	for i := range name {
		name[i] = letters[rand.Intn(len(letters))]
	}
	return string(name)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	filename := "data.tsv"

	// ----- 1. Create TSV file -----
	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating file: %v\n", err)
		os.Exit(1)
	}

	writer := bufio.NewWriter(file)

	for id := 1; id <= 5; id++ {
		line := fmt.Sprintf("%d\t%s\n", id, randomName())
		_, err := writer.WriteString(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error writing tsv: %v\n", err)
			os.Exit(1)
		}
	}

	writer.Flush()
	file.Close()

	// ----- 2. Read TSV file -----
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening tsv file: %v\n", err)
		os.Exit(1)
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)

	uniqueNames := make(map[string]struct{})

	// ----- 3. Parse TSV -----
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		if len(parts) < 2 {
			continue
		}
		name := parts[1]
		uniqueNames[name] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error scanning tsv: %v\n", err)
		os.Exit(1)
	}

	// Convert to slice
	var sortedNames []string
	for n := range uniqueNames {
		sortedNames = append(sortedNames, n)
	}

	// ----- 4. Sort -----
	sort.Strings(sortedNames)

	// Print sorted unique names
	for _, n := range sortedNames {
		fmt.Println(n)
	}

	// ----- 5. Delete the file -----
	if err := os.Remove(filename); err != nil {
		fmt.Fprintf(os.Stderr, "error deleting file: %v\n", err)
		os.Exit(1)
	}
}

