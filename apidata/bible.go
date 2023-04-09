package apidata

import (
	"encoding/json"
	"net/http"
)

type Bible struct {
	Data []struct {
		BookId     string `json:"book_id"`
		Chapter    int    `json:"chapter"`
		VerseStart int    `json:"verse_start"`
		VerseEnd   int    `json:"verse_end"`
		VerseText  string `json:"verse_text"`
	} `json:"data"`
}

func DecodeBibleData(client *http.Client, url string) (Bible, error) {
	bible := Bible{}

	data, err := fetchRemoteResource(client, url)
	if err != nil {
		return bible, err
	}

	err = json.Unmarshal(data, &bible)
	return bible, err
}
