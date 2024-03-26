package fileutils

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func createFile() (*os.File, error) {
	currentTime := time.Now()
	fileName := fmt.Sprintf("./out/%s_wget.html", currentTime.Format("2006-01-02_15-04-05"))
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func WriteToFile(data []byte) error {
	file, err := createFile()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	_, err = writer.Write(data)
	if err != nil {
		return err
	}
	return writer.Flush()
}
