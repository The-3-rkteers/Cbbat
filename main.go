package main

import "github.com/grantseltzer/Cbbat/imagemake"

func main() {
	imagemake.BuildTree()

	// Build tree of directory entries from root
	// Save as JSON file
	// Create container image with JSON file and executable inside
	// Reconstruct directory tree using JSON file (mock objs)
	// Run executable
	// Build new directory tree, save as JSON
	// Compare two JSON files, report back to user
}
