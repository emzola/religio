package cmd

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/emzola/religio/apidata"
)

func ParseBible(w io.Writer, args []string) error {
	var b bible

	fs := flag.NewFlagSet("bible", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&b.lang, "lang", "ENG", "Bible language")
	fs.StringVar(&b.version, "version", "ESV", "Bible version (KJV, ESV, NKJV, etc)")
	fs.Usage = func() {
		var usageMessage = `
bible: A client for reading The Holy Bible.
	
bible: [options] passage`
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
	err := b.extractChapterAndVerse(passage)
	if err != nil {
		return err
	}

	headers := map[string]string{
		"key": os.Getenv("API_KEY"),
		"v":   "4",
	}

	client := apidata.Client(headers)
	url := b.apiUrl()
	bible, err := apidata.DecodeBibleData(client, url)
	if err != nil {
		return err
	}

	printBorder(w, passage)

	bibleData := bible.Data
	for _, data := range bibleData {
		fmt.Fprintf(w, "(%d) %s\n", data.VerseStart, data.VerseText)
	}

	return nil
}
