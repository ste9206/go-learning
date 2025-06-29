package reader

import (
	"os"
	"strings"
)

func Read(filePath string) ([]string, error) {
	byteArray, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	content := string(byteArray)

	lines := strings.Split(content, "\n")

	return lines, err

}
