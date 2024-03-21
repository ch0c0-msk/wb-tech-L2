package fileutils

import (
	"bufio"
	"os"
)

func GetInputFile(fileName string) (*os.File, error) {
	if len(fileName) == 0 {
		return os.Stdin, nil
	}
	inputFile, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	return inputFile, nil
}

func ReadLinesFromFile(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func WriteLinesToFile(file *os.File, lines []string) error {
	for _, line := range lines {
		if _, err := file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}
