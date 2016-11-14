package main

import (
	"errors"
	"fmt"
	"github.com/GDGPetropolis/DojoGo/src/happynumbers/happy"
	"os"
	"strconv"
)

func main() {
	var err error
	start, stop := 1, 11

	switch numArgs := len(os.Args); {
	case numArgs == 2:
		stop, err = strconv.Atoi(os.Args[1])
	case numArgs == 3:
		start, err = strconv.Atoi(os.Args[1])
		stop, err = strconv.Atoi(os.Args[2])
	case numArgs > 3:
		err = errors.New("Too many arguments!")
	}

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("usage: %s [[start] stop]")
		os.Exit(1)
	}

	for i := start; i < stop; i++ {
		if happy.IsHappy(i) {
			fmt.Println(":)")
		} else {
			fmt.Println(i)
		}
	}
}
