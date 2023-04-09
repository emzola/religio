package cmd

import (
	"fmt"
	"strconv"
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

func (b *bible) extractChapterAndVerse(passage string) error {
	passage = strings.TrimSpace(passage)
	parts := strings.Split(passage, " ")

	if strings.Count(passage, " ") == 2 {
		// set book
		bookId, err := setBibleBookId(fmt.Sprintf("%s %s", parts[0], parts[1]))
		if err != nil {
			return err
		}
		b.book = bookId
		// set chapter
		if strings.Contains(parts[2], ":") {
			chapter := strings.Split(parts[2], ":")
			b.chapter = chapter[0]
			// set verse
			if strings.Contains(chapter[1], "-") {
				verses := strings.Split(chapter[1], "-")
				for _, verse := range verses {
					number, err := strconv.Atoi(verse)
					if err != nil {
						return err
					}
					b.verse = append(b.verse, number)
				}
			} else {
				number, err := strconv.Atoi(chapter[1])
				if err != nil {
					return err
				}
				b.verse = append(b.verse, number)
			}
		} else {
			b.chapter = parts[2]
		}
	} else if strings.Count(passage, " ") == 1 {
		// set book
		b.book = parts[0]
		// set chapter
		if strings.Contains(parts[1], ":") {
			chapter := strings.Split(parts[1], ":")
			b.chapter = chapter[0]

			// set verse
			if strings.Contains(chapter[1], "-") {
				verses := strings.Split(chapter[1], "-")
				for _, verse := range verses {
					number, err := strconv.Atoi(verse)
					if err != nil {
						return err
					}
					b.verse = append(b.verse, number)
				}
			} else {
				number, err := strconv.Atoi(chapter[1])
				if err != nil {
					return err
				}
				b.verse = append(b.verse, number)
			}
		} else {
			b.chapter = parts[1]
		}
	} else {
		return ErrInvalidPassageSpecified
	}

	return nil
}
