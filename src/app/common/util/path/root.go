package path

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindProjectRootPath() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return searchRootPath(currentDir)
}

func searchRootPath(startPath string) (string, error) {
	// Check if the go.mod file exists in the current directory
	goModPath := filepath.Join(startPath, "go.mod")
	_, err := os.Stat(goModPath)
	if err == nil {
		// go.mod found, return the current directory as the root
		return startPath, nil
	}

	// If go.mod is not found and we have reached the root directory, return an error
	if startPath == filepath.Dir(startPath) {
		return "", fmt.Errorf("go.mod not found")
	}

	// Recursively search in the parent directory
	return searchRootPath(filepath.Dir(startPath))
}
