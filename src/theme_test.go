package main

import "testing"

// We probably want to test this by checking
// If the function ran without calling an error (might add further tests later)
func TestGetThemes(t *testing.T) {
	filename := "../data/themes.json"

	// Checks if code runs without calling its error/breaking
	getThemes(filename)

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
			0.5,
			0.5,
		},
	}
	B := theme{
		Theme:    "test2",
		Topic:    "test",
		Mentions: 10,
		Sentiment: sentiment{
			0.5,
			0.5,
		},
	}
	themes := []theme{A, B}

	topics := groupByTopic(themes)
	// Check if both themes are in this topic
	if len(topics["test"].themes) != 2 {
		t.Errorf("Themes not grouped by topic")
	}
}
