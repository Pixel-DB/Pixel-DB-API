package utils

import "github.com/google/uuid"

func GenerateFilename(OldFilename, Username string) string {
	Filename := uuid.New().String() + "-" + Username + "-" + OldFilename

	return Filename
}
