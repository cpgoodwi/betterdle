package word

import (
	"bufio"
	"fmt"
	"os"
)

type WordService struct {
	filePath string
}

func NewWordService(filePath string) *WordService {
	return &WordService{
		filePath: filePath,
	}
}

func (s *WordService) GetWordFromLine(lineNumber int) (string, error) {
	file, err := os.Open(s.filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentLine := 1

	for scanner.Scan() {
		if currentLine == lineNumber {
			return scanner.Text(), nil
		}
		currentLine++
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	return "", fmt.Errorf("line %d not found", lineNumber)
}
