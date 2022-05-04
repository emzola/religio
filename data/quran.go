package data

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Quran struct {
	Data struct {
		Verse []QuranVerse `json:"ayahs"`
	} `json:"data"`
}

type QuranVerse struct {
	Number int    `json:"numberInSurah"`
	Text   string `json:"text"`
}

type LanguageEdition struct {
	Data []DataDetails `json:"data"`
}

type DataDetails struct {
	Identifier string `json:"identifier"`
	Language   string `json:"language"`
}

// LanguageEditionRequest sends an HTTP request to the Quran Language Edition API endpoint.
func LanguageEditionRequest(client *http.Client) (LanguageEdition, error) {
	edition := LanguageEdition{}
	url := "http://api.alquran.cloud/v1/edition"
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return edition, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return edition, err
	}
	defer resp.Body.Close()
	if resp.Header.Get("Content-Type") != "application/json" {
		return edition, nil
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return edition, err
	}
	err = json.Unmarshal(data, &edition)
	return edition, err
}

// LanguageIdentifier gets the language identifier from LanguageEdition.
func LanguageIdentifier(editions LanguageEdition, language string) (string, error) {
	var identifier string
	var isValidLanguage bool
	data := editions.Data
	for _, value := range data {
		if value.Language == language {
			identifier = value.Identifier
			isValidLanguage = true
			break
		}
	} 
	if !isValidLanguage {
		return identifier, errors.New("language not found")
	}
	return identifier, nil
}

// SendHTTPRequest sends an HTTP request to the Quran Surah endpoint.
func SendQuranHTTPRequest(client http.Client, url string) (Quran, error) {
	quran := Quran{}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return quran, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return quran, err
	}
	defer resp.Body.Close()
	if resp.Header.Get("Content-Type") != "application/json" {
		return quran, nil
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return quran, err
	}
	err = json.Unmarshal(data, &quran)
	return quran, err
}