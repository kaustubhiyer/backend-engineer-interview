package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type sentiment struct {
	Negative float64 `json:"negative"`
	Positive float64 `json:"positive"`
}

type theme struct {
	Theme     string    `json:"theme"`
	Topic     string    `json:"topic"`
	Mentions  int       `json:"mentions"`
	Sentiment sentiment `json:"sentiment"`
}

type topic struct {
	themes       []theme
	netSentiment int
	mentions     int
}

// Retreives a list of themes from a file
func getThemes(filename string) []theme {
	themes := []theme{}

	plainText, err := ioutil.ReadFile(filename)

	if err != nil {
		// Throw and error and close program
		panic(err)
	}

	json.Unmarshal(plainText, &themes)

	fmt.Println(themes)

	return themes
}

// Groups a list of themes by topic, collating statistics about them
func groupByTopic(tSlice []theme) map[string]topic {
	topicSlice := map[string]topic{}

	for _, t := range tSlice {
		if _, ok := topicSlice[t.Topic]; ok {
			// add to topicslice's attributes the current theme
		}
	}

	return topicSlice
}
