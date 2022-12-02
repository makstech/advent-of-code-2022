package main

import (
	"advent-of-code-2022/day2"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Bₒₒₜᵢₙg ₐₗₗ ₜₕₑ ᵤₙₙₑcₑₛₛₐᵣy cₒₛₘᵢc ᵣₐdᵢₐₜᵢₒₙ fₒᵣ ₜₕᵢₛ ₛₕᵢₜcₒdₑ ₜₒ wₒᵣₖ")
	fmt.Print("ₚᵢcₖ ₐ dₐy ... ")
	var day string
	fmt.Scanln(&day)

	funcs := map[string]func(){
		"1": Day1,
		"2": day2.Day2,
	}

	if fun, ok := funcs[day]; ok {
		fmt.Println("ₛₜₐᵣₜᵢₙg dₐy ...", day)
		fun()
	} else {
		log.Fatal("day doesn't exist")
	}
}
