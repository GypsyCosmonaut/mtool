package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

// randomName generates a random 5-letter uppercase name.
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

	filename := "data.csv"

	// ----- 1. Create CSV File -----
	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating file: %v\n", err)
		os.Exit(1)
	}
	writer := csv.NewWriter(file)

	for id := 1; id <= 5; id++ {
		name := randomName()
		record := []string{fmt.Sprintf("%d", id), name}
		err := writer.Write(record)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error writing csv: %v\n", err)
			os.Exit(1)
		}
	}

	writer.Flush()
	file.Close()

	// ----- 2. Read CSV File -----
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
		os.Exit(1)
	}
	reader := csv.NewReader(readFile)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading csv: %v\n", err)
		os.Exit(1)
	}
	readFile.Close()

	// ----- 3. Extract unique names -----
	uniqueNames := make(map[string]struct{})

	for _, row := range records {
		if len(row) < 2 {
			continue
		}
		name := row[1]
		uniqueNames[name] = struct{}{}
	}

	// Convert map → slice
	var sortedNames []string
	for n := range uniqueNames {
		sortedNames = append(sortedNames, n)
	}

	// ----- 4. Sort names -----
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

