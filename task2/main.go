package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	var err error

	fmt.Println("Программа ведёт поиск всех простых чисел от 0 до N")
	fmt.Print("Введите число N: ")

	if _, err = fmt.Scanf("%d", &n); err != nil {
		fmt.Println("Ошибка ввода числа N")
		os.Exit(1)
	}

	if n < 0 {
		fmt.Println("Число N не должно быть меньше нуля")
		os.Exit(1)
	}

	for i := 0; i <= n; i++ {
		if isPrimeNumber(i) {
			fmt.Printf("Найдено новое простое число: %d\n", i)
		}
	}

	return
}

// Функция возвращает true, если число в аргументе n является простым
func isPrimeNumber(n int) bool {
	var result bool = true

	for j := 2; j < n; j++ {
		if n%j == 0 {
			result = false
			break
		}
	}

	return result
}
