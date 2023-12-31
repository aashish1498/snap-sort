package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	outputFolder := "output"
	rootFolder := "files"
	fileInfos, _ := os.ReadDir(rootFolder)

	for _, file := range fileInfos {
		filename := file.Name()
		if IsHiddenFile(filename) {
			continue
		}
		fullPath := filepath.Join(rootFolder, filename)
		dateTime := GetDateTime(fullPath)
		outputPath := filepath.Join(outputFolder, toString(dateTime.Year()), toString((int)(dateTime.Month())), toString(dateTime.Day()))
		fmt.Println(filename + ": " + outputPath)
		Copy(fullPath, outputPath)
	}
}

func toString(number int) string {
	return strconv.FormatInt(int64(number), 10)
}
