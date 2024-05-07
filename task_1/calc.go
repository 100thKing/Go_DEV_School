package main

import (
	"fmt"
	"os"
	"strconv"
)

func calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("деление на ноль невозможно")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("некорректная операция. Пожалуйста, используйте символы +, -, * или /")
	}
}

func isNumber(input string) float64 {
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Printf("Ошибка: %s не является числом\n", input)
		os.Exit(1)
	}
	return num
}

func main() {
	var inp1, inp2, operator string
	var num1, num2 float64

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&inp1)
	num1 = isNumber(inp1)

	fmt.Print("Выберите операцию (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&inp2)
	num2 = isNumber(inp2)

	result, err := calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}

	fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
