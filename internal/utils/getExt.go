package utils

import (
	"path"
	"strings"
)

func GetExt(Filename string) string {
	ext := strings.ToLower(strings.TrimPrefix(path.Ext(Filename), "."))
	return ext
}
