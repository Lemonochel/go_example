// программа сообщает сдал ли пользователь экзамен
package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Print("Введите значение: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "сдал"
	} else {
		status = "не сдал"
	}
	fmt.Println("Ваше значение равно: ", grade, ". И ты ", status)

}
