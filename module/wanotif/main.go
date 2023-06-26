package main

import (
	"fmt"

	"github.com/pquerna/otp/totp"
)

func main() {
	// Menentukan kunci rahasia (secret key) yang akan digunakan untuk verifikasi OTP.
	secret := "KUNCI_RAHASIA_ANDA"

	// Menghasilkan URI OTP berdasarkan kunci rahasia.
	uri := totp.URI("AccountName", secret)
	fmt.Println("URI OTP:", uri)

	totp.

	// Mengirim OTP ke pengguna (contoh: melalui email atau SMS).
	sendOTPToUser("USER_EMAIL_OR_PHONE_NUMBER", uri)

	// Mengambil OTP yang dimasukkan oleh pengguna.
	userOTP := getUserInput()

	// Memverifikasi kevalidan OTP yang dimasukkan oleh pengguna.
	valid := verifyOTP(uri, userOTP)
	if valid {
		fmt.Println("OTP valid. Akses diberikan.")
	} else {
		fmt.Println("OTP tidak valid. Akses ditolak.")
	}
}

func sendOTPToUser(userContact string, uri string) {
	// Mengirim OTP ke pengguna (contoh: melalui email atau SMS).
	fmt.Printf("OTP dikirim ke: %s\n", userContact)
	fmt.Printf("Silakan masukkan OTP yang Anda terima:\n")
}

func getUserInput() string {
	// Mendapatkan input OTP dari pengguna.
	var userOTP string
	fmt.Scanln(&userOTP)
	return userOTP
}

func verifyOTP(uri string, userOTP string) bool {
	// Memverifikasi kevalidan OTP yang dimasukkan oleh pengguna.
	valid := totp.Validate(userOTP, uri)
	return valid
}
