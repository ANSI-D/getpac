package main

import (
	"os/exec"
	"strings"
)

type PacmanPackage struct {
	Name        string
	Version     string
	Description string
	Repository  string
}

// checkPacmanPackage checks if a package exists in official Pacman repositories
func checkPacmanPackage(pkgName string) (*PacmanPackage, error) {
	// Use pacman -Si to get package info from sync database
	cmd := exec.Command("pacman", "-Si", pkgName)
	output, err := cmd.Output()
	if err != nil {
		// Package not found in official repos
		return nil, nil
	}
	
	// Parse the output
	lines := strings.Split(string(output), "\n")
	pkg := &PacmanPackage{Name: pkgName}
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Version") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				pkg.Version = strings.TrimSpace(parts[1])
			}
		} else if strings.HasPrefix(line, "Description") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				pkg.Description = strings.TrimSpace(parts[1])
			}
		} else if strings.HasPrefix(line, "Repository") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				pkg.Repository = strings.TrimSpace(parts[1])
			}
		}
	}
	
	return pkg, nil
}

// installPacmanPackage installs a package from official repositories
func installPacmanPackage(pkgName string) error {
	cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", pkgName)
	cmd.Stdin = nil
	return cmd.Run()
}
