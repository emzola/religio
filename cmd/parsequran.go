package cmd

import (
	"flag"
	"fmt"
	"io"

	"github.com/emzola/religio/data"
)

func ParseQuran(w io.Writer, args []string) error {
	var q quran

	fs := flag.NewFlagSet("quran", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&q.lang, "lang", "en", "Quran language")
	fs.Usage = func() {
		var usageMessage = `
quran: a client for reading the quran.
	
quran: [options] passage`
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
	err := q.extractChapterAndVerse(passage)
	if err != nil {
		return err
	}

	client := data.Client()
	langId, err := q.langId(client)
	if err != nil {
		return err
	}

	url := q.apiUrl(langId)
	quran, err := data.SendQuranHTTPRequest(*client, url)
	if err != nil {
		return err
	}

	printBorder(w, passage)

	quranData := quran.Data.Verse

	for _, value := range quranData {
		fmt.Fprintf(w, "(%d) %s\n", value.Number, value.Text)
	}

	return nil
}
