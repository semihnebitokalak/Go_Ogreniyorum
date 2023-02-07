package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Hello, world!")
	Toplama()
	Çikartma()
	Üsalma()
}
func Toplama() {

	var sayi1, sayi2 int
	fmt.Println("toplama yapmak istediğiniz birinci sayiyi giriniz:")
	fmt.Scan(&sayi1)
	fmt.Println("toplama yapmak istediğiniz ikinci sayiyi giriniz:")
	fmt.Scan(&sayi2)

	fmt.Println("toplamanin sonucu :", sayi1+sayi2)

}

func Çikartma() {

	var sayi1, sayi2 int
	fmt.Println("çikartma yapmak istediğiniz birinci sayiyi giriniz")
	fmt.Scan(&sayi1)
	fmt.Println("çikartma yapmak istediğiniz ikinci sayiyi giriniz")
	fmt.Scan(&sayi2)
	fmt.Println("çikartmanin sonucu: ", sayi1-sayi2)

}
func Üsalma() {
	c := math.Pow10(2) //burda 10 sayisinin üssünü aldik.
	fmt.Println(c)
}
