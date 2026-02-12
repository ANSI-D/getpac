# getpac

A unified package manager for Arch Linux that seamlessly handles both official repository packages and AUR packages.

## Current Features (Increment 1)

- Install packages from official Pacman repositories or AUR
- Automatically detects package source (official repos first, then AUR)
- For official packages: installs directly using pacman
- For AUR packages:
  - Automatically downloads PKGBUILD from AUR git repository
  - Builds packages using makepkg
  - Installs built packages using pacman

## Prerequisites

- Go 1.16 or higher
- git
- base-devel (for makepkg)
- root access (sudo/doas)

## Installation

```bash
go build -o getpac
sudo mv getpac /usr/local/bin/
```

## Usage

### Available Commands (more to be added)

- `install` (or `-S`) - Install a package from official repos or AUR

### Installing Packages

Install a package (from official repos or AUR):
```bash
getpac install <package-name>
# or
getpac -S <package-name>
```

## Examples

```bash
# Install a package from official repositories
getpac install neovim

# Install a package from AUR (automatically falls back if not in official repos)
getpac install yay
```

## Planned Features

- Increment 2: CLI search functionality
- Increment 3: GUI search interface

## How it works

1. Checks if package exists in official Pacman repositories using `pacman -Si`
2. If found in official repos:
   - Installs directly using `sudo pacman -S --noconfirm`
3. If not found in official repos, checks AUR:
   - Queries AUR RPC API to verify package exists
   - Clones the AUR git repository for the package
   - Runs `makepkg -s --noconfirm` to build the package
   - Installs using `sudo pacman -U --noconfirm`
