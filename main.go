package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
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
	for _, file := range fileInfos {
		fullPath := filepath.Join(rootFolder, file.Name())
		dateTime := GetDateTime(fullPath)
		outputPath := filepath.Join(outputFolder, toString(dateTime.Year()), toString((int)(dateTime.Month())), toString(dateTime.Day()))
		Copy(fullPath, outputPath)
	}
}

func toString(number int) string {
	return strconv.FormatInt(int64(number), 10)
}

func getPathFromUser(pathName string, defaultValue string, validate bool) string {
	path := defaultValue
	fmt.Print("Enter a path for " + pathName + " (" + defaultValue + "): ")
	fmt.Scanln(&path)
	for !exists(path) && validate {
		fmt.Print("Invalid path, please try again: ")
		fmt.Scanln(&path)
	}
	return path
}
