package main

import (
	"fmt"
	"os"
)

func main() {
	var operator string
	var operand1, operand2 float64
	var err error

	fmt.Println("Программа калькулятор")
	fmt.Print("Введите оператор (+, -, /, *, %, **): ")

	if _, err = fmt.Scan(&operator); err != nil {
		fmt.Println("Недопустимый ввод")
		os.Exit(1)
	}

	if operator != "+" &&
		operator != "-" &&
		operator != "/" &&
		operator != "*" &&
		operator != "%" &&
		operator != "**" {

		fmt.Println("Недопустимый математический оператор")
		os.Exit(1)
	}

	if operator == "**" {
		fmt.Println("Введите операнд для отбезания десятичной части")
		if _, err = fmt.Scanf("%f", &operand1); err != nil {
			fmt.Println("Недопустимый формат операнда")
			os.Exit(1)
		}

		fmt.Printf("Результат возведения в степень числа %f равен %f\n", operand1, operand1*operand1)

	} else if operator == "%" {
		fmt.Println("Введите операнд для позведения его в степень")
		if _, err = fmt.Scanf("%f", &operand1); err != nil {
			fmt.Println("Недопустимый формат операнда")
			os.Exit(1)
		}

		fmt.Printf("Результат обрезания от числа %f десятичной части равен %d\n", operand1, int(operand1))
	} else {
		fmt.Println(`Введите два операнда через пробел в формате "%f %f":`)
		if _, err = fmt.Scanf("%f %f", &operand1, &operand2); err != nil {
			fmt.Println("Недопустимый формат операндов")
			os.Exit(1)
		}

		switch operator {
		case "+":
			fmt.Printf("Сумма чисел %f и %f равна %f\n", operand1, operand2, operand1+operand2)
		case "-":
			fmt.Printf("Разность чисел %f и %f равна %f\n", operand1, operand2, operand1-operand2)
		case "*":
			fmt.Printf("Произведение чисел %f и %f равна %f\n", operand1, operand2, operand1*operand2)
		case "/":
			fmt.Printf("Отношение чисел %f и %f равна %f\n", operand1, operand2, operand1/operand2)
		}
	}

	return
}
