package main

import (
	"bytes"
	"errors"
	"testing"
)

func TestHandleCommand(t *testing.T) {
	usageMessage := "Usage [bible|quran] -h\n\nbible: A client for reading The Holy Bible.\n\t\nbible: [options] passage\n\noptions: \n  -language string\n    \tBible language (format: ENG, SPA, DEU) (default \"ENG\")\n  -version string\n    \tBible version (format: KJV, ESV, NKJV, etc) (default \"ESV\")\n\nquran: A client for reading The Holy Quran.\n\t\nquran: [options] passage\n\noptions: \n  -language string\n    \tQuran language (format: en, es) (default \"en\")\n"
	bibleOutput := "+----------------------------------------------------------------------------------------------------+\n*** JOHN 3:16 ***\n+----------------------------------------------------------------------------------------------------+\n(16) “For God so loved the world, that he gave his only Son, that whoever believes in him should not perish but have eternal life.\n"
	quranOutput := "+----------------------------------------------------------------------------------------------------+\n*** AL-FATIHAH:1-5 ***\n+----------------------------------------------------------------------------------------------------+\n(1) Во имя Аллаха, Милостивого, Милосердного!\n(2) Хвала Аллаху, Господу миров,\n(3) Милостивому, Милосердному,\n(4) Властелину Дня воздаяния!\n(5) Тебе одному мы поклоняемся и Тебя одного молим о помощи.\n"

	testCases := []struct {
		name     string
		args     []string
		expected string
		err      error
	}{
		{"ParseNoArgument", []string{}, usageMessage, ErrInvalidSubCommand},
		{"ParseHelpArgument", []string{"-h"}, usageMessage, nil},
		{"ParseInvalidArgument", []string{"-foo"}, usageMessage, errors.New("flag provided but not defined: -foo\n")},
		{"ParseValidBibleArgument", []string{"bible", "-language", "ENG", "-version", "ESV", "John 3:16"}, bibleOutput, nil},
		{"ParseValidQuranArgument", []string{"quran", "-language", "ru", "Al-Fatihah:1-5"}, quranOutput, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer
			if err := handleCommand(&buffer, tc.args); err != nil && tc.err == nil {
				t.Fatal(err)
			}
			res := buffer.String()
			if tc.expected != res {
				t.Errorf("Expected %q, got %q instead\n", tc.expected, res)
			}
		})
	}
}
