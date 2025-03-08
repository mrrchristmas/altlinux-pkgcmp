package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	version "github.com/knqyf263/go-rpm-version"
)

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Release string `json:"release"`
	Arch    string `json:"arch"`
}

type ApiResponse struct {
	Packages []Package `json:"packages"`
}
type PackageVersion struct {
	Name            string `json:"name"`
	P10Version      string `json:"p10_version"`
	SisyphusVersion string `json:"sisyphus_version"`
}
type ComparisonResult struct {
	Arch            string           `json:"arch"`
	OnlyInP10       []string         `json:"only_in_p10"`
	OnlyInSisyphus  []string         `json:"only_in_sisyphus"`
	NewerInSisyphus []PackageVersion `json:"newer_in_sisyphus"`
}

func FetchPackages(branch string) (*ApiResponse, error) {
	url := fmt.Sprintf("https://rdb.altlinux.org/api/export/branch_binary_packages/%s", branch)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result ApiResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func ComparePackages(sisyphusPackages, p10Packages []Package) []ComparisonResult {
	sisyphusMap := make(map[string]map[string]Package)
	p10Map := make(map[string]map[string]Package)
	// Заполняем карты для sisyphus
	for _, pkg := range sisyphusPackages {
		if _, exists := sisyphusMap[pkg.Arch]; !exists {
			sisyphusMap[pkg.Arch] = make(map[string]Package)
		}
		sisyphusMap[pkg.Arch][pkg.Name] = pkg
	}

	// Заполняем карты для p10
	for _, pkg := range p10Packages {
		if _, exists := p10Map[pkg.Arch]; !exists {
			p10Map[pkg.Arch] = make(map[string]Package)
		}
		p10Map[pkg.Arch][pkg.Name] = pkg
	}
	// Сравниваем пакеты для каждой архитектуры
	var results []ComparisonResult
	for arch := range sisyphusMap {
		result := ComparisonResult{Arch: arch}

		// Пакеты которые есть только в sisyphus
		for name := range sisyphusMap[arch] {
			if _, exists := p10Map[arch][name]; !exists {
				result.OnlyInSisyphus = append(result.OnlyInSisyphus, name)
			}

		}

		// Пакеты которые есть только в p10
		for name := range p10Map[arch] {
			if _, exists := sisyphusMap[arch][name]; !exists {
				result.OnlyInP10 = append(result.OnlyInP10, name)
			}

		}

		// Пакеты, у которых версия в sisyphus больше
		for name, sisyphusPkg := range sisyphusMap[arch] {
			if p10Pkg, exists := p10Map[arch][name]; exists {
				if isNewer(sisyphusPkg.Version+"-"+sisyphusPkg.Release, p10Pkg.Version+"-"+p10Pkg.Release) {
					result.NewerInSisyphus = append(result.NewerInSisyphus, PackageVersion{
						Name:            name,
						P10Version:      p10Pkg.Version + "-" + p10Pkg.Release,
						SisyphusVersion: sisyphusPkg.Version + "-" + sisyphusPkg.Release,
					})
				}
			}
		}
		results = append(results, result)
	}
	return results

}

func isNewer(version1, version2 string) bool {
	v1 := version.NewVersion(version1)
	v2 := version.NewVersion(version2)
	return v1.GreaterThan(v2)
}

func SaveResultJson(results []ComparisonResult, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	if err := encoder.Encode(results); err != nil {
		return fmt.Errorf("failed to encode results to JSON: %v", err)
	}
	return nil
}
