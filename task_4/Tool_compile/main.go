package main

func main() {
	rows := 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			print("*")
		}
		println()
	}

	for i := rows - 1; i >= 1; i-- {
		for j := 1; j <= rows-i; j++ {
			print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			print("*")
		}
		println()
	}
}
