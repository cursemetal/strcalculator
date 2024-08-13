package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rdr := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение: ")
	rust, _ := rdr.ReadString('\n')
	rust = strings.TrimSpace(rust)

	x := strings.Split(rust, " ")
	if len(x) < 3 {
		fmt.Println("Некорректный ввод.") //Проверка ввода что бы было 3 части
		return
	}

	var frop, num, sdop string
	if len(x) > 3 {
		frop = strings.Trim(x[0], "\"") + " " + strings.Trim(x[1], "\"")
		num = x[2]
		sdop = strings.Trim(x[3], "\"")
	} else {
		frop = strings.Trim(x[0], "\"")
		num = x[1]
		sdop = strings.Trim(x[2], "\"")
	}

	var surgery string // операторы
	var n int
	var err error

	switch num {
	case "+":
		surgery = frop + sdop
	case "-":
		surgery = strings.Replace(frop, sdop, "", 1)
	case "*":

		n, err = strconv.Atoi(sdop)
		if err != nil {
			fmt.Println("Некорректый ввод.") // конверт 2 число
			return
		}
		surgery = strings.Repeat(frop, n)
	case "/":

		n, err = strconv.Atoi(sdop)
		if err != nil {
			fmt.Println("Некорректный ввод .") // конверт 2 число
			return
		}

		if len(frop)%n != 0 {
			fmt.Println("Некорректный ввод.") //деление строки на число, что бы можно было поделить колво символов без остатка
			return
		}
		surgery = frop[:len(frop)/n]
	default:
		fmt.Println("Некорректный ввод")
		return
	}

	if len(surgery) > 40 {
		surgery = surgery[:40] + "..." // 40 + символов
	}

	fmt.Println("Результат:", surgery)
}
