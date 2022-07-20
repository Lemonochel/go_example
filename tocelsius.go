package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Println("Введите температуру по Фаренгейту: ")
	fahr, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	cels := (fahr - 32) * 5 / 9
	fmt.Printf("%0.2f в Цельсиях\n", cels)
}
