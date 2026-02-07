package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: getpac <command> <package>")
		fmt.Println("Commands:")
		fmt.Println("  install - Install an AUR package")
		os.Exit(1)
	}

	command := os.Args[1]
	packageName := os.Args[2]

	switch command {
	case "install", "-S":
		if err := installPackage(packageName); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func installPackage(pkgName string) error {
	fmt.Printf("Installing package: %s\n", pkgName)
	
	// Check if package exists in AUR
	pkgInfo, err := getAURPackageInfo(pkgName)
	if err != nil {
		return fmt.Errorf("failed to get package info: %w", err)
	}
	
	if pkgInfo == nil {
		return fmt.Errorf("package '%s' not found in AUR", pkgName)
	}
	
	fmt.Printf("Found: %s (%s)\n", pkgInfo.Name, pkgInfo.Version)
	
	// Download package
	buildDir, err := downloadPackage(pkgName)
	if err != nil {
		return fmt.Errorf("failed to download package: %w", err)
	}
	
	// Build package
	pkgFile, err := buildPackage(buildDir)
	if err != nil {
		return fmt.Errorf("failed to build package: %w", err)
	}
	
	// Install package
	if err := installBuiltPackage(pkgFile); err != nil {
		return fmt.Errorf("failed to install package: %w", err)
	}
	
	fmt.Printf("Successfully installed %s\n", pkgName)
	return nil
}
