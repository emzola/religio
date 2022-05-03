package cmd

import (
	"flag"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/emzola/religio/data"
)

type config struct {
	chapter int64
	verse []int64
}

// setChapterAndVerse retrieves the quran chapter and verse and
// associates them with the appropriate fields in the quran config.
func setChapterAndVerse(c *config, scripture string) error {
	// check whether scripture contains chapter and verse. If it contains both,
	// retrieve both, if it contains only chapter, retrieve only chapter
	if !strings.Contains(scripture, ":") {
		// retrieve only chapter
		c.chapter = getChapterNumber(scripture)
	} else {
		parts := strings.Split(scripture, ":")
		// retrieve chapter
		chapter := getChapterNumber(parts[0])
		c.chapter = chapter
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

// getUrl determines which API url endpoint to make a request to.
func getUrl(c *config, identifier string) string {
	var url string
	var limit int64
	switch {
		case len(c.verse) == 1:
			limit = 1
			url = fmt.Sprintf("https://api.alquran.cloud/v1/surah/%d/%s?offset=%d&limit=%d", c.chapter, identifier, c.verse[0]-1, limit)
		case len(c.verse) == 2:
			limit = (c.verse[1] + 1) - c.verse[0]
			url = fmt.Sprintf("https://api.alquran.cloud/v1/surah/%d/%s?offset=%d&limit=%d", c.chapter, identifier, c.verse[0]-1, limit)
		default:
			url = fmt.Sprintf("http://api.alquran.cloud/v1/surah/%d/%s", c.chapter, identifier)
	}
	return url
}

// QuranCommand implements the quran sub-command.
func QuranCommand(w io.Writer, args []string) error {
	var language string
	c := &config{}
	fs := flag.NewFlagSet("quran", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&language, "lang", "en", "Language of Quran")
	fs.Usage = func() {
		var usageString = `
	quran: a sub-command for reading the quran.
	
	quran: [options] scripture`
		fmt.Fprintln(w, usageString)
		fmt.Fprintln(w)
		fmt.Fprintln(w)
		fmt.Fprintln(w, "options: ")
		fs.PrintDefaults()
	}
	fs.Parse(args)
	if fs.NArg() != 1 {
		return InvalidInputError{ErrNoScripture}
	}
	scripture := fs.Arg(0)
	// Set the chapter and verse of the config
	err := setChapterAndVerse(c, scripture)
	if err != nil {
		return err
	}	
	httpClient := httpClient()

	// Get language edition and identifier
	var identifier string
	if len(language) != 0 {
		languageEditions, err := data.LanguageEditionRequest(httpClient)
		if err != nil {
			return err
		}
		identifier, err = data.LanguageIdentifier(languageEditions, language)
		if err != nil {
			return err
		}
	} else {
		identifier = "en.asad"
	}

	// Send HTTP requests to Quran API
	url := getUrl(c, identifier)	
	quran, err := data.SendHTTPRequest(*httpClient, url)
	if err != nil {
		return err
	}
	quranData := quran.Data.Verse

	fmt.Fprintf(w, "+%s+\n", strings.Repeat("-", 100))
	fmt.Fprintf(w, "*** %s ***\n", strings.ToUpper(scripture))	
	fmt.Fprintf(w, "+%s+\n", strings.Repeat("-", 100))
	
	for _, value := range quranData {
		fmt.Fprintf(w, "(%d) %s\n", value.Number, value.Text)
	}
	return nil
}