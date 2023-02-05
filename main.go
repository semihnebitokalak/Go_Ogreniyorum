package main

import "fmt"

func main() {

	fmt.Println("Hello, world!")
	Toplama()
}
func Toplama() {

	var sayi1, sayi2 int
	fmt.Scan(&sayi1)
	fmt.Scan(&sayi2)

	fmt.Println(sayi1 + sayi2)

}
