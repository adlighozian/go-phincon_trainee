package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// menentukan total file yang ingin dibuat
const totalFile = 10

// menetukan panjang konten di satu file
const contentLength = 5000

func randomString(length int) string {
	// membuat randomizer dengan unix
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	// membuat slice untuk menampung letters dengan array of char
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// membuat slice dengan panjang dari parameter length dengan tipedara alias rune
	b := make([]rune, length)

	// melakukan perulangan dari slice b, dan setiap index slice b di isi oleh value slice letters yang sudah di randomize
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}

	// mengembalikan nilai b yang sudah diubah menjadi string
	return string(b)
}

func generateFiles() {
	// membuat path untuk menyimpan file
	var tempPath = filepath.Join("./storage")

	// menghapus semua file yang ada di tempPath
	os.RemoveAll(tempPath)

	os.MkdirAll(tempPath, os.ModePerm)

	// perulangan for untuk membuat jumlah total file yang sudah ditentukan dari totalFile
	for i := 1; i <= totalFile; i++ {

		// membuat file di path tempPath dengan nama dari Sprintf
		filename := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", i))

		// membuat variabel baru untuk menampung random content yang sudah dibuat sebelumnnya di functioan randomString, dan membawa param total length konten yang ingin dibuat
		content := randomString(contentLength)

		
		err := ioutil.WriteFile(filename, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("Error writing file", filename)
		}

		if i%100 == 0 && i > 0 {
			log.Println(i, "files created")
		}
		time.Sleep(1 * time.Second)
	}

	log.Printf("%d of total files created", totalFile)
}

func main() {
	log.Println("start")
	start := time.Now()

	generateFiles()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}
