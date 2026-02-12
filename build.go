package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const aurGitURL = "https://aur.archlinux.org"

// downloadPackage clones the AUR git repository for the package
func downloadPackage(pkgName string) (string, error) {
	buildDir := filepath.Join("/tmp", "getpac-"+pkgName)
	
	// removes existing build directory if it exists
	os.RemoveAll(buildDir)
	
	gitURL := fmt.Sprintf("%s/%s.git", aurGitURL, pkgName)
	
	cmd := exec.Command("git", "clone", gitURL, buildDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		return "", err
	}
	
	return buildDir, nil
}

// runs makepkg to build the package
func buildPackage(buildDir string) (string, error) {
	fmt.Println("Building package...")
	
	cmd := exec.Command("makepkg", "-s", "--noconfirm")
	cmd.Dir = buildDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		return "", err
	}
	
	// Find the built package file
	files, err := filepath.Glob(filepath.Join(buildDir, "*.pkg.tar.zst"))
	if err != nil {
		return "", err
	}
	
	if len(files) == 0 {
		return "", fmt.Errorf("no package file found after build")
	}
	
	return files[0], nil
}

// installs the built pkg with pacman
func installBuiltPackage(pkgFile string) error {
	fmt.Printf("Installing %s...\n", filepath.Base(pkgFile))
	
	cmd := exec.Command("sudo", "pacman", "-U", "--noconfirm", pkgFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	
	return cmd.Run()
}
