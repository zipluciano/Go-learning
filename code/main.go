package main

import "fmt"

func main() {
	// to run the code -> go run <file>
	fmt.Println("Hello, World!")

	// defining variables
	// three ways to write to declare a variable
	var name_one string = "Hello,"
	var name_two = "World!"
	// can't declare a variable and dont use, but can declare and use without value
	var name_three string
	fmt.Println(name_one, name_two, name_three)

	// now using
	name_three = "World!"
	fmt.Println(name_one, name_three)

	// declaring this way, only inside a function
	my_fname := "Luciano"
	fmt.Println(my_fname)

	var num_1 int = 42
	// consedering IEEE 754
	var num_2 int16 = -42
	num_3 := 424
	num_4 := 42
	var num_5 uint = 255
	fmt.Println(num_1, num_2, num_3, num_4, num_5)

	var num_6 float64 = 42.42
	num_7 := 42.42
	fmt.Println(num_6, num_7)

	// difference between Print and Prinln
	fmt.Print("Enough hello World! ")
	fmt.Print("Now, hello Universe!")

	fmt.Print("Enough hello World! \n")
	fmt.Print("Now, hello Universe!")

	fmt.Println("Enough hello World!")
	fmt.Println("Now, hello Universe!")

	age := 24
	fmt.Printf("my age is %d, and my first name is %v", age, my_fname)

	// float formating
	fmt.Printf("\nPoints: %f\n", num_6)
	fmt.Printf("Points: %0.1f\n", num_6)
	fmt.Printf("Points: %0.2f\n", num_6)
}
