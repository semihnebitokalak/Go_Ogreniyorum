package main

import "fmt"

func main() {

	var name, age, height, weight, isMarried, shoesnumber = "Semih", 18, 196, 85, false, 46.5

	//name,age,height,weight,isMarried = "Semih",18,196,85,false

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(height)
	fmt.Println(weight)
	fmt.Println(isMarried)
	fmt.Println(shoesnumber)

	/*
		     var name string
			 fmt.Println(name)       bunun çıktısı bana hiç bir şey vermez.

			 var age int
			 fmt.Println(age)       bunun çıktısı bana 0 değerini verir.

			 Note = string için zero value değeri boştur, integer için zero value değeri 0 dır,bool'un zero value değeri "false"dur.

	*/

}
