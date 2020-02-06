package uuid

import (
	"github.com/google/uuid"
)

// NewID generates UUID V4
func NewID() string {
	return uuid.Must(uuid.NewRandom()).String()
}

// NewIDFor generates the same UUID (V5) for the given namespace and value
func NewIDFor(namespace string, value string) string {
	nsUUID := uuid.NewSHA1(uuid.NameSpaceOID, []byte(namespace))
	return uuid.NewSHA1(nsUUID, []byte(value)).String()
}

// ValidateID checks whether the given string is a valid RFC4122-compliant UUID
func ValidateID(s string) bool {
	id, err := uuid.Parse(s)
	return err == nil && id.Variant() == uuid.RFC4122
}
