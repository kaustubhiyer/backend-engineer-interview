package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestRetreiveReviews(t *testing.T) {
	os.Remove("_testreviews")

	fileContents := `[{"text": "Its so bad", "score": 0.0, "date": "2018-06-09", "themes": [{"theme": "ttheme, "topic": "ttopic", "sentiment": "positive"}]}]`
	filename := "_testreviews"

	// write contents to file
	ioutil.WriteFile(filename, []byte(fileContents), 0666)

	// Checks if code runs without calling its error/breaking
	reviews := retreiveReviews(filename)

	// check if themes contains testTheme
	if len(reviews) != 1 {
		t.Errorf("Did not load themes into slice")
	} else if reviews[0].Text != "Its so bad" {
		t.Errorf("Json did not load correctly, loaded %v but expected %v", reviews, fileContents)
	}

	os.Remove("_testreviews")
}
