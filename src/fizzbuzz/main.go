package main

import (
	"fmt"
	"./dojo"
)

func main() {

	// var values = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25}

	fmt.Println("-------------------------------");
	fmt.Println("Bem vindo ao GDG Petrópolis Day");
	fmt.Println("-------------------------------");

	var value = 7;

	fmt.Println(dojo.FizzBuzz(value));
}