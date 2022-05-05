package cmd

import "errors"

var (
	ErrInvalidCommand = errors.New("invalid command specified")
	ErrNoScripture = errors.New("you have to specify a scripture")
	ErrInvalidArgs = errors.New("you have to specify only 1 scripture")
)

type InvalidInputError struct {
	Err error
}

func (e InvalidInputError) Error() string {
	return e.Err.Error()
}

type FlagParsingError struct {
	Err error
}

func (e FlagParsingError) Error() string {
	return e.Err.Error()
}