// Варинат 4 c ООП
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Game struct {
	randNum  int
	attempts int
	win      bool
}

func NewGame() Game {
	return Game{
		randNum:  rand.Intn(10) + 1,
		attempts: 0,
		win:      false,
	}
}

func (g Game) Play() {
	var str string
	var userNum int
	var err error

	fmt.Println("Угадайте число от 1 до 10!")

	for {
		fmt.Print("Введите количество попыток: ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			str = scanner.Text()
			g.attempts, err = strconv.Atoi(str)
			if err != nil {
				fmt.Println("Ошибка! Введено не число!")
			} else {
				break
			}
		}
	}
	for i := 0; i < g.attempts; i++ {
		fmt.Print("Ваш вариант: ")

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			str = scanner.Text()
			userNum, err = strconv.Atoi(str)

			if err != nil {
				fmt.Println("Ошибка! Введено не число!")

			} else if userNum >= 1 && userNum <= 10 {

				if userNum == g.randNum {
					g.win = true
					fmt.Println("Вы угадали! Это", g.randNum)
					break
				} else if userNum < g.randNum {
					fmt.Println("Загаданное число больше!")
				} else if userNum > g.randNum {
					fmt.Println("Загаданное число меньше!")
				}
			} else {
				fmt.Println("Ошибка! Введенное число не попадает в диапазон от 1 до 10!")
			}

		}
	}
}

func main() {

	game := NewGame()
	game.Play()

}
