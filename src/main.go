package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	feedbackFile := dirname + "/f.json"

	// Get topics to display
	themes := getThemes(themeFile)
	topics := groupByTopic(themes)
	reviews := retreiveReviews(feedbackFile)
	fmt.Println("Welcome to CLI Customer Review Analytics.")
	fmt.Println()

	// Keeps track of whether we're in feedback or pulse
	var feedback bool
	var id int
	var err error

	for {

		if feedback {

			fmt.Println("Feedback not implemented yet. Look forward to it!")

			// We need the following for the feedback:
			// 1. Results per page
			// 2. sort type
			// 3. topic it's filtered by
			// 4. (later) themes under the topic
			pageSize := 5
			sortType := "date"
			var filterTopic string
			for name, topic := range topics {
				if id == topic.id {
					filterTopic = name
				}
			}

			reviewList := loadReviews(reviews, pageSize, sortType, filterTopic)

			// Section 1
			fmt.Println("Feedback")
			fmt.Println()

			displayFeedback(reviewList)

			fmt.Println()
			fmt.Println("Options: ")
			fmt.Println("Enter sort:<sort-type> where sort-type can be: ")
			fmt.Println("date(default), highest score, lowest score")
			fmt.Println("Enter return to return to pulse")
			fmt.Println("Enter \"exit\" to quit the program")
			scanner := bufio.NewScanner(os.Stdin)
			var line string
			if scanner.Scan() {
				line = strings.TrimSpace(scanner.Text())
			}

			if line == "exit" { // Exit program
				fmt.Println("Thank you for using CLI Customer Review Analytics.")

				os.Exit(0)
			} else if line[:5] == "sort:" { // sort by line
				line = line[5:]
				if line != "highest score" && line != "lowest score" {
					fmt.Println("Invalid sort type, try again")
				} else {
					reviewList = loadReviews(reviews, pageSize, line, filterTopic)
				}
			} else if line == "return" {
				feedback = false
			} else {
				fmt.Println("Invalid Input, try again")
			}
			fmt.Println()
			fmt.Println()
			fmt.Println()

		} else {
			fmt.Println("PULSE")
			fmt.Println()

			displayTopics(topics)

			// Present Menu
			fmt.Println()
			fmt.Println("Options: ")
			fmt.Println("Enter the ID of the topic to expand")
			fmt.Println("Enter the ID followed by F to view feedback analytics")
			fmt.Println("Enter \"exit\" to quit the program")
			scanner := bufio.NewScanner(os.Stdin)
			var line string
			if scanner.Scan() {
				line = strings.TrimSpace(scanner.Text())
			}

			if line == "exit" { // Exit program
				fmt.Println("Thank you for using CLI Customer Review Analytics.")

				os.Exit(0)
			} else if line[len(line)-1] == 'F' { // go to feedback for ID
				line = line[:len(line)-1]
				fmt.Println(line)
				fmt.Println(topics)
				id, err = strconv.Atoi(line)
				id--
				if err != nil {
					fmt.Println("Invalid ID, please try again")
				} else if id >= len(topics) {
					fmt.Println("Invalid ID, please try again")
				} else {
					feedback = true
				}
			} else {
				id, err = strconv.Atoi(line)
				if err != nil {
					fmt.Println("Invalid ID, please try again")
				} else if id >= len(topics) {
					fmt.Println("Invalid ID, please try again")
				} else {
					for name, topic := range topics {
						if id == topic.id {
							selectTopic(topics, name)
						}
					}
				}
			}
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
	}

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
