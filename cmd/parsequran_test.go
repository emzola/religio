package cmd

import (
	"bytes"
	"errors"
	"testing"
)

func TestParseQuran(t *testing.T) {
	usageMessage := "\nquran: A client for reading The Holy Quran.\n\t\nquran: [options] passage\n\noptions: \n  -language string\n    \tQuran language (format: en, es) (default \"en\")\n"
	testCases := []struct {
		name     string
		args     []string
		expected string
		err      error
	}{
		{"ParseHelpArgument", []string{"-h"}, usageMessage, errors.New("flag: help requested")},
		{"ParseInvalidArgument", []string{"-foo"}, "flag provided but not defined: -foo\n" + usageMessage, ErrInvalidCommand},
		{"ParseValidArgument", []string{"-language", "ru", "Al-Fatihah:1-5"}, "+----------------------------------------------------------------------------------------------------+\n*** AL-FATIHAH:1-5 ***\n+----------------------------------------------------------------------------------------------------+\n(1) Во имя Аллаха, Милостивого, Милосердного!\n(2) Хвала Аллаху, Господу миров,\n(3) Милостивому, Милосердному,\n(4) Властелину Дня воздаяния!\n(5) Тебе одному мы поклоняемся и Тебя одного молим о помощи.\n", ErrInvalidPassageSpecified},
		{"ParseInvalidPositionalArgument", []string{"-language", "ru", "Al-Fatihah:1-5", "Al-Fatihah:1-5"}, "", ErrInvalidPassageSpecified},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer
			if err := ParseQuran(&buffer, tc.args); err != nil && tc.err == nil {
				t.Fatal(err)
			}
			res := buffer.String()
			if tc.expected != res {
				t.Errorf("Expected %q, got %q instead\n", tc.expected, res)
			}
		})
	}
}
