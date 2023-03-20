package apidata

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Bible struct {
	Data []struct {
		Bible   string       `json:"bible"`
		Passage string       `json:"passage"`
		Chapter string       `json:"chapter"`
		Verse   []BibleVerse `json:"verse"`
	} `json:"data"`
}

type BibleVerse struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

type BibleLanguage struct {
	Data []Languages `json:"data"`
}

type Languages struct {
	Bible     string   `json:"bible"`
	Languages []string `json:"languages"`
}

// BibleLanguageRequest sends an HTTP request to the Quran Language Edition API endpoint.
func BibleLanguageRequest(client *http.Client) (BibleLanguage, error) {
	bibles := BibleLanguage{}
	url := "https://bible-references.p.rapidapi.com/api/bibles"
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return bibles, err
	}
	req.Header.Add("X-RapidAPI-Host", "bible-references.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "fc6060058dmsh115e1a12702b12bp1d48ccjsn679feaca9662")
	resp, err := client.Do(req)
	if err != nil {
		return bibles, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return bibles, err
	}
	err = json.Unmarshal(data, &bibles)
	return bibles, err
}

// LanguageIdentifier gets the language identifier from BibleLanguage.
func BibleLanguageIdentifier(bibles BibleLanguage, language string) (string, error) {
	var identifier string
	var isValidLanguage bool
	data := bibles.Data
	for _, value := range data {
		if value.Languages[0] == language {
			identifier = value.Bible
			isValidLanguage = true
			break
		}
	}
	if !isValidLanguage {
		return identifier, errors.New("language not found")
	}
	return identifier, nil
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
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return bible, err
	}
	err = json.Unmarshal(data, &bible)
	return bible, err
}
