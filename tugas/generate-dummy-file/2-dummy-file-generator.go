package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func proceed() {
	// membuat path untuk menyimpan file
	var tempPath = filepath.Join("./storage")

	// membuat variabel untuk menampung jumlah file yang error
	counterTotal := 0

	// membuat variabel untuk menampung jumlah file yang sudah direname
	counterRenamed := 0

	err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {

		// if there is an error, return immediatelly
		if err != nil {
			return err
		}

		// if it is a sub directory, return immediatelly
		if info.IsDir() {
			return nil
		}

		counterTotal++

		// read file
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// sum it
		sum := fmt.Sprintf("%x", md5.Sum(buf))

		// rename file
		destinationPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", sum))
		err = os.Rename(path, destinationPath)
		if err != nil {
			return err
		}

		counterRenamed++
		return nil
	})
	if err != nil {
		log.Println("ERROR:", err.Error())
	}

	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)
}

func main() {
	log.Println("start")
	start := time.Now()

	proceed()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}
