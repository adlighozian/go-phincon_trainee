package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type FileInfo struct {
	FilePath  string // file location
	Content   []byte // file content
	Sum       string // md5 sum of content
	IsRenamed bool   // indicator whether the particular file is renamed already or not
}

func main() {
	// ...
	log.Println("start")
	start := time.Now()
	duration := time.Since(start)

	// pipe 1: baca file
	chanFileContent := readFiles()

	// pipe 2: MD5 Hash Konten File
	chanFileSum1 := getSum(chanFileContent)
	chanFileSum2 := getSum(chanFileContent)
	chanFileSum3 := getSum(chanFileContent)
	chanFileSum := mergeChanFileInfo(chanFileSum1, chanFileSum2, chanFileSum3)

	// pipe 3: rename files
	chanRename1 := rename(chanFileSum)
	chanRename2 := rename(chanFileSum)
	chanRename3 := rename(chanFileSum)
	chanRename4 := rename(chanFileSum)
	chanRename := mergeChanFileInfo(chanRename1, chanRename2, chanRename3, chanRename4)

	// print output
	counterRenamed := 0
	counterTotal := 0
	for fileInfo := range chanRename {
		if fileInfo.IsRenamed {
			counterRenamed++
		}
		counterTotal++
	}

	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)
	log.Println("done in", duration.Seconds(), "seconds")
}

func readFiles() <-chan FileInfo {
	// membuat path untuk menyimpan file
	var tempPath = filepath.Join("./storage1")

	// membuat channel dengan tipe data struct FileInfo
	chanOut := make(chan FileInfo)

	// menjalankan goroutine
	go func() {

		// filepath.walk adalah fungsi untuk mencari file yang berada di path tempPath
		err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {

			// ngecek path
			if err != nil {
				log.Println(err.Error())
				return err
			}

			// mengecek apakah info itu directory folder atau bukan, jika itu folder maka akan error
			if info.IsDir() {
				return nil
			}

			// untuk membaca file dan mengembalikan contentnya
			buf, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			// mengisi channel chanin dengan struct
			chanOut <- FileInfo{
				FilePath: path,
				Content:  buf,
			}

			return nil
		})
		if err != nil {
			log.Println("ERROR:", err.Error())
		}

		close(chanOut)
	}()

	return chanOut
}

func getSum(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for fileInfo := range chanIn {
			fileInfo.Sum = fmt.Sprintf("%x", md5.Sum(fileInfo.Content))
			chanOut <- fileInfo
		}
		close(chanOut)
	}()

	return chanOut
}

func mergeChanFileInfo(chanInMany ...<-chan FileInfo) <-chan FileInfo {
	wg := new(sync.WaitGroup)
	chanOut := make(chan FileInfo)

	wg.Add(len(chanInMany))
	for _, eachChan := range chanInMany {
		go func(eachChan <-chan FileInfo) {
			for eachChanData := range eachChan {
				chanOut <- eachChanData
			}
			wg.Done()
		}(eachChan)
	}

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

func rename(chanIn <-chan FileInfo) <-chan FileInfo {
	// membuat path untuk menyimpan file
	var tempPath = filepath.Join("./storage")
	chanOut := make(chan FileInfo)

	go func() {
		for fileInfo := range chanIn {
			newPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", fileInfo.Sum))
			err := os.Rename(fileInfo.FilePath, newPath)
			fileInfo.IsRenamed = err == nil
			chanOut <- fileInfo
		}

		close(chanOut)
	}()

	return chanOut
}
