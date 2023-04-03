package belajar

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"
)

const timeoutDuration = 3 * time.Second
const numOfWorkers = 5
const buffer = 5

var wgroup = new(sync.WaitGroup)

func test2() {
	// Log that the program is starting.
	log.Println("start")

	// Record the start time.
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	namaSiswa := "muhammad"

	log.Println(len(daftarNama2), "nama siswa yang terdaftar")
	daftar := cariDaftarSiswaWithContext(ctx, namaSiswa, daftarNama2...)
	log.Println("daftar =", daftar)

	// Calculate the duration of the file generation process.
	duration := time.Since(start)

	// Log that the program has finished, along with the duration of the file generation process.
	log.Println("selesai dalam", duration, "detik")
}

func cariDaftarSiswaWithContext(ctx context.Context, namaSiswa string, daftarSiswa ...string) []string {
	var daftarSiswaFound []string

	done := make(chan int)

	go func() {
		chanDaftarSiswa := inputDataToChan(ctx, daftarSiswa)

		chanDaftarSiswaFound := findNamaSiswa(ctx, chanDaftarSiswa, namaSiswa, numOfWorkers)

		// track and print output
		counterSuccess := 0
		for result := range chanDaftarSiswaFound {
			counterSuccess++
			daftarSiswaFound = append(daftarSiswaFound, result)
		}

		done <- counterSuccess
	}()

	select {
	case <-ctx.Done():
		log.Println("finding stopped, %s", ctx.Err())
		return []string{}
	case counterTotal := <-done:
		// Log the total number of files created.
		log.Printf("%d of total found", counterTotal)
		return daftarSiswaFound
	}
}

func inputDataToChan(ctx context.Context, daftarSiswa []string) <-chan string {
	chanOut := make(chan string, buffer)

	go func() {
		for _, siswa := range daftarSiswa {
			select {
			case <-ctx.Done():
				break
			default:
				chanOut <- siswa
			}
		}
		close(chanOut)
	}()

	return chanOut
}

func findNamaSiswa(ctx context.Context, chanIn <-chan string, namaSiswa string, numOfWorkers int) <-chan string {
	chanOut := make(chan string, buffer)

	// Create a new WaitGroup to control the workers.
	wgroup := new(sync.WaitGroup)

	// Allocate number of workers.
	wgroup.Add(numOfWorkers)

	go func() {
		// Start a goroutine for each worker.
		for workerIndex := 0; workerIndex < numOfWorkers; workerIndex++ {
			go func(workerIndex int) {
				// Receive jobs from chanIn and process them.
				for siswa := range chanIn {
					select {
					case <-ctx.Done():
						break
					default:
						// Find nama siswa
						if strings.Contains(strings.ToLower(siswa), namaSiswa) {
							log.Println("worker", workerIndex, "found", siswa)

							chanOut <- siswa
						}
					}
				}

				// if chanIn is closed, and remaining jobs are finished, only then we mark the workers as done
				defer wgroup.Done()
			}(workerIndex)
		}
	}()

	// wait until `chanIn` closed and all workers are done, after that, close the `chanOut` channel
	go func() {
		wgroup.Wait()
		close(chanOut)
	}()

	return chanOut
}

var daftarNama2 = []string{
	"Aditya Ananta Putra",
	"ADNAN NUR JAILANI",
	"Afrila Zahra Prasetyo",
	"Aida Nirmala",
	"Amalia Rahma",
	"ANDINI DWI RIZKY PUTRI",
	"Angga Putra",
	"CAHAYA WIJAYA IJAM",
	"Dea Christina",
	"DEVITA NANDA OKTAVIA",
	"DEWI AYU LESTARI",
	"Dhevina Ananda Fitri",
	"DHEVY YANI RIZKI",
	"Dwi Mahani",
	"Eri Santosos",
	"Faiza Amalia Mahfudz",
	"FAUZIYATUNNISA",
	"Febriyanti Syabina",
	"Hafifah Nur Azmi Pratiwi",
	"Juliansyah Husien",
	"Kharisa Nur Aziza",
	"Lun Wina",
	"MAHREVA ROESTHA RAMADANIATY",
	"MUHAMMAD FADHILAH",
	"Muhammad Fatah Firdaus",
	"Muhammad Rafly Ahya",
	"Muhammad Rahmin Noor",
	"Muhammad Rivaldi",
	"Muhammad Rizky Rachmadhani",
	"Nadiah Nahdhah Islamiyah",
	"NOLA FEBRIANA SAPUTRI",
	"RAFADYAN REYHAN MARITZA WERAT",
	"Robby Satria",
	"Salshabilla Wahyu Anafi",
	"STEVI VIONI PAKULLA",
	"Winny Rahmah Nia",
	"Cintya Abrianto",
	"Tifany Fernanda Nizliandry",
	"Riski Ratnasari",
	"Novi Nurfalah",
	"Imroatun Mawaldi",
	"Umi Usra",
	"Iga Hamada",
	"Ardha Abidah",
	"Desyandi Ningsih",
	"Neva Adhitama",
	"Vinny Yahya",
	"Faradhia Annisa",
	"Dela Maulidah",
	"Abi Christalline Erdyaning",
	"Ossi Tanjung",
	"Elma Kusuma",
	"Puspa Nathania",
	"Farah Sukosulistiowani",
	"Dita Rismiarti",
	"Novita Indriany",
	"Fauziyah Octavia",
	"Putri Anindya",
	"Raelita Fachrully",
	"Shabrina Claudia",
	"Milati Izzatul Utomo",
	"Magdalena Evany Harkart",
	"Reny Salma",
	"Sulalah Noor",
	"Bytta Pratiko",
	"Hanny Sri Zahlia",
	"Muhammad Hilman Daniel Irianto",
	"Humam Susanto",
	"Juno Puspitasari",
	"Syahid Widyaningtias",
	"Rizka Mubarak",
	"Rico Ario Wahyuningtias",
	"Adisdi Imran",
	"Alvino Ramadhan",
	"Azrul Maulinda",
	"Aburachman A",
	"Ressy Syahrani",
	"Emir Nadya",
	"Hudzaifah Purba",
	"Izhar Riahdita",
	"Fatahillah Maulidah",
	"Wahyu Priyohadi Fitriani",
	"Fahmi Larassati",
	"Azmi Bayhacki",
	"Deristya Wirawan",
	"Jeremy Natasya",
	"Lukman Ervya",
	"Adi Mulyanti",
	"Farizi Sukmamuliawanty",
	"Roy Fiora",
	"Anton Apriyanto",
	"Riechal Azzad",
	"Andrilla Kasta",
	"Taufik Primanelza",
	"Rifat Bagir",
	"Ekka Saraswati",
	"Adli",
	"Gibran",
	"Rafli",
	"Luthfi",
}
