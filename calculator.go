package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Input() (string, string, string) {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	initial_input := input.Text()
	arr := strings.Split(initial_input, " ")
	if len(arr) != 3 {
		fmt.Println("Ошибка: математическая операция должна состоять из двух элементов")
		os.Exit(1)
	}
	x, sign, y := arr[0], arr[1], arr[2]
	x1, y1 := translationInt(x, y)
	x2 := InputRoma(x)
	y2 := InputRoma(y)
	if x1 > 10 || y1 > 10 || x1 < 0 || y1 < 0 {
		fmt.Println("Ошибка: числа находятся в диапозоне от 1 до 10")
		os.Exit(1)
	}
	if x2 > 10 || y2 > 10 || x2 < 0 || y2 < 0 {
		fmt.Println("Ошибка: числа находятся в диапозоне от 1 до 10")
		os.Exit(1)
	}
	return sign, x, y
}

func InputRoma(r string) int {
	var tmp, tmpa, tmps int
	roma := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := len(r) - 1; i >= 0; i-- {
		tmps = roma[r[i]]
		if tmps < tmpa {
			tmp -= tmps
		} else {
			tmp += tmps
		}
		tmpa = tmps
	}
	return tmp
}

func OutputRoma(number int) string {
	conver := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	roma := ""
	for _, converP := range conver {
		for number >= converP.value {
			roma += converP.digit
			number -= converP.value
		}
	}
	return roma
}

func translationInt(x, y string) (int, int) {
	var number1, number2 int
	number1, _ = strconv.Atoi(x)
	number2, _ = strconv.Atoi(y)
	return number1, number2
}

func Calculator(sign, x, y string) {
	arab := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	roma := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	isRoman := false
	for i := 0; i < len(arab); i++ {
		for j := 0; j < len(arab); j++ {
			if (x == roma[i] && y == arab[j]) || (x == arab[i] && y == roma[j]) {
				fmt.Println("Ошибка: разные типы счисления")
				return
			}
			if x == arab[i] && y == arab[j] {
				number1, number2 := translationInt(x, y)
				performOperation(sign, number1, number2)
				return
			}
		}
		if x == roma[i] || y == roma[i] {
			isRoman = true
		}
	}

	if isRoman {
		number1 := InputRoma(x)
		number2 := InputRoma(y)
		performRomanOperation(sign, number1, number2)
	} else {
		fmt.Println("Ошибка: неизвестные значения")
	}
}

func performOperation(sign string, x, y int) {
	switch sign {
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	case "*":
		fmt.Println(x * y)
	case "/":
		fmt.Println(x / y)
	default:
		fmt.Println("Ошибка: неизвестная математическая операция")
	}
}

func performRomanOperation(sign string, x, y int) {
	switch sign {
	case "+":
		fmt.Println(OutputRoma(x + y))
	case "-":
		if x-y < 1 {
			fmt.Println("Ошибка: в римской системе нет отрицательных чисел и нуля")
		} else {
			fmt.Println(OutputRoma(x - y))
		}
	case "*":
		fmt.Println(OutputRoma(x * y))
	case "/":
		fmt.Println(OutputRoma(x / y))
	default:
		fmt.Println("Ошибка: неизвестная математическая операция")
	}
}

func main() {
	var sign, x, y string
	sign, x, y = Input()
	Calculator(sign, x, y)
}
