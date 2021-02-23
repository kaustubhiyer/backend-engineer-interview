package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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
	themes       []theme
	netSentiment int
	mentions     int
	pMentions    int
	nMentions    int
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
	fmt.Println(themes)

	return themes
}

// Groups a list of themes by topic, collating statistics about them
func groupByTopic(tSlice []theme) map[string]topic {
	topicSlice := map[string]topic{}

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
				themes:       []theme{t},
				mentions:     t.Mentions,
				pMentions:    pMentions,
				nMentions:    nMentions,
				netSentiment: int(float64(pMentions)/float64(t.Mentions)*100) - int(float64(nMentions)/float64(t.Mentions)*100),
			}
		}
	}

	return topicSlice
}
