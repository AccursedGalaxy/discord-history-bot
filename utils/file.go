package utils

import (
	"os"
)

func SaveToFile(filename string, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func GetFileReader(filename string) *os.File {
	file, _ := os.Open(filename)
	return file
} 