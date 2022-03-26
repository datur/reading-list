package utils

import "github.com/google/uuid"

func ValidateUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}
