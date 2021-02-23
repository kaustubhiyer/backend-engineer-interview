package main

import (
	"flag"
	"fmt"
	"os"
)

// I should initialize the dataset from here

func main() {

	argCount := len(os.Args[1:])

	var dirname string

	flag.StringVar(&dirname, "dir", "data", "Specify directory to use. Defaults to data")
	flag.Parse()

	if argCount != 1 {
		usageInstructions()
		os.Exit(0)
	}

	// Create filenames for use
	themeFile := dirname + "/themes.json"
	// feedbackFile := dirname + "/f.json"

	// Get topics to display
	themes := getThemes(themeFile)
	topics := groupByTopic(themes)
	fmt.Println("Welcome to CLI Customer Review Analytics.")
	fmt.Println()
	fmt.Println("PULSE")
	fmt.Println()

	displayTopics(topics)

	// Present Menu
}

func usageInstructions() {
	fmt.Println("Welcome to CLI Customer Review Analytics. Here are instructions on usage:")
	fmt.Println()
	fmt.Println("To use, create a build using \"go build\" and follow the below format:")
	fmt.Println()
	fmt.Println("> ./output_file -dir=directory_where_data_is")
	fmt.Println()
	fmt.Println("Files containing your datasets must be labelled f.json and themes.json")
	fmt.Println("And be within \"directory_where_data_is\"")
}
