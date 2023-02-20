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
	değişkenler()

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

func değişkenler() {
	var name string // name := Semih
	name = "Semih"

	var age int // age := 18
	age = 18

	var uStudent bool // uStudent := true
	uStudent = true

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(uStudent)
	/*
		     var name string
			 fmt.Println(name)       bunun çıktısı bana hiç bir şey vermez.

			 var age int
			 fmt.Println(age)       bunun çıktısı bana 0 değerini verir.

			 Note = string için zero value değeri boştur, integer için zero value değeri 0 dır,bool'un zero value değeri "false"dur.




	*/

	/*
		fmt.Printf("%T", name)
		fmt.Printf("%T", age)
		fmt.Printf("%T", uStudent)    Printf ile atadığımız değişkenlerin türlerini öğrenebiliriz.

	*/

	// option + shift + a komutu seçtiğimiz kodları yorum satırına alır.

}
