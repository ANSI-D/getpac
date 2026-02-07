# getpac

A simple AUR package manager written in Go.

## Current Features (Increment 1)

- Install packages from AUR
- Automatically downloads PKGBUILD from AUR git repository
- Builds packages using makepkg
- Installs built packages using pacman

## Prerequisites

- Go 1.16 or higher
- git
- base-devel (for makepkg)
- sudo access

## Installation

```bash
go build -o getpac
sudo mv getpac /usr/local/bin/
```

## Usage

Install an AUR package:
```bash
getpac install <package-name>
# or
getpac -S <package-name>
```

## Examples

```bash
# Install a package from AUR
getpac install yay
```

## Planned Features

- Increment 2: CLI search functionality
- Increment 3: GUI search interface

## How it works

1. Queries AUR RPC API to verify package exists
2. Clones the AUR git repository for the package
3. Runs `makepkg -s --noconfirm` to build the package
4. Installs using `sudo pacman -U --noconfirm`
