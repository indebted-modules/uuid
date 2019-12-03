package uuid

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// NewID generates UUID V4
func NewID() string {
	v4, err := uuid.NewRandom()
	if err != nil {
		log.
			Fatal().
			Err(err).
			Msg("Failed generating UUID")
	}
	return v4.String()
}

// ValidateID checks whether the given string is a valid UUID V4
func ValidateID(s string) bool {
	id, err := uuid.Parse(s)
	return err == nil && id.Variant() == uuid.RFC4122
}
