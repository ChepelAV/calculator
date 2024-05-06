package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToArabic(romanNum string) int {
	romanMap := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	return romanMap[romanNum]
}

func arabicToRoman(arabicNum int) string {
	romanMap := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}
	return romanMap[arabicNum]
}

func calculate(operand1 int, operator string, operand2 int) int {
	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		return operand1 / operand2
	default:
		panic("Неправильная арифметическая операция")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите выражение (например, 1 + 1):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("Введено неверное выражение")
	}

	operand1Str := parts[0]
	operator := parts[1]
	operand2Str := parts[2]

	operand1, err := strconv.Atoi(operand1Str)
	if err != nil {
		operand1 = romanToArabic(operand1Str)
	}

	operand2, err := strconv.Atoi(operand2Str)
	if err != nil {
		operand2 = romanToArabic(operand2Str)
	}

	if operand1 < 1 || operand1 > 10 || operand2 < 1 || operand2 > 10 {
		panic("Калькулятор принимает только числа от 1 до 10 включительно")
	}

	result := calculate(operand1, operator, operand2)

	if operand1Str != "" {
		fmt.Printf("Результат: %d\n", result)
	} else {
		fmt.Printf("Результат: %s\n", arabicToRoman(result))
	}
}

//func main() {
//	fmt.Println("Hello World")
//	var chislo1 int64
//	var chislo2 int64
//	var pro string
//	fmt.Println("Введите число1")
//fmt.Scan(&chislo1)
//	fmt.Println("Введите число2")
//	fmt.Scan(&chislo2)
//	fmt.Println("операция?")
//	fmt.Scan(&pro)
//	if pro == "+" {
//		sum := chislo1 + chislo2
//		fmt.Println(sum)
//	} else if pro == "-" {
//		sum := chislo1 - chislo2
//		fmt.Println(sum)
//	} else if pro == "*" {
//		sum := chislo1 * chislo2
//		fmt.Println(sum)
//	} else if pro == "/" {
//		sum := chislo1 / chislo2
//		fmt.Println(sum)
//	}
//	fmt.Println("ITS ALL")

//}
