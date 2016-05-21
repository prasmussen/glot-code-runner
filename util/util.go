package util

import (
	"path/filepath"
)

func FilterByExtension(files []string, ext string) []string {
	var newFiles []string
	suffix := "." + ext

	for _, file := range files {
		if filepath.Ext(file) == suffix {
			newFiles = append(newFiles, file)
		}
	}

	return newFiles
}
