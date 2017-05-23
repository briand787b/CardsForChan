package main

import (
	"errors"
	"github.com/briand787b/validation"
)

var (
	ErrNoGameName = validation.ValidationError(errors.New("You must name a game"))
	ErrInvalidPlayerName = validation.ValidationError(errors.New("Player must have name"))
)
