package cmd

import (
	"flag"
	"fmt"
	"io"

	"github.com/emzola/religio/apidata"
)

func ParseQuran(w io.Writer, args []string) error {
	var q quran

	fs := flag.NewFlagSet("quran", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&q.lang, "language", "en", "Quran language (format: en, es)")
	fs.Usage = func() {
		var usageMessage = `
quran: A client for reading The Holy Quran.
	
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

	client := apidata.Client(nil)
	langId, err := q.langId(client)
	if err != nil {
		return err
	}

	url := q.apiUrl(langId)
	quran, err := apidata.DecodeQuranData(client, url)
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
