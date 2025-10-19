package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting CheckPath")

	fmt.Println("Getting PATH...")
	envPath := os.Getenv("PATH")
	_paths := strings.Split(envPath, ";")
	// Clean up. Trim spaces and remove empty entries
	var paths []string
	for _, s := range _paths {
		if strings.TrimSpace(s) != "" {
			paths = append(paths, s)
		}
	}

	if len(paths) == 0 {
		log.Panic("Did not find any paths")
	}

	// Check that each entry has a directory
	var invalidPaths []string
	for _, p := range paths {
		_, err := os.Stat(p)
		if err != nil {
			invalidPaths = append(invalidPaths, p)
		}
	}

	if len(invalidPaths) == 0 {
		fmt.Println("Found no invalid paths! 👍")
	} else {
		fmt.Println("Found these invalid paths:")
		for _, p := range invalidPaths {
			fmt.Printf("\t%s\n", p)
		}
	}

	fmt.Println("Bye! 👋")
}
