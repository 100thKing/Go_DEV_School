package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomArray(n int) []int {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	randomArray := make([]int, n)
	for i := 0; i < n; i++ {
		randomArray[i] = rand.Intn(100) // Генерация случайного числа от 0 до 99
	}
	return randomArray
}

func FindMax(arr []int, n int) int {
	//	if len(arr) == 0 {
	//		return 0, -1 // Возвращаем 0 и -1 в случае пустого массива
	//		}

	max := arr[0] // Предполагаем, что первый элемент - максимальный
	maxIndex := 0

	for i, num := range arr {
		if num > max {
			max = num // Обновляем максимальное значение, если находим большее число
			maxIndex = i
		}
	}

	return maxIndex
}

func main() {
	var n int
	var size int
	fmt.Print("Введите размер массива: ")
	fmt.Scanln(&size)
	arr := generateRandomArray(size)
	fmt.Println(arr)
	fmt.Print("Введите номер числа для вывода: ")
	fmt.Scanln(&n)
	//fmt.Println(n)
	for i := 0; i < n-1; i++ {
		maxIndex := FindMax(arr, n)
		arr[maxIndex] = 0
	}
	maxIndex := FindMax(arr, n)
	fmt.Println(arr[maxIndex])
}
