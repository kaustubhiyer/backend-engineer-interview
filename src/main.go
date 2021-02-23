package main

import "fmt"

func main() {
	themes := getThemes("../data/themes.json")

	topics := groupByTopic(themes)

	fmt.Printf("%+v", topics)
}
