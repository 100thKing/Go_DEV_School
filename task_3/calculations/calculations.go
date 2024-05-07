package calculations

import "log"

func Calculate(n int, logFlag bool) int {
	if logFlag {
		log.Println("Начаты вычисления...")
		log.Printf("Вычисляется %d!", n)
	}

	result := int(1)
	for i := int(1); i <= n; i++ {
		result *= i
	}

	if logFlag {
		log.Println("Вычисления завершены!")
	}

	return result
}
