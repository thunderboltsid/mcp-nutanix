package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/thunderboltsid/mcp-nutanix/internal/codegen/templates"
)

func main() {
	// Parse command-line arguments
	outputDir := flag.String("output", ".", "Output directory for generated files")
	flag.Parse()

	// Get absolute path
	absPath, err := filepath.Abs(*outputDir)
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generating resource and tool files in: %s\n", absPath)

	// Generate all resource files
	if err := templates.GenerateResourceFiles(absPath); err != nil {
		fmt.Printf("Error generating files: %v\n", err)
		os.Exit(1)
	}

	// Generate all resource files
	if err := templates.GenerateToolFiles(absPath); err != nil {
		fmt.Printf("Error generating files: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Resource and tool files generated successfully!")
}
