package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/emzola/religio/data"
)

type quran struct {
	chapter int
	verse   []int
	lang    string
}

func (q quran) apiUrl(langId string) string {
	var url string
	var limit int

	switch {
	case len(q.verse) == 1:
		limit = 1
		url = fmt.Sprintf("https://api.alquran.cloud/v1/surah/%d/%s?offset=%d&limit=%d", q.chapter, langId, q.verse[0]-1, limit)
	case len(q.verse) == 2:
		limit = (q.verse[1] + 1) - q.verse[0]
		url = fmt.Sprintf("https://api.alquran.cloud/v1/surah/%d/%s?offset=%d&limit=%d", q.chapter, langId, q.verse[0]-1, limit)
	default:
		url = fmt.Sprintf("http://api.alquran.cloud/v1/surah/%d/%s", q.chapter, langId)
	}

	return url
}

func (q quran) langId(client *http.Client) (string, error) {
	langId := "en.asad"

	if len(q.lang) != 0 {
		langEditions, err := data.LanguageEditionRequest(client)
		if err != nil {
			return "", err
		}

		langId, err = data.LanguageIdentifier(langEditions, q.lang)
		if err != nil {
			return "", err
		}
	}

	return langId, nil
}

func (q *quran) extractChapterAndVerse(passage string) error {
	passage = strings.TrimSpace(passage)

	// Extract chapter
	if !strings.Contains(passage, ":") {
		q.chapter = getQuranChapterNumber(passage)
		return nil
	}

	passageParts := strings.Split(passage, ":")
	q.chapter = getQuranChapterNumber(passageParts[0])

	// Extract verse
	if !strings.Contains(passageParts[1], "-") {
		number, err := strconv.Atoi(passageParts[1])
		if err != nil {
			return err
		}
		q.verse = append(q.verse, number)
		return nil
	}

	verseParts := strings.Split(passageParts[1], "-")
	for _, verse := range verseParts {
		number, err := strconv.Atoi(verse)
		if err != nil {
			return err
		}
		q.verse = append(q.verse, number)
	}

	return nil
}
