package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/emzola/religio/cmd"
)

var ErrInvalidSubCommand = errors.New("invalid sub-command specified")

// printUsage displays help information.
func printUsage(w io.Writer) {
	fmt.Fprintln(w, "Religio [bible|quran] -h")
	cmd.BibleCommand(w, []string{"-h"})
	cmd.QuranCommand(w, []string{"-h"})
}

// handleCommand determines which sub-command to execute based on user input.
func handleCommand(w io.Writer, args []string) error {
	var err error
	if len(args) < 1 {
		err = cmd.InvalidInputError{Err: ErrInvalidSubCommand}
	} else {
		switch args[0] {
		case "-h", "-help":
			printUsage(w)
		case "bible":
			err = cmd.BibleCommand(w, args[1:])
		case "quran":
			err = cmd.QuranCommand(w, args[1:])	
		default:
			err = cmd.InvalidInputError{Err: ErrInvalidSubCommand}
		}
		if err != nil {
			if errors.As(err, &cmd.FlagParsingError{}) {
				fmt.Fprintln(w, err.Error())
			}
			if errors.As(err, &cmd.InvalidInputError{}) {
				printUsage(w)
			}
		}
	}
	return nil
}

func main() {
	err := handleCommand(os.Stdout, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}