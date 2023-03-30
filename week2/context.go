package main

import (
	"context"
	"fmt"
)

// - Context merupakan sebuah data yang membawa value, sinyal cancel,
//   sinyal timeout dan sinyal deadline.
// - Context biasanya dibuat per request (misal setiap ada request
// 	 masuk ke server web melalui http request).
// - Context digunakan untuk mempermudah kita meneruskan value,
//   mengirim data request dan sinyal antar proses.
// - batalkan semua proses -> kirim sinyal ke context
// - sangat berguna dalam program bersamaan, di mana banyak goroutine
//   dapat dieksekusi secara paralel dan perlu mengoordinasikan
// 	 tindakan mereka.
// - package context ada di https://pkg.go.dev/context

// - Context direpresentasikan di dalam sebuah interface Context,
//   untuk itu kita perlu membuat struct yang sesuai dengan interface
// 	 Context
// - tidak usah khawatir, ada fungsi built-in-nya

//   - context.Background() -> context kosongan, no cancel, no timeout,
//     dan no value. biasa di main func, inisiasi func, test, atau awal
//     request
//   - context.TODO() -> sama seperti Background(), namun biasa digunakan
//     ketika blum jelas context yg ingin digunakan atau tersedia
//     (menerima parameter context)
func main() {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

// - Context menganut konsep parent dan child,
//   dimana satu parent banyak child
// - konsep ini mirip Inheritance OOP

// 								 +---+
// 				+---+   -| D |
// +---+  -| B |--/ +---+
// | A |-/ +---+  \ +---+
// +---+ \         -| E |
//        \ +---+	 +---+
// 			  -| C |-  +---+
// 				 +---+ \-| F |
// 							   +---+
// - Jika A dibatalkan maka semua child dan subchild A dibatalkan
// - Jika B dibatalkan maka semua child dan subchild B dibatalkan,
// 	namun parent B (yaitu A) tidak
// - Begitu juga dengan penyisipan data

// - Context merupakan Immutable Object, tidak bisa diubah ketika dibuat
// - Jika ingin menambah value, timeout, dan cancel, otomatis membuat child
// 	baru

// - tidak usah khawatir, untuk membuat child context ada fungsi
//   built-in-nya

// 	- context.WithCancel() -> membuat context baru dengan CancelFunc yang
// 	  dapat digunakan untuk membatalkan context.
// 	- context.WithDeadline() -> membuat context baru dengan deadline.
// 		deadline adalah waktu dimana context akan dibatalkan.
// 	- context.WithTimeout() -> mirip sama WithDeadline(), namun alih-alih
// 		menetapkan deadline absolut, WithTimeout menetapkan durasi timeout
// 		relatif. Setelah timeout berakhir, context dibatalkan
// 	- context.WithValue() -> membuat context baru dengan data key-value,
// 		data ini bisa diteruskan antar fungsi tanpa menggunakan parameter
// 		fungsi. Perlu diingat tipe data key harus sebanding (misalnya string
// 		atau int) dan value-nya harus thread-safe, untuk mencegah race
// 		conditions
