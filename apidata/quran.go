package apidata

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Quran struct {
	Data struct {
		Verse []Verse `json:"ayahs"`
	} `json:"data"`
}

type Verse struct {
	Number int    `json:"numberInSurah"`
	Text   string `json:"text"`
}

type Language struct {
	Data []langData `json:"data"`
}

type langData struct {
	Identifier string `json:"identifier"`
	Language   string `json:"language"`
}

func DecodeQuranData(client *http.Client, url string) (Quran, error) {
	quran := Quran{}

	data, err := fetchRemoteResource(client, url)
	if err != nil {
		return quran, err
	}

	err = json.Unmarshal(data, &quran)
	return quran, err
}

func DecodeLangEdition(client *http.Client, url string) (Language, error) {
	lang := Language{}

	data, err := fetchRemoteResource(client, url)
	if err != nil {
		return lang, err
	}

	err = json.Unmarshal(data, &lang)
	return lang, err
}

func LanguageIdentifier(langEdition Language, language string) (string, error) {
	var identifier string
	var isValidLang bool

	for _, edition := range langEdition.Data {
		if edition.Language == language {
			identifier = edition.Identifier
			isValidLang = true
			break
		}
	}

	if !isValidLang {
		return identifier, errors.New("language not found")
	}

	return identifier, nil
}
