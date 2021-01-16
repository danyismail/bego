package main

import (
	"fmt"
	"runtime"
	"time"
)

func PesenNasiPadang(jumlah int, detik time.Duration) {
	for i := 0; i < jumlah; i++ {
		time.Sleep(detik * time.Second)
		fmt.Println("Pesen nasi padang ke ", i+1)
	}
	fmt.Println("Pesen nasi padang selesai")
}

func BelanjaKePasar(jumlah int, detik time.Duration) {
	for i := 0; i < jumlah; i++ {
		time.Sleep(detik * time.Second)
		fmt.Println("Belanja ke pasar ke", i+1)
	}
	fmt.Println("Belanja ke pasar selesai")
}

func AmbilJahitan(jumlah int, detik time.Duration) {
	for i := 0; i < jumlah; i++ {
		time.Sleep(detik * time.Second)
		fmt.Println("Ambil jahitan ke", i+1)
	}
	fmt.Println("Ambil jahitan selesai")
}

func Mikir() {
	fmt.Println("..hmmm")
}

func TestTanpaGoRoutine() {
	fmt.Println("Bismillah..")
	PesenNasiPadang(3, 1)
	BelanjaKePasar(3, 1)
	AmbilJahitan(3, 1)
	fmt.Println("Tugas dah selesai Alhamdulillah..pulang ah")
}

func TestGoRoutine(detik int) {
	runtime.GOMAXPROCS(4)
	fmt.Println("Bismillah..")
	go PesenNasiPadang(3, 1)
	go BelanjaKePasar(3, 1)
	// Mikir() //Bisa disertakan sebagai varian testing untuk melihat sequence dari prosesnya
	go AmbilJahitan(3, 1)
	time.Sleep(4 * time.Second)
	/*Sleep => ini sifatnya mem-blocking baris selanjutnya ga akan dieksekusi sebelum waktu habis..,
	disini sebagai pemberi waktu bagi go routines untuk menyelesaikan semua pekerjaan nya*/
	fmt.Println("Tugas dah selesai Alhamdulillah..pulang ah")
}

func main() {
	/*
		Hipotesa saya sebuah pekerjaan yg seharusnya butuh waktu 9 detik untuk menyelesaikannya
		(asumsi 3 function masing2 melakukan print 3x dengan jeda 1 detik ==  3 x 3 x 1 detik),
		hanya memerlukan waktu 5 detik bahkan bisa kurang untuk menyelesaikannya karena dihandle oleh go routine.

		Disini jelas bahwa artinya go routine melakukan proses tersebut secara pararel, karena kalau saling menunggu / blocking maka
		1 x print butuh waktu 1 detik, sedang jumlah yg dilakukan sebanyak 9 x 1 detik,tapi nyatanya
		.:KURANG DARI 5 DETIK PEKERJAAN TERSEBUT DAPAT DISELESAIKAN:.
		Happy coding , never stop to learn :)

		Note: bahkan di waktu 4 detik dengan 1 CPU pun semua proses dapat diselesaikan artinya, 9 - 5 = 4.
		Up to 50% atau 2x lipat komputasi dilakukan lebih cepat, silahkan lakukan perbandingan dengan TestTanpaGoRoutine()
	*/
	TestGoRoutine(5)
	//TestTanpaGoRoutine()

}
