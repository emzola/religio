package cmd

import "errors"

var (
	ErrInvalidCommand          = errors.New("invalid command specified")
	ErrInvalidPassageSpecified = errors.New("invalid passage specified")
	ErrInvalidArgs             = errors.New("you have to specify only 1 scripture")
)
