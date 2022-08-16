package main

import (
	"fmt"
	"log"
)

func main() {
	ak47, err := GetGun("ak47")
	if err != nil {
		log.Println("Error occured", err.Error())
	}
	PrintDetails(ak47)
	musket, err := GetGun("musket")
	if err != nil {
		log.Println("Error occured", err.Error())
	}
	PrintDetails(musket)

}

func PrintDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
