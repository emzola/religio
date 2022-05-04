package data

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Bible struct {
	Data []struct {
		Bible string `json:"bible"`
		Passage string `json:"passage"`
		Chapter string `json:"chapter"`
		Verse []BibleVerse `json:"verse"`
	} `json:"data"`
}

type BibleVerse struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

// SendHTTPRequest sends an HTTP request to the Bible Get Verses endpoint.
func SendBibleHTTPRequest(client http.Client, url string) (Bible, error) {
	bible := Bible{}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return bible, err
	}
	req.Header.Add("X-RapidAPI-Host", "bible-references.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "fc6060058dmsh115e1a12702b12bp1d48ccjsn679feaca9662")
	resp, err := client.Do(req)
	if err != nil {
		return bible, err
	}
	defer resp.Body.Close()
	if resp.Header.Get("Content-Type") != "application/json" {
		return bible, nil
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return bible, err
	}
	err = json.Unmarshal(data, &bible)
	return bible, err
}