package main // Define Main Package

import "fmt" //Import fmt

func main() {
	fmt.Println("Hello World") //PrintLn
	//fmt.Println("Hello World")
	fmt.Println("Type of values")

	fmt.Println("String")
	fmt.Println(24)   // int
	fmt.Println(2.5)  //float 32 or 64 bits
	fmt.Println(true) //boolean
	// if we declare a variable we  must use it yes or yes
	var name string //set type

	name = "Grego" // add value

	fmt.Println(name) // use them

	name = "Pablo" //Change the value

	fmt.Println(name)

	var firtname, lastname string = "Grego", "Garcia" //we can set multiple var and set de value in the same line

	fmt.Println(firtname + " " + lastname) // as well as print 2 or more concat with "+" like Java

	var city, country string

	city = "Cordoba"
	country = "Arg"

	fmt.Println(city + " " + country)

	var a, b, c, d int = 1, 2, 3, 4
	// if we declare int vars , we can print them concat use ","
	fmt.Println(a, b, c, d)

	var (
		pi      float64
		boolean bool
		//if we declare a value of a variable, this will always have to have the same type of value
		spell = "Mamalona"
		// I can't do this : {spell = 23} because "spell" is s String variable now.
		ent = 24
	)

	pi = 3.141592
	boolean = true

	println(pi, boolean, spell, ent)

}
