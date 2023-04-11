package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"readAndWriteCsv/helper"
	"readAndWriteCsv/model"
	"sync"
	"time"
)

var wg = new(sync.WaitGroup)

func main() {
	jumlahData := 1000000
	jumlahWorker := 1

	start := time.Now()

	wg.Add(1)
	go func() {
		chanFile := insertToChannel(jumlahData, jumlahWorker)
		insertToCsv(chanFile, jumlahWorker)
		wg.Done()
	}()
	wg.Wait()

	csvfile2, err := os.Open("file/social.csv")
	if err != nil {
		log.Panicln("Error open data", err)
	}
	csvreader := csv.NewReader(csvfile2)
	ReadAllData(csvreader)

	duration := time.Since(start)
	fmt.Println(duration)

}

func insertToChannel(jumlahData int, worker int) <-chan model.Social {

	chanOut := make(chan model.Social, worker)

	go func() {
		var sliceName = []string{"ulil", "bagas", "nandy", "fian", "siapa", "ini", "itu"}

		for i := 0; i < jumlahData; i++ {
			chanOut <- model.Social{
				User:    sliceName[rand.Intn(len(sliceName))],
				Comment: helper.RandomString(30),
			}
		}
		close(chanOut)
	}()

	return chanOut
}

func insertToCsv(chanIn <-chan model.Social, worker int) {
	numberOfWorkers := worker
	wg.Add(numberOfWorkers)

	go func() {
		fmt.Println("start with worker", numberOfWorkers)
		csvfile, err := os.Create("file/social.csv")
		if err != nil {
			log.Panicln("Error create data", err)
		}
		csvwriter := csv.NewWriter(csvfile)
		for worker := 1; worker <= numberOfWorkers; worker++ {
			go func(chanwork <-chan model.Social, i int) {
				var dataTempSlice [][]string
				defer wg.Done()
				for v := range chanwork {
					var temp = []string{v.User, v.Comment}
					dataTempSlice = append(dataTempSlice, temp)
					// fmt.Println("worker :", i, temp)
				}
				if err := csvwriter.WriteAll(dataTempSlice); err != nil {
					log.Fatalln("Error writing data", err)
				}
			}(chanIn, worker)
		}
		defer csvwriter.Flush()
		defer func() {
			err := csvfile.Close()
			if err != nil {
				log.Panicln("Error closing file", err)
			}
		}()

		wg.Wait()
	}()
}

func ReadAllData(csvreader *csv.Reader) {
	number := 1
	if booksDataSlice, err := csvreader.ReadAll(); err != nil {
		log.Fatalln("Error cant read data from csv")
	} else {
		for _, v := range booksDataSlice {
			fmt.Println(number+1, "User :", v[0], "| Comment :", v[1])
			number++

			if number == 100 {
				break
			}
		}
	}
}
