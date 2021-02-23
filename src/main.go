package main

// I should initialize the dataset from here

func main() {
	themes := getThemes("../data/themes.json")

	topics := groupByTopic(themes)

	displayTopics(topics)
}
