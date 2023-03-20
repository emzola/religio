package cmd

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/emzola/religio/data"
)

type bible struct{}

type bibleConfig struct {
	passage string
	chapter string
	verse   []int64
	lang    string
}

func (b bible) apiUrl(langId string) string {
	return ""
}

func (b bible) langId(client *http.Client) (string, error) {
	return "", nil
}

// setBibleChapterAndVerse retrieves the bible chapter and verse and
// associates them with the appropriate fields in the bible config.
func PassageChapterAndVerse(c *bibleConfig, scripture string) error {
	scripture = strings.TrimSpace(scripture)
	// retrieve passage
	parts := strings.Split(scripture, " ")
	if strings.Count(scripture, " ") == 2 {
		c.passage = fmt.Sprintf("%s %s", parts[0], parts[1])
	} else {
		c.passage = parts[0]
	}

	// check whether scripture contains chapter and verse. If it contains both,
	// retrieve both, if it contains only chapter, retrieve only chapter
	if !strings.Contains(scripture, ":") {
		// retrieve only chapter
		c.chapter = scripture
	} else {
		parts := strings.Split(scripture, ":")
		// retrieve chapter
		c.chapter = parts[0]
		// retrieve verse
		if strings.Contains(parts[1], "-") {
			verseParts := strings.Split(parts[1], "-")
			for _, value := range verseParts {
				number, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return err
				}
				c.verse = append(c.verse, number)
			}
		} else {
			number, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return err
			}
			c.verse = append(c.verse, number)
		}
	}
	return nil
}

// bibleAPIUrl determines which API url endpoint to make a request to.
func bibleAPIUrl(c *bibleConfig, identifier string) string {
	var url string
	switch {
	case len(c.verse) == 1:
		url = fmt.Sprintf("https://bible-references.p.rapidapi.com/api/verses/%s/%s/%s?verse_start=%d&verse_end=%d", identifier, c.passage, c.chapter, c.verse[0], c.verse[0])
	case len(c.verse) == 2:
		url = fmt.Sprintf("https://bible-references.p.rapidapi.com/api/verses/%s/%s/%s?verse_start=%d&verse_end=%d", identifier, c.passage, c.chapter, c.verse[0], c.verse[1])
	default:
		url = fmt.Sprintf("https://bible-references.p.rapidapi.com/api/verses/%s/%s/%s", identifier, c.passage, c.chapter)
	}
	return url
}

// BibleCommand implements the bible sub-command.
func ParseBible(w io.Writer, args []string) error {
	c := &bibleConfig{}

	fs := flag.NewFlagSet("bible", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&c.lang, "lang", "en", "Bible language")
	fs.Usage = func() {
		var usageMessage = `
bible: a sub-command for reading the bible.
	
bible: [options] scripture`
		fmt.Fprintln(w, usageMessage)
		fmt.Fprintln(w)
		fmt.Fprintln(w, "options: ")
		fs.PrintDefaults()
	}

	fs.Parse(args)

	if fs.NArg() != 1 {
		return ErrInvalidPassageSpecified
	}

	passage := fs.Arg(0)

	// Set the passage, chapter and verse of the config
	err := PassageChapterAndVerse(c, passage)
	if err != nil {
		return err
	}

	httpClient := data.Client()

	// Get language edition and identifier
	var identifier string
	if len(c.lang) != 0 {
		bibleLanguage, err := data.BibleLanguageRequest(httpClient)
		if err != nil {
			return err
		}
		identifier, err = data.BibleLanguageIdentifier(bibleLanguage, c.lang)
		if err != nil {
			return err
		}
	} else {
		identifier = "kjv"
	}

	// Send HTTP requests to Bible API
	url := bibleAPIUrl(c, identifier)
	bible, err := data.SendBibleHTTPRequest(*httpClient, url)
	if err != nil {
		return err
	}
	bibleData := bible.Data[0].Verse

	fmt.Fprintf(w, "+%s+\n", strings.Repeat("-", 100))
	fmt.Fprintf(w, "*** %s ***\n", strings.ToUpper(passage))
	fmt.Fprintf(w, "+%s+\n", strings.Repeat("-", 100))

	for _, value := range bibleData {
		fmt.Fprintf(w, "(%d) %s\n", value.Number, value.Text)
	}
	return nil
}
