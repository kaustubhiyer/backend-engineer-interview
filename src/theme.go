package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

type sentiment struct {
	Negative string `json:"negative"`
	Positive string `json:"positive"`
}

type theme struct {
	Theme     string    `json:"theme"`
	Topic     string    `json:"topic"`
	Mentions  int       `json:"mentions"`
	Sentiment sentiment `json:"sentiment"`
}

type topic struct {
	id           int
	themes       []theme
	netSentiment int
	mentions     int
	pMentions    int
	nMentions    int
	selected     bool // checks if a topic is selected so it can be expanded on, default is false
}

// Retreives a list of themes from a file
func getThemes(filename string) []theme {
	themes := []theme{}

	plainText, err := ioutil.ReadFile(filename)

	if err != nil {
		// Throw and error and close program
		panic(err)
	}

	e := json.Unmarshal(plainText, &themes)

	if e != nil {
		fmt.Println(e)
	}

	return themes
}

// Groups a list of themes by topic, collating statistics about them
func groupByTopic(tSlice []theme) map[string]topic {
	topicSlice := map[string]topic{}
	i := 1

	for _, t := range tSlice {
		posSentiment, _ := strconv.ParseFloat(strings.TrimSpace(t.Sentiment.Positive), 64)
		negSentiment, _ := strconv.ParseFloat(strings.TrimSpace(t.Sentiment.Negative), 64)
		if _, ok := topicSlice[t.Topic]; ok {
			// add to topicslice's attributes the current theme
			curTopic := topicSlice[t.Topic]

			curTopic.themes = append(topicSlice[t.Topic].themes, t)
			curTopic.mentions += t.Mentions
			curTopic.pMentions += int(posSentiment * float64(t.Mentions))
			curTopic.nMentions += int(negSentiment * float64(t.Mentions))

			curTopic.netSentiment = int(float64(curTopic.pMentions)/float64(curTopic.mentions)*100) - int(float64(curTopic.nMentions)/float64(curTopic.mentions)*100)

			topicSlice[t.Topic] = curTopic
		} else {
			// create a new topic and add this theme to it for now
			pMentions := int(posSentiment * float64(t.Mentions))
			nMentions := int(negSentiment * float64(t.Mentions))
			topicSlice[t.Topic] = topic{
				id:           i,
				themes:       []theme{t},
				mentions:     t.Mentions,
				pMentions:    pMentions,
				nMentions:    nMentions,
				netSentiment: int(float64(pMentions)/float64(t.Mentions)*100) - int(float64(nMentions)/float64(t.Mentions)*100),
				selected:     false,
			}
			i++
		}
	}

	return topicSlice
}

// This function displays topics to stdout in an organized fashion
func displayTopics(topics map[string]topic) {
	w := tabwriter.NewWriter(os.Stdout, 10, 4, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "\tID\tTopic\t+Sentiment\t-Sentiment\tNetSentiment\tMentions\t")
	fmt.Fprintln(w, "\t--\t-------------\t----------\t----------\t------------\t--------\t")

	for topic, contents := range topics {
		fmt.Fprintf(w, "\t%d\t%s\t%f\t%f\t%d\t%d\t\n", contents.id, topic, float64(contents.pMentions)/float64(contents.mentions), float64(contents.nMentions)/float64(contents.mentions), contents.netSentiment, contents.mentions)
		// If this topic is selected, present the themes in it

		if contents.selected {
			for _, t := range contents.themes {
				fmt.Fprintf(w, "\t\t - %s\t%s\t%s\t%s\t%d\t\n", t.Theme, t.Sentiment.Positive, t.Sentiment.Negative, " ", t.Mentions)
			}
		}
	}
	w.Flush()
}

func selectTopic(topics map[string]topic, t string) {
	// reset everything to deselected
	for _, top := range topics {
		top.selected = false
	}
	// select topic
	top := topics[t]
	top.selected = !top.selected
	topics[t] = top
}
