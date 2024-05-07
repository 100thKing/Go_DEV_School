package main

import (
    "fmt"
    "time"
	"math/rand"
)

func generateRandomArray(n int) []int {
    rand.Seed(time.Now().UnixNano())
    arr := make([]int, n)

    for i := 0; i < n; i++ {
        arr[i] = i
    }

    // Перемешиваем массив
    for i := range arr {
        j := rand.Intn(i + 1)
        arr[i], arr[j] = arr[j], arr[i]
    }

    return arr
}


func bubbleSort(arr []int) time.Duration {
    start := time.Now()
    n := len(arr)

    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }

    fmt.Printf("Bubblesort: %v\n", time.Since(start))
    return time.Since(start)
}

func main() {
    numbers := generateRandomArray(100000)
    duration := bubbleSort(numbers)
    //fmt.Println(numbers)
}
