package cmd

import (
	"fmt"
	"strings"
)

type bible struct {
	book    string
	chapter string
	verse   []int
	lang    string
	version string
}

func (b bible) apiUrl() string {
	var url string
	switch {
	case len(b.verse) == 1:
		url = fmt.Sprintf("https://4.dbt.io/api/bibles/filesets/%s%s/%s/%s?verse_start=%d&verse_end=%d", b.lang, b.version, b.book, b.chapter, b.verse[0], b.verse[0])
	case len(b.verse) == 2:
		url = fmt.Sprintf("https://4.dbt.io/api/bibles/filesets/%s%s/%s/%s?verse_start=%d&verse_end=%d", b.lang, b.version, b.book, b.chapter, b.verse[0], b.verse[1])
	default:
		url = fmt.Sprintf("https://4.dbt.io/api/bibles/filesets/%s%s/%s/%s", b.lang, b.version, b.book, b.chapter)
	}
	return url
}

func (b *bible) Passage(passage string) error {
	passage = strings.TrimSpace(passage)
	parts := strings.Split(passage, " ")

	if strings.Count(passage, " ") < 1 || strings.Count(passage, " ") > 2 {
		return ErrInvalidPassageSpecified
	}

	if strings.Count(passage, " ") == 2 {
		// set book
		bookId, err := bibleBookId(fmt.Sprintf("%s %s", parts[0], parts[1]))
		if err != nil {
			return err
		}
		b.book = bookId

		// set chapter
		chapter := bibleChapter(parts[2], b)

		// set verse
		err = bibleVerse(chapter[1], b)
		if err != nil {
			return err
		}
	}

	if strings.Count(passage, " ") == 1 {
		// set book
		b.book = parts[0]

		// set chapter
		chapter := bibleChapter(parts[1], b)

		// set verse
		err := bibleVerse(chapter[1], b)
		if err != nil {
			return err
		}
	}

	return nil
}
