package main

import (
	"fmt"
	"math"
)

func main() {
	// var height = 1.75
	// var weight float64 = 79

	// height := 1.75
	// weight := 79.0
	// var imt = weight / math.Pow(height, 2)

	fmt.Println("Калькулятор индекса массы тела")

	const POWER = 2

	var height float64
	var weight float64

	fmt.Println("Введите ваш рост:")
	_, err := fmt.Scan(&height)
	if err != nil {
		fmt.Println("Ошибка ввода роста:", err)
		return
	}

	height = height / 100

	fmt.Println("Введите ваш вес:")
	_, err = fmt.Scan(&weight)
	if err != nil {
		fmt.Println("Ошибка ввода веса:", err)
		return
	}

	imt := weight / math.Pow(height, POWER)

	fmt.Printf("Ваш индекс массы тела: %.2f\n", imt)
}
