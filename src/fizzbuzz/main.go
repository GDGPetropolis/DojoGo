package main

import (
	"fmt"
	"github.com/GDGPetropolis/DojoGo/src/fizzbuzz/fizzbuzz"
)

func main() {
	fmt.Println("-------------------------------");
	fmt.Println("Bem vindo ao GDG Petrópolis Day");
	fmt.Println("-------------------------------");

	var value = 10;

	for i := 1; i <= value; i++ {
		fmt.Println(fizzbuzz.FizzBuzz(i));
	}
}
