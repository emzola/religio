package cmd

import (
	"bytes"
	"errors"
	"testing"
)

func TestParseBible(t *testing.T) {
	usageMessage := "\nbible: A client for reading The Holy Bible.\n\t\nbible: [options] passage\n\noptions: \n  -language string\n    \tBible language (format: ENG, SPA, DEU) (default \"ENG\")\n  -version string\n    \tBible version (format: KJV, ESV, NKJV, etc) (default \"ESV\")\n"
	testCases := []struct {
		name     string
		args     []string
		expected string
		err      error
	}{
		{"ParseHelpArgument", []string{"-h"}, usageMessage, errors.New("flag: help requested")},
		{"ParseInvalidArgument", []string{"-foo"}, "flag provided but not defined: -foo\n" + usageMessage, ErrInvalidCommand},
		{"ParseValidArgument", []string{"-language", "ENG", "-version", "ESV", "John 3:16"}, "+----------------------------------------------------------------------------------------------------+\n*** JOHN 3:16 ***\n+----------------------------------------------------------------------------------------------------+\n(16) “For God so loved the world, that he gave his only Son, that whoever believes in him should not perish but have eternal life.\n", nil},
		{"ParseInvalidPositionalArgument", []string{"-language", "ENG", "-version", "ESV", "2 Peter 3:1-5", "John 3:16"}, "", ErrInvalidPassageSpecified},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer
			if err := ParseBible(&buffer, tc.args); err != nil && tc.err == nil {
				t.Fatal(err)
			}
			res := buffer.String()
			if tc.expected != res {
				t.Errorf("Expected %q, got %q instead\n", tc.expected, res)
			}
		})
	}
}
