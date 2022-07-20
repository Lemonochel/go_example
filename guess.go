// эта игра угадай число
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Я загадал число от 1 до 100")
	fmt.Println("Угадай это число")
	//fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("У вас осталось ", 10-guesses, " попыток")
		fmt.Print("Напиши свое число: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Ваше число меньше загаданного")
		} else if guess > target {
			fmt.Println("Ваше число больше загаданного")
		} else {
			success = true
			fmt.Println("Отличная работа! Вы Угадали!")
			break
		}
	}
	if !success {
		fmt.Println("Вы проиграли, было загадано число: ", target)
	}
}
