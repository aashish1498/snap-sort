package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/h2non/filetype"
	"github.com/rwcarlsen/goexif/exif"
)

func GetDateTime(filename string) time.Time {
	dateTime, err := getDatetimeFromTag(filename)
	if err != nil || dateTime.Year() < 1970 {
		log.Println(filename, "Could not get get exif tags, using time from file ...")
		dateTime, err = GetDateTimeFromFileInfo(filename)
	}
	return dateTime
}

func GetDateTimeFromFileInfo(filename string) (time.Time, error) {
	dateTime := time.Now()
	info, err := os.Stat(filename)
	if err == nil {
		dateTime = info.ModTime()
	}
	return dateTime, err
}

func getDatetimeFromTag(filename string) (time.Time, error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Println(err.Error())
	}
	var datetime = time.Now()
	x, err := exif.Decode(f)
	if err == nil {
		datetime, _ = x.DateTime()
	}
	return datetime, err
}

func IsValidType(fileBuffer []byte) bool {
	return filetype.IsImage(fileBuffer) || filetype.IsVideo(fileBuffer)

}

func Copy(srcpath, dstFolder string) (err error) {
	filename := filepath.Base(srcpath)
	input, err := os.ReadFile(srcpath)
	if err != nil || !IsValidType(input) {
		log.Println("Skipping file: " + srcpath)
		return err
	}
	os.MkdirAll(dstFolder, os.ModePerm)
	destinationFile := filepath.Join(dstFolder, filename)
	err = os.WriteFile(destinationFile, input, 0644)
	return err
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
