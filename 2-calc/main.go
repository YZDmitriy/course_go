package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите операцию (AVG, SUM, MED): ")
	opInput, _ := reader.ReadString('\n')
	operation := strings.ToUpper(strings.TrimSpace(opInput))

	if(operation != "AVG" && operation != "SUM" && operation != "MED") {
		fmt.Println("Неизвестная операция. Используйте AVG, SUM или MED.")
		return
	}

	fmt.Print("Введите числа через запятую: ")
	numsInput, _ := reader.ReadString('\n')
	numStrs := strings.Split(numsInput, ",")
	var nums []float64

	for _, s := range numStrs {
		n, err := strconv.ParseFloat(strings.TrimSpace(s), 64)

		if err != nil {
			fmt.Printf("Некорректное число: %s\n", s)
			return
		}
		nums = append(nums, n)
	}

	switch operation {
	case "AVG":
		fmt.Printf("AVG: %.2f\n", avg(nums))
	case "SUM":
		fmt.Printf("SUM: %.2f\n", sum(nums))
	case "MED":
		fmt.Printf("MED: %.2f\n", median(nums))
	default:
		fmt.Println("Неизвестная операция. Используйте AVG, SUM или MED.")
	}
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
