package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: getpac <command> <package>")
		fmt.Println("Commands:")
		fmt.Println("  install - Install a package from official repos or AUR")
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
	
	// First, check if package exists in official Pacman repositories
	pacmanPkg, err := checkPacmanPackage(pkgName)
	if err != nil {
		return fmt.Errorf("failed to check pacman repositories: %w", err)
	}
	
	if pacmanPkg != nil {
		// Package found in official repos
		fmt.Printf("Found in %s repository: %s (%s)\n", pacmanPkg.Repository, pacmanPkg.Name, pacmanPkg.Version)
		if pacmanPkg.Description != "" {
			fmt.Printf("Description: %s\n", pacmanPkg.Description)
		}
		
		fmt.Println("Installing from official repository...")
		if err := installPacmanPackage(pkgName); err != nil {
			return fmt.Errorf("failed to install from pacman: %w", err)
		}
		
		fmt.Printf("Successfully installed %s\n", pkgName)
		return nil
	}
	
	// Package not in official repos, check AUR
	fmt.Println("Not found in official repositories, checking AUR...")
	pkgInfo, err := getAURPackageInfo(pkgName)
	if err != nil {
		return fmt.Errorf("failed to get package info: %w", err)
	}
	
	if pkgInfo == nil {
		return fmt.Errorf("package '%s' not found in official repositories or AUR", pkgName)
	}
	
	fmt.Printf("Found in AUR: %s (%s)\n", pkgInfo.Name, pkgInfo.Version)
	
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
