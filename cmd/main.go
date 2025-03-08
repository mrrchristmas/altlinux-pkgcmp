package main

import (
	"fmt"

	"github.com/mrrchristmas/altlinux-pkgcmp/cmd/pkg/api"
)

func main() {
	sisyphusResponse, err := api.FetchPackages("sisyphus")
	if err != nil {
		fmt.Println("Error fetching packages", err)
		return
	}
	p10Response, err := api.FetchPackages("p10")
	if err != nil {
		fmt.Println("Error fetching packages", err)
		return
	}

	// fmt.Println("Sisyphus packages:")
	// for _, pkg := range sisyphusResponse.Packages {
	// 	fmt.Printf("Name: %s, Version: %s, Release: %s, Arch: %s\n",
	// 		pkg.Name, pkg.Version, pkg.Release, pkg.Arch)
	// }

	// fmt.Println("\nP10 packages:")
	// for _, pkg := range p10Response.Packages {
	// 	fmt.Printf("Name: %s, Version: %s, Release: %s, Arch: %s\n",
	// 		pkg.Name, pkg.Version, pkg.Release, pkg.Arch)
	// }
	// Сравнение списка пакетов
	results := api.ComparePackages(sisyphusResponse.Packages, p10Response.Packages)
	err = api.SaveResultJson(results, "result.json")
	if err != nil {
		fmt.Println("Error saving result to json file:", err)
	} else {
		fmt.Println("Result saved to result.json")

		fmt.Println("Comparison results:")
		for _, result := range results {
			fmt.Printf("Architecture: %s\n", result.Arch)
			fmt.Println("Only in p10:\n", result.OnlyInP10)
			fmt.Println("Only in sisyphus:\n", result.OnlyInSisyphus)
			fmt.Println("Newer in sisyphus:")
			fmt.Println("-------------------------------------------")
			for _, pkg := range result.NewerInSisyphus {
				fmt.Printf("Name: %s, P10 version: %s, Sisyphus version: %s\n",
					pkg.Name, pkg.P10Version, pkg.SisyphusVersion)
			}
			fmt.Println("-------------------------------------------")
		}

	}
}
