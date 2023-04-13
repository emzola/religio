package cmd

import (
	"reflect"
	"testing"
)

func TestQuranChapterNumber(t *testing.T) {
	testCases := []struct {
		name     string
		chapter  string
		expected int
		err      error
	}{
		{"InvalidQuranChapter", "foo", 0, ErrInvalidPassageSpecified},
		{"ValidQuranChapter", "Al-Muddaththir", 74, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := quranChapterNumber(tc.chapter)
			if err != nil && tc.err == nil {
				t.Fatal(err)
			}

			if output != tc.expected {
				t.Errorf("Expected '%d', got '%d' instead\n", tc.expected, output)
			}
		})
	}
}

func TestBibleBookId(t *testing.T) {
	testCases := []struct {
		name     string
		book     string
		expected string
		err      error
	}{
		{"InvalidBibleBookId", "foo", "", ErrInvalidPassageSpecified},
		{"ValidBibleBookId", "genesis", "GEN", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := bibleBookId(tc.book)
			if err != nil && tc.err == nil {
				t.Fatal(err)
			}

			if output != tc.expected {
				t.Errorf("Expected '%s', got '%s' instead\n", tc.expected, output)
			}
		})
	}
}

func TestBibleChapter(t *testing.T) {
	testCases := []struct {
		name     string
		passage  string
		expected []string
	}{
		{"ValidBibleChapterNoVerse", "3", nil},
		{"ValidBibleChapterWithVerse", "3:16", []string{"3", "16"}},
	}

	var b bible
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := bibleChapter(tc.passage, &b)
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("Expected '%q', got '%q' instead\n", tc.expected, output)
			}
		})
	}
}
