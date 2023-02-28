/* uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32 */

package main

import "fmt"

func main() {

	/* x := 10
	y := 10.0

	fmt.Printf("%v %T \n", x, x)
	fmt.Printf("%v %T \n", y, y)

	fmt.Println(x + int(y))

	fmt.Printf("%v %T \n", y, y) */

	// fmt.Println(x + y)   yazdığımızda hata alırız çünkü veri tipleri farklı olduğu için.

	// type conversion = type(value) => int(y) = 10.0 => 10
	// type conversion ile biz sayının ilk atanan değerini ve türünü değiştirmiyoruz.

	/* var x int8 = 10
	var y int16 = 10

	fmt.Printf(x + y )*/ // hata alırız çünkü int8 ile int16 aynı değişken değildir.
	/*
		x := 10
		y := 10.0

		fmt.Printf("%v %T \n", x, x)
		fmt.Printf("%v %T \n", y, y)

		fmt.Println(float64(x) + y)

		fmt.Printf("%v %T \n", x, x)
		fmt.Printf("%v %T \n", y, y) */

	// eğer sonucumuzun integer çıkmasını istiyorsak float olan kısmı integera dönüştürürüz.Diğer durum için ise tam tersini yaparız.
	/*
		var x int16 = 127
		var y int8
		y = int8(x)
		fmt.Println(y)
		// y = x   //cannot use x (variable of type int8) as type int16 in assignment */
	/*
		var x int16 = 128
		var y int8

		y = int8(x)

		fmt.Println(y) // int8 değeri -128 ile 128 arasında olduğu için int8 değerinin en küçük değerini çıktı olarak alırız.
		// Bu sebeple değerleri birbirine dönüştürürken küçük olan değeri büyük olana benzetmemiz hata payını oratadan kaldırır.
	*/
	/*
		x := 10
		y := "10"

		fmt.Printf("%v %T \n", x, x)
		fmt.Printf("%v %T \n", y, y)

		fmt.Println(x + y)  // invalid operation: x + y (mismatched types int and string) */
	/*
		x := 10
		y := "10"

		fmt.Printf("%v %T \n", x, x)
		fmt.Printf("%v %T \n", y, y)

		fmt.Println(x + int(y)) //cannot convert y (variable of type string) to type int
		// type corversion yöntemi ile bir string i bir integera dönüştüremeyiz. */

	num1 := 106
	str1 := string(106)

	fmt.Printf("%v %T \n", num1, num1)
	fmt.Println()
	fmt.Printf("%v %T \n", str1, str1) // https://theasciicode.com.ar/
	// Burada yaptığımız veri değişimi oluyor çünkü 106 sayısı yukarıdaki linktede var j harfine denk geliyor.

}
