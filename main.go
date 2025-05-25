package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type context struct {
	currentPath string
	levelPaths  map[int]string
}

func main() {
	// CLI flags
	inputPath := flag.String("input", "", "Path to the .txt or .md file")
	outputDir := flag.String("output", "", "Path to the root folder to create structure")
	dryRun := flag.Bool("dry", false, "Dry run (no files/folders will be created)")
	// TODO: "help" flag
	flag.Parse()

	if *inputPath == "" || *outputDir == "" {
		fmt.Println("Usage (optional flag: -dry): go run main.go -input=structure.txt -output=./output")
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
	ctx := context{levelPaths: make(map[int]string)}
	ctx.levelPaths[0] = *outputDir

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			// Count the level
			level := strings.Count(line, "#")
			folderName := strings.TrimSpace(line[level:])
			parentPath := ctx.levelPaths[level-1]
			currentFolder := filepath.Join(parentPath, folderName)

			if !*dryRun {
				err := os.MkdirAll(currentFolder, os.ModePerm)
				if err != nil {
					fmt.Printf("Error creating folder %s: %v\n", currentFolder, err)
				}
			}
      // TODO: improve visualization
			fmt.Printf("[FOLDER] %s\n", currentFolder)
			ctx.levelPaths[level] = currentFolder
			ctx.currentPath = currentFolder
		} else if strings.HasPrefix(line, "-") {
			// File inside current folder
			fileName := strings.TrimSpace(strings.TrimPrefix(line, "-"))
			if ctx.currentPath == "" {
				fmt.Println("Error: File defined before any folder.")
				continue
			}
			filePath := filepath.Join(ctx.currentPath, fileName)
			// TODO: improve visualization
			fmt.Printf("[FILE]   %s\n", filePath)
			if !*dryRun {
				err := os.WriteFile(filePath, []byte("// File created by CLI tool\n"), 0644)
				if err != nil {
					fmt.Printf("Error creating file %s: %v\n", filePath, err)
				}
			}
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
}
