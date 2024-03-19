package fileutils

import (
	"bufio"
	"dev03/cmdutils"
	"fmt"
	"os"
	"time"
)

func GetInputFile(fileName string) (*os.File, error) {
	inputFile, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	return inputFile, nil
}

func CreateOutputFile() (*os.File, error) {
	currentTime := time.Now()
	fileName := fmt.Sprintf("./out/%s_sorted.txt", currentTime.Format("2006-01-02_15-04-05"))
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ReadLinesFromFile(file *os.File, opt *cmdutils.Options) ([]string, error) {
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
