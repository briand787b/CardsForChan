package main

import (
	"errors"
	"github.com/briand787b/validation"
)

var (
	ErrNoGameName = validation.ValidationError(errors.New("You must name a game"))
	ErrGameNotFound = validation.ValidationError(errors.New("Game not found"))
	ErrGameInactive = validation.ValidationError(errors.New("Game is inactive"))
	ErrInvitationExpired = validation.ValidationError(errors.New("Invitation has expired"))

)
