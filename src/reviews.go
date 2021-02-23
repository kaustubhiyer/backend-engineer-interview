package main

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
