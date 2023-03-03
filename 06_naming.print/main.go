/*
	İNTEGER

%b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%O	base 8 with 0o prefix
%q	a single-quoted character literal safely escaped with Go syntax.
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04X"

BOOLEAN
%t	the word true or false

Floating-point and complex constituents:

%b	decimalless scientific notation with exponent a power of two,

	in the manner of strconv.FormatFloat with the 'b' format,
	e.g. -123456p-78

%e	scientific notation, e.g. -1.234456e+78
%E	scientific notation, e.g. -1.234456E+78
%f	decimal point but no exponent, e.g. 123.456
%F	synonym for %f
%g	%e for large exponents, %f otherwise. Precision is discussed below.
%G	%E for large exponents, %F otherwise
%x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
%X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

The default format for %v is:

bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
*/
package main

import "fmt"

func main() {
	/*
		// ( Print  -  Println )  - Printf

		fmt.Print("Hello")
		fmt.Println("Helloo") // olduğu gibi değerler yan yana yazılır.   çıktısı := HelloHelloo

		// fakat Println çıktısını üst satıra yazarsak println sayesinde satır atlayıp ayrı ayrı yazılır.

		fmt.Println("Hello")
		fmt.Print("Helloo")
		fmt.Println()       // yaparsak satır atlatmış oluruz.
		fmt.Printf("Hello") // aslında bu üçüde aynıdır fakat Println den sonra diğer satıra geçer.
	*/

	// name := "Semih"
	/*
		fmt.Print(name)
		fmt.Println(name)
		fmt.Printf(name)
	*/

	/* fmt.Print("Benim adim ", name)
	fmt.Println()
	fmt.Println("Benim adim ", name) */
	//fmt.Printf("Benim adim ", name) // printfte ise formatlandırmamız gerekiyor.Onuda %v ile name değerini çıkartabiliriz eğer %T yaparsak o değişkenin ismini yazdırır.

	/* Benim adim Semih        print ile println arasındaki farklardan bir diğeri ise değişkenden sonra bir boşluk koyar.
	   Benim adim  Semih */
	/* fmt.Printf("Benim adim %v", name)
	fmt.Println()
	fmt.Printf("Benim adim %T ", name)
	fmt.Println()
	fmt.Printf("Benim adim %X", name)

	*/
	/*
		x := 100
		y := 25
		z := 30

		fmt.Printf("%b %d %o", x, y, z) */

	//name, age := "Semih Nebi", 5

	// fmt.Print("My name is ", name, " I am ", age, " years old")
	// fmt.Println(" My name is ", name, " I am ", age, "years old")
	// fmt.Printf("My name is %v , I am %v yaşindayim  ", name, age)

	//  VISIBILITY = GÖRÜNÜRLÜLÜK

	x := 10
	fmt.Println(x)
	/* eğer x değişkenini biz packet ekleme yerine eklersek x:=10 diye değilde

		var x = 10 şeklinde eklememiz gerekir.
		Mesela
	package main

	import "fmt"

	x := 10 olmaz

	paket ekleme yerine eklersek var x = 10 diye eklemeliyiz.

	func main() {

	 x:= 10 olur.

	 }

	*/

	// Değişkenlerde isimlendirme yaparken kolay ve anlaşılır olması için kurallara dikkat edilmelidir.
	// Yani var co string yerine var coin string denebilir.
	// Başkaları tarafından ve kendi tarafımızdan daha kolay anlaşılması için bu kurallara dikkatt edilir.

	// Kısaltmalar büyük harflerle yazılır örneğin var URL doğru iken var Url yanlıştır. var HTTP gibi ...

}
