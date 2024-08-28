package main

import "fmt"

func main() {

	//variaveis
	var name string = "Belmondo"
	var surname = "Jr"
	var version float32 = 1.1

	fmt.Println("Hello, ", name, surname)
	fmt.Println("My first program in Go Lang")
	fmt.Print("I'm in my version: ", version)

	//another way to declare variables
	anothername := "Belmondo"
	anothersurname := "Jr"
	anotherversion := 1.1

	fmt.Println("Hello, ", anothername, anothersurname)
	fmt.Println("My first program in Go Lang")
	fmt.Print("I'm in my version: ", anotherversion)

}
