// main.go
package main

import (
	"flag"
	"fmt"
	"mymodule/calculations"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

func main() {
	var n int
	var logFlag bool
	var result int
	args := os.Args

	if len(args) < 2 {
		fmt.Print("Введите число: ")
		fmt.Scanf("%d\n", &n)
		fmt.Print("Использовать логирование?(введите 1/0): ")
		fmt.Scanf("%t\n", &logFlag)
		result := calculations.Calculate(n, logFlag)
		fmt.Printf("Факториал %d = %d\n", n, result)
		os.Exit(1)
	}

	x, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Ошибка при преобразовании числа. Пожалуйста, введите корректное число.")
		return
	}

	fmt.Printf("Число 'n' получено из аргументов командной строки: %d\n", x)
	// Задаем флаг '-log' с типом bool и значением по умолчанию false
	logFlag_bool := flag.Bool("log", false, "Включить логирование")

	flag.Parse()

	fmt.Printf("Значение флага '-log': %t\n", *logFlag_bool)

	// Если установлен флаг '-log', выводим сообщение
	if *logFlag_bool {
		fmt.Println("Логирование включено.")
	} else {
		fmt.Println("Логирование выключено.")
	}

	// Отображаем справку по флагу '-log'
	flag.PrintDefaults()

	// Проверяем, были ли указаны флаги
	if flag.NFlag() == 0 {
		fmt.Println("Не были указаны флаги.")
		os.Exit(1)
	}

	result = calculations.Calculate(x, *logFlag_bool)

	// Создаем новый логгер
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.Out = os.Stdout
	logger.Infof("Факториал %d = %d\n", x, result)
	logger.Error(err)
}
