package main

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

func GetDateTime(filename string) time.Time {
	dateTime, err := getDatetimeFromTag(filename)
	if err != nil || dateTime.Year() < 1970 {
		slog.Warn("Could not get get exif tags, using time from file ...")
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
		slog.Error(err.Error())
	}
	var datetime = time.Now()
	x, err := exif.Decode(f)
	if err == nil {
		datetime, _ = x.DateTime()
	}
	return datetime, err
}

func IsHiddenFile(filename string) bool {
	return filename[0] == '.'
}

func Copy(srcpath, dstFolder string) (err error) {
	r, err := os.Open(srcpath)
	if err != nil {
		return err
	}
	defer r.Close()
	filename := filepath.Base(srcpath)
	os.MkdirAll(dstFolder, os.ModePerm)
	w, err := os.Create(filepath.Join(dstFolder, filename))
	if err != nil {
		return err
	}

	defer func() {
		if c := w.Close(); c != nil && err == nil {
			err = c
		}
	}()

	_, err = io.Copy(w, r)
	return err
}
