package utils

import "github.com/google/uuid"

func GenerateFilename(OldFilename string) string {
	Filename := uuid.New().String() + "-" + OldFilename

	return Filename
}
