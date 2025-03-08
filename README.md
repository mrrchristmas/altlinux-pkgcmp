# Package Comparison Tool

This tool compares packages between two branches of the ALT Linux repository: `sisyphus` and `p10`. It identifies:
- Packages that are only present in one branch.
- Packages where the version in `sisyphus` is newer than in `p10`.

## Table of Contents

1. [Installation](#installation)
2. [Example Output](#example-output)
---

## Installation
git clone https://github.com/your-username/altlinux-pkgcmp.git
cd altlinux-pkgcmp
go build -o pkgcmp cmd/main.go

## Example Output
[
  {
    "Arch": "x86_64",
    "OnlyInP10": ["package5"],
    "OnlyInSisyphus": ["package1", "package2"],
    "NewerInSisyphus": [
      {
        "Name": "package1",
        "P10Version": "1.0-alt1",
        "SisyphusVersion": "1.0-alt2"
      }
    ]
  }
]

### Prerequisites

1. Install **Go** (version 1.20 or higher). Check your Go version with:
   ```bash
   go version