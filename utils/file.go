package utils

import (
	"path/filepath"
	"strings"
)

var imageExtMap = map[string]struct{}{
	".jpg":  {},
	".jpeg": {},
	".png":  {},
}

func IsImage(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	if _, ok := imageExtMap[ext]; ok {
		return true
	}

	return false
}
