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
	rust, _ := rdr.ReadString('\n') //считываем ввод от пользователя до новой строки
	rust = strings.TrimSpace(rust)  //

	x := strings.Split(rust, " ") // разжделение строки пробелами

	
	var strng string
	if len(x[0]) >= 2 && x[0][0] == '"' && x[0][len(x[0])-1] == '"' {
		// Если есть кавычки, удаляем их
		strng = x[0][1 : len(x[0])-1]
	} else {
		// Если нет кавычек, просто используем первый элемент
		strng = x[0]
	}

	if len(strng) > 10 {
		panic("Строка не должна быть длиннее 10 символов") 
	}

	var operateons string
	var num int
	switch len(x) { // Проверка на верноесть ввода что 1 это 1
	case 3:
		operateons = x[1]
		num, _ = strconv.Atoi(x[2]) // конверт
	case 4:
		operateons = x[2]
		num, _ = strconv.Atoi(x[3]) // конверт
	default:
		panic("Некорректный формат выражения")
	}

	if operateons == "*" || operateons == "/" { // (работает для *, /)
		if num < 1 || num > 10 {
			panic("Число должно быть от 1 до 10")
		}
	}

	var surgery string
	switch operateons {
	case "+":
		
		surgery = strng + x[len(x)-1][1:len(x[len(x)-1])-1]
	case "-":
		
		surgery = strings.ReplaceAll(strng, x[len(x)-1][1:len(x[len(x)-1])-1], "") // Действия
	case "*":
		surgery = strings.Repeat(strng, num)
	case "/":
		if len(strng)%num != 0 {
			panic("Невозможно разделить строку на равные части")
		}
		lgLength := len(strng) / num
		surgery = strng[0:lgLength] // только часть строки кратную числу
	default:
		panic("Неподдерживаемая операция")
	}

	
	surgery = strings.Trim(surgery, "\"")

	if len(surgery) > 40 {
		fmt.Println(surgery[0:40] + "...") 
	} else {
		fmt.Println(surgery)
	}
}
