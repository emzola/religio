package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/emzola/religio/cmd"
)

var ErrInvalidSubCommand = errors.New("invalid sub-command specified")

func printUsage(w io.Writer) {
	fmt.Fprintln(w, "Usage [bible|quran] -h")
	cmd.ParseBible(w, []string{"-h"})
	cmd.ParseQuran(w, []string{"-h"})
}

func handleCommand(w io.Writer, args []string) error {
	var err error
	if len(args) < 1 {
		err = ErrInvalidSubCommand
	} else {
		switch args[0] {
		case "bible":
			err = cmd.ParseBible(w, args[1:])
		case "quran":
			err = cmd.ParseQuran(w, args[1:])
		case "-h", "-help":
			printUsage(w)
		default:
			err = ErrInvalidSubCommand
		}
	}

	if err != nil {
		if errors.Is(err, ErrInvalidSubCommand) || errors.Is(err, cmd.ErrInvalidCommand) || errors.Is(err, cmd.ErrInvalidPassageSpecified) || errors.Is(err, cmd.ErrInvalidArgs) {
			fmt.Println(w, err)
			printUsage(w)
		} else {
			fmt.Fprintln(w, err)
		}
	}

	return nil
}
