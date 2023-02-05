package main

import "fmt"

func main() {

	fmt.Println("Hello, world!")
	Toplama()
	Çikartma()
}
func Toplama() {

	var sayi1, sayi2 int
	fmt.Scan(&sayi1)
	fmt.Scan(&sayi2)

	fmt.Println("toplamanin sonucu :", sayi1+sayi2)

}

func Çikartma() {

	var sayi1, sayi2 int
	fmt.Println("bir sayi giriniz")
	fmt.Scan(&sayi1)
	fmt.Println("bir sayi giriniz")
	fmt.Scan(&sayi2)
	fmt.Println("çikartmanin sonucu: ", sayi1-sayi2)

}
