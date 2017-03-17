package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

type WatsonConnection struct {
	URL      string
	Version  string
	Username string
	Password string
}

func (w WatsonConnection) getToneAnalysis(text string) WatsonToneResponse {
	client := &http.Client{}
	req, err := http.NewRequest("GET", w.URL+"?version="+w.Version+"&text="+url.QueryEscape(text), nil)
	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth(w.Username, w.Password)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	data := WatsonToneResponse{}
	json.Unmarshal(body, &data)

	return data
}

func hasBadWords(text string) bool {
	matched, err := regexp.MatchString(`(?i)(\s|^)f..k(\s|$)`, text)
	if err != nil {
		fmt.Println(err)
	}
	return matched
}

func calculatePriority(w WatsonConnection, text string) int {

	if hasBadWords(text) {
		fmt.Println("Detected no-no words, lowest priority.")
		return 5
	}

	var factorsConsidered, normalizedScore int
	var totalScore float64

	watsonData := w.getToneAnalysis(text)

	for _, toneCategory := range watsonData.DocumentTone.ToneCategories {
		for _, tone := range toneCategory.Tones {
			if tone.Score > 0 {
				switch tone.ID {
				case
					"conscientiousness_big5",
					"agreeableness_big5",
					"analytical":
					factorsConsidered++
					totalScore += tone.Score
					if tone.Score > 0.75 {
						fmt.Printf("%s (positive): high probability.\n", tone.Name)
					}
				case
					"anger",
					"disgust",
					"emotional_range_big5":
					factorsConsidered++
					totalScore += 1 - tone.Score
					if tone.Score > 0.75 {
						fmt.Printf("%s (negative): high probability.\n", tone.Name)
					}
				}
			}
		}
	}

	normalizedScore = int(5 - math.Floor(5*totalScore/float64(factorsConsidered)))
	return normalizedScore
}

func main() {
	watson := WatsonConnection{
		"https://gateway.watsonplatform.net/tone-analyzer/api/v3/tone",
		"2016-05-19",
		os.Getenv("WATSON_USERNAME"),
		os.Getenv("WATSON_PASSWORD"),
	}

	analysis := calculatePriority(watson, os.Args[1])
	fmt.Printf("Priority: %d.\n", analysis)
}
