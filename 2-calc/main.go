package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Определяем тип для функций операций
type OperationFunc func([]float64) float64

// Создаем map для хранения функций операций
var operations = map[string]OperationFunc{
	"SUM": sum,
	"AVG": avg,
	"MED": median,
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите операцию (AVG, SUM, MED): ")
	opInput, _ := reader.ReadString('\n')
	operation := strings.ToUpper(strings.TrimSpace(opInput))

	// Проверяем, существует ли операция в нашей map
	opFunc, exists := operations[operation]
	if !exists {
		fmt.Println("Неизвестная операция. Используйте AVG, SUM или MED.")
		return
	}

	fmt.Print("Введите числа через запятую: ")
	numsInput, _ := reader.ReadString('\n')
	numStrs := strings.Split(strings.TrimSpace(numsInput), ",")
	var nums []float64

	for _, s := range numStrs {
		n, err := strconv.ParseFloat(strings.TrimSpace(s), 64)

		if err != nil {
			fmt.Printf("Некорректное число: %s\n", s)
			return
		}
		nums = append(nums, n)
	}

	// Вызываем функцию из map
	result := opFunc(nums)
	fmt.Printf("%s: %.2f\n", operation, result)
}

func sum(nums []float64) float64 {
    if len(nums) < 2 {
        fmt.Println("Введите минимум два числа")
        return 0
    }
    total := 0.0
    for _, n := range nums {
        total += n
    }
    return total
}

func avg(nums []float64) float64 {
    if len(nums) < 2 {
        fmt.Println("Введите минимум два числа")
        return 0
    }
    return sum(nums) / float64(len(nums))
}

func median(nums []float64) float64 {
    if len(nums) < 2 {
        fmt.Println("Введите минимум два числа")
        return 0
    }
    sort.Float64s(nums)
    mid := len(nums) / 2
    if len(nums)%2 == 0 {
        return (nums[mid-1] + nums[mid]) / 2
    }
    return nums[mid]
}
