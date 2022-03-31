package main

import "fmt"

func angka_kata(x int) {
	satudigit := [10]string{"Nol ", "Satu ", "Dua ", "Tiga ", "Empat ", "Lima ", "Enam ", "Tujuh ", "Delapan ", "Sembilan "}
	duadigit := [10]string{"Sepuluh ", "Sebelas ", "Duabelas ", "Tigabelas ", "Empatbelas ", "Limabelas ", "Enambelas ", "Tujuhbelas ", "Delapanbelas ", "Sembilanbelas "}
	banyak := [5]string{"Puluh ", "Seratus ", "Ratus ", "Ribu ", "Seribu "}
	var kata string
	kata = ""
	for x > 0 {
		if x >= 100000 {
			if (x/1000)%100 == 0 {
				if x/100000 != 1 {
					kata = kata + satudigit[x/100000] + banyak[2] + banyak[3]
				} else {
					kata = kata + banyak[1] + banyak[3]
				}
			} else {
				if x/100000 != 1 {
					kata = kata + satudigit[x/100000] + banyak[2]
				} else {
					kata = kata + banyak[1]
				}
			}
			x = x % 100000
		} else if x >= 10000 {
			if x/1000 >= 20 {
				if (x/1000)%10 != 0 {
					kata = kata + satudigit[x/10000] + banyak[0] + satudigit[(x/1000)%10] + banyak[3]
				} else {
					kata = kata + satudigit[x/10000] + banyak[0] + banyak[3]
				}
			} else if x/1000 >= 10 {
				kata = kata + duadigit[(x/1000)-10] + banyak[3]
			}
			x = x % 1000
		} else if x >= 1000 {
			if x/1000 != 1 {
				kata = kata + satudigit[x/1000] + banyak[3]
			} else {
				kata = kata + banyak[4]
			}
			x = x % 1000
		} else if x >= 100 {
			if x/100 != 1 {
				kata = kata + satudigit[x/100] + banyak[2]
			} else {
				kata = kata + banyak[1]
			}
			x = x % 100
		} else if x >= 20 {
			if x%10 != 0 {
				kata = kata + satudigit[x/10] + banyak[0] + satudigit[x%10]
			} else {
				kata = kata + satudigit[x/10] + banyak[0]
			}
			x = 0
		} else if x >= 10 {
			kata = kata + duadigit[x-10]
			x = 0
		} else {
			kata = kata + satudigit[x]
			x = 0
		}
	}
	if len(kata) == 0 {
		kata = satudigit[0]
	}
	fmt.Println(kata)
}
func main() {
	var a int
	fmt.Println("Masukkan angka yang diinginkan (angka harus kurang dari satu juta)")
	fmt.Scanln(&a)
	angka_kata(a)
}
