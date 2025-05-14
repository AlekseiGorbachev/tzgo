// Варинат 2
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	var str string
	var randNum, userNum, attempts int
	var err error

	randNum = rand.Intn(10) + 1

	fmt.Println("Угадайте число от 1 до 10!")

   for {

	fmt.Print("Введите количество попыток: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
		attempts, err = strconv.Atoi(str)
		if err != nil {
			fmt.Println("Ошибка! Введено не целое число!")
		} else {
			break
		}
	 }
	
	}

	for i := 1; i <= attempts; i++ {

		fmt.Print("Ваш вариант: ")

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {

			str = scanner.Text()

			userNum, err = strconv.Atoi(str)

			if err != nil {

				fmt.Println("Ошибка! Введено не число!")

			} else if userNum >= 1 && userNum <= 10 {

				switch {

				case userNum == randNum:
					fmt.Println("Вы угадали! Это", randNum)
					i = attempts

				case userNum < randNum:
					fmt.Println("Загаданное число больше!")

				case userNum > randNum:
					fmt.Println("Загаданное число меньше!")
				}

			} else {
				fmt.Println("Ошибка! Введенное число не попадает в диапазон от 1 до 10!")
			}

		}
	
	}
}
