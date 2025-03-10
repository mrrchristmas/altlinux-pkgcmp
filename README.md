
# Package Comparison Tool

This tool compares packages between two branches of the ALT Linux repository: `sisyphus` and `p10`. It identifies:
- Packages that are only present in one branch.
- Packages where the version in `sisyphus` is newer than in `p10`.

## Table of Contents

1. [Installation](#installation)
2. [Example Output](#example-output)
---

## Features

- Fetches package data from the ALT Linux API.
- Compares packages across architectures (e.g., `x86_64`, `noarch`).
- Outputs results in JSON format.
- Saves results to a file for further analysis.

---

## Installation
```bash
git clone https://github.com/mrrchristmas/altlinux-pkgcmp.git
cd altlinux-pkgcmp
go build -o pkgcmp cmd/main.go
```
## Example Output
```bash
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
```
### Prerequisites

1. Install **Go** (version 1.20 or higher). Check your Go version with:
   ```bash
   go version
