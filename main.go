package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// CLI flags
	inputPath := flag.String("input", "", "Path to the .txt or .md file")
	outputDir := flag.String("output", "", "Path to the root folder to create structure")
	flag.Parse()

	if *inputPath == "" || *outputDir == "" {
		fmt.Println("Usage: go run main.go -input=structure.txt -output=./output")
		os.Exit(1)
	}

	// Open the input file
	file, err := os.Open(*inputPath)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentFolder string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") {
			// Folder
			folderName := strings.TrimSpace(strings.TrimPrefix(line, "#"))
			currentFolder = filepath.Join(*outputDir, folderName)
			err := os.MkdirAll(currentFolder, os.ModePerm)
			if err != nil {
				fmt.Printf("Error creating folder %s: %v\n", currentFolder, err)
			}
		} else if strings.HasPrefix(line, "-") {
			// File inside the current folder
			if currentFolder == "" {
				fmt.Println("Error: File defined before any folder.")
				continue
			}
			fileName := strings.TrimSpace(strings.TrimPrefix(line, "-"))
			filePath := filepath.Join(currentFolder, fileName)
			emptyFile, err := os.Create(filePath)
			if err != nil {
				fmt.Printf("Error creating file %s: %v\n", filePath, err)
			}
			emptyFile.Close()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
