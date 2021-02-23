package main

import (
	"io/ioutil"
	"os"
	"testing"
)

// We probably want to test this by checking
// If the function ran without calling an error (might add further tests later)
// create a temp json with 1-2 themes in it and see if they load
func TestGetThemes(t *testing.T) {
	os.Remove("_testthemes")

	fileContents := `[{"theme": "testTheme", "topic": "testTopic", "mentions": 300, "sentiment": {"negative": 0.4, "positive": 0.6}}]`
	filename := "_testthemes"

	// write contents to file
	ioutil.WriteFile(filename, []byte(fileContents), 0666)

	// Checks if code runs without calling its error/breaking
	themes := getThemes(filename)

	// check if themes contains testTheme
	if len(themes) != 1 {
		t.Errorf("Did not load themes into slice")
	} else if themes[0].Theme != "testTheme" {
		t.Errorf("Json did not load correctly, loaded %v but expected %v", themes, fileContents)
	}

	os.Remove("_testthemes")

}

// We want to see if a simple case actually groups themes by their common topic
// Build two themes with a common topic, run groupByTopic, and check if both are under 1 topic
// and their statistics are aggregated
func TestGroupByTopic(t *testing.T) {
	A := theme{
		Theme:    "test1",
		Topic:    "test",
		Mentions: 10,
		Sentiment: sentiment{
			"0.5",
			"0.5",
		},
	}
	B := theme{
		Theme:    "test2",
		Topic:    "test",
		Mentions: 10,
		Sentiment: sentiment{
			"0.5",
			"0.5",
		},
	}
	themes := []theme{A, B}

	topics := groupByTopic(themes)
	// Check if both themes are in this topic
	if len(topics["test"].themes) != 2 {
		t.Errorf("Themes not grouped by topic")
	}
}
