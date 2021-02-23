package main

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
