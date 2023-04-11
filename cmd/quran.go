package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/emzola/religio/apidata"
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
		url := "http://api.alquran.cloud/v1/edition"
		lang, err := apidata.DecodeLangEdition(client, url)
		if err != nil {
			return "", err
		}

		langId, err = apidata.LanguageIdentifier(lang, q.lang)
		if err != nil {
			return "", err
		}
	}

	return langId, nil
}

func (q *quran) extractChapterAndVerse(passage string) error {
	var err error
	passage = strings.TrimSpace(passage)

	// Extract chapter
	if !strings.Contains(passage, ":") {
		q.chapter, err = quranChapterNumber(passage)
		if err != nil {
			return err
		}
		return nil
	}

	passageParts := strings.Split(passage, ":")
	q.chapter, err = quranChapterNumber(passageParts[0])
	if err != nil {
		return err
	}

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
