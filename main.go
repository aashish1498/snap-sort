package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/schollz/progressbar/v3"
)

func main() {
	f, err := os.OpenFile(".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	rootFolder := getPathFromUser("input stored images", "test", true)
	outputFolder := getPathFromUser("output sorted images", filepath.Join(rootFolder, "output"), false)

	processMediaInDirectory(rootFolder, outputFolder)
}

func processMediaInDirectory(rootFolder string, outputFolder string) {
	fileInfos, _ := os.ReadDir(rootFolder)
	log.Println(toString(len(fileInfos)) + " files found. Processing ...")
	bar := progressbar.Default(int64(len(fileInfos)))

	for _, file := range fileInfos {
		fullPath := filepath.Join(rootFolder, file.Name())
		dateTime := GetDateTime(fullPath)
		outputPath := filepath.Join(outputFolder, toString(dateTime.Year()), toString((int)(dateTime.Month())), toString(dateTime.Day()))
		Copy(fullPath, outputPath)
		bar.Add(1)
	}
}

func toString(number int) string {
	return strconv.FormatInt(int64(number), 10)
}

func getPathFromUser(pathName string, defaultValue string, validate bool) string {
	fmt.Print("Enter a path for " + pathName + " (" + defaultValue + "): ")
	path := scanText(defaultValue)
	for !exists(path) && validate {
		fmt.Print("Invalid path, please try again: ")
		path = scanText(defaultValue)
	}
	return path
}

func scanText(defaultValue string) string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		scanned := scanner.Text()
		if scanned != "" {
			return scanned
		}
	}
	return defaultValue
}
