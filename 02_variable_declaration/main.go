package main

import "fmt"

func main() {

	var name1 string = "Nebi"   // var - name of variable - static type
	fmt.Println("Hello", name1) // var name , surname string = "Semih Nebi " , "Tokalak"

	var name string     // var name string = "Semih Nebi"
	name = "Semih Nebi" // name := "Semih Nebi"
	fmt.Println(name)

	var age int // var age int = 19
	age = 19    // age := 19
	fmt.Println(age)

	var isMarried bool // var isMarried bool = true
	isMarried = false  // isMarried := true
	fmt.Println(isMarried)

}

// age := 40 değerini atadıktan sonra bu değeri değiştirebiliriz.
// age = 41 yazdığımızda değerimiz 41 olur yani en son atadığımız değer yazılır.
