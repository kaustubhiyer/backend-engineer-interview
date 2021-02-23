package main

import "fmt"

func main() {
	themes := getThemes("../data/themes.json")

	fmt.Println(themes)
}
