package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
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
	Date   string        `json:"date"`
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

func dateCompare(date1 string, date2 string) bool {
	t1, _ := time.Parse("02-01-2006", date1)
	t2, _ := time.Parse("02-01-2006", date2)

	return t1.After(t2)
}

func sentimentCompare(s1 float64, s2 float64) bool {
	return s1 > s2
}

func loadReviews(reviews []review, pageSize int, sortType string, filterTopic string, pageNumber int) (int, []review) {
	// Need to get the top pagesize reviews ordered in decreasing order of sorttype
	if sortType == "date" {
		sort.Slice(reviews, func(i int, j int) bool {
			return dateCompare(reviews[i].Date, reviews[j].Date)
		})
		return len(reviews) / pageSize, reviews[pageSize*(pageNumber-1) : pageSize*(pageNumber)]

	} else if sortType == "highest score" {
		sort.Slice(reviews, func(i int, j int) bool {
			return sentimentCompare(reviews[i].Score, reviews[j].Score)
		})
		return len(reviews) / pageSize, reviews[:pageSize]
	} else {
		sort.Slice(reviews, func(i int, j int) bool {
			return !sentimentCompare(reviews[i].Score, reviews[j].Score)
		})
		return len(reviews) / pageSize, reviews[:pageSize]
	}
}

func insertNth(s string, n int) string {
	var buffer bytes.Buffer
	var n1 = n - 1
	var l1 = len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n1 && i != l1 {
			buffer.WriteRune('\n')
		}
	}
	return buffer.String()
}

func displayFeedback(reviewList []review, pageNumber int, totalPages int) {
	w := tabwriter.NewWriter(os.Stdout, 10, 4, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Score\tReview\tThemes\tDate")
	for _, r := range reviewList {
		text := insertNth(r.Text, 120)
		textArray := strings.Split(text, "\n")
		fmt.Fprintf(w, "%g\t%s\t%s: %s\t%s\n", r.Score, textArray[0], r.Themes[0].Topic, r.Themes[0].Theme, r.Date)
		//print remaining members of textArray as well as remaining themes/topics
		for i, ti := range textArray {
			if i < len(r.Themes) && i > 0 {
				fmt.Fprintf(w, "\t%s\t%s: %s\t\n", ti, r.Themes[i].Topic, r.Themes[i].Theme)
			} else if i > 0 {
				fmt.Fprintf(w, "\t%s\t\t\n", ti)
			}

		}
	}
	w.Flush()
	fmt.Printf("Page: %d of %d\n", pageNumber, totalPages)
}
