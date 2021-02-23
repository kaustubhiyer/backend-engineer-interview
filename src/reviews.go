package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type themeReview struct {
	Theme     string `json:"theme"`
	Topic     string `json:"topic"`
	Sentiment string `json:"sentiment"`
}

type review struct {
	Text   string        `json:"text"`
	Themes []themeReview `json:"themes"`
	Score  float64       `json:"score"`
}

func retreiveReviews(filename string) []review {
	reviews := []review{}

	plainText, err := ioutil.ReadFile(filename)

	if err != nil {
		// Throw and error and close program
		panic(err)
	}

	e := json.Unmarshal(plainText, &reviews)

	if e != nil {
		fmt.Println(e)
	}
	return reviews
}
