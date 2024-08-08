package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	rdr := bufio.NewReader(os.Stdin) // ридер для строки
	fmt.Print("Введите выражение: ")
	rust, _ := rdr.ReadString('\n') // считываем ввод от пользователя до новой строки
	rust = strings.TrimSpace(rust)  // игнорируем пробелы в начале и в конце

	x := strings.Split(rust, " ") // разжделение строки пробелами

	if !strings.Contains(x[0], "\"") {
		panic("Первый операнд должна быть строка") // функция для проверки 1го операнда

	}

	var result string
	switch len(x) {
	case 1:
		result = x[1]
	case 2:
		result = x[2]
	default:
		panic("Неверный формат выражение")
	}

	var num int
	switch len(x) { //достаем число
	case 1:
		num, _ = strconv.Atoi(x[2])
	case 2:
		num, _ = strconv.Atoi(x[3])
	}

	if num < 1 || num > 10 { // указываем диапазон от 1 до 10
		panic("Допустимый диапазон для ввода от 1 до 10")
	}

	strngs := x[0][1 : len(x[0])-1]
	if len(strngs) > 10 {
		panic("Допустимая длинна строки 10 символов")
	}

	var surgery string
	switch result {
	case "+":
		surgery = strngs + x[len(x)-1][1:len(x[len(x)-1])-1]
	case "-":
		surgery = strings.ReplaceAll(strngs, x[len(x)-1][1:len(x[len(x)-1])-1], "")
	case "*":
		surgery = strings.Repeat(strngs, num)
	case "/":
		if len(strngs)%num != 0 {
			panic("Невозможно разделить строку на равыне части")
		}
		surgery = strngs[0 : len(strngs)/num]
	default:
		panic("Неверная операция")

	}

	if len(surgery) > 40 {
		fmt.Println(surgery[0:40] + "...")
	} else {
		fmt.Println(surgery)
	}

}
