// Варинат 1
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
	var randNum, userNum int
	var err error

	randNum = rand.Intn(10) + 1

	fmt.Println("Угадайте число от 1 до 10!")

	for {

		fmt.Print("Ваш вариант: ")

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {

			str = scanner.Text()

			userNum, err = strconv.Atoi(str)

			if err != nil {

				fmt.Println("Ошибка! Введено не число!")

			} else if userNum >= 1 && userNum <= 10 {

				if userNum == randNum {

					fmt.Println("Вы угадали! Это", randNum)
					break

				} else if userNum < randNum {

					fmt.Println("Загаданное число больше!")

				} else if userNum > randNum {

					fmt.Println("Загаданное число меньше!")
				}
			} else {
				fmt.Println("Ошибка! Введенное число не попадает в диапазон от 1 до 10!")
			}
		}

	}
}
