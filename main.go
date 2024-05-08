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
		"I":        1,
		"II":       2,
		"III":      3,
		"IV":       4,
		"V":        5,
		"VI":       6,
		"VII":      7,
		"VIII":     8,
		"IX":       9,
		"X":        10,
		"XI":       11,
		"XII":      12,
		"XIII":     13,
		"XIV":      14,
		"XV":       15,
		"XVI":      16,
		"XVII":     17,
		"XVIII":    18,
		"XIX":      19,
		"XX":       20,
		"XXI":      21,
		"XXII":     22,
		"XXIII":    23,
		"XXIV":     24,
		"XXV":      25,
		"XXVI":     26,
		"XXVII":    27,
		"XXVIII":   28,
		"XXIX":     29,
		"XXX":      30,
		"XXXI":     31,
		"XXXII":    32,
		"XXXIII":   33,
		"XXXIV":    34,
		"XXXV":     35,
		"XXXVI":    36,
		"XXXVII":   37,
		"XXXVIII":  38,
		"XXXIX":    39,
		"XL":       40,
		"XLI":      41,
		"XLII":     42,
		"XLIII":    43,
		"XLIV":     44,
		"XLV":      45,
		"XLVI":     46,
		"XLVII":    47,
		"XLVIII":   48,
		"XLIX":     49,
		"L":        50,
		"LI":       51,
		"LII":      52,
		"LIII":     53,
		"LIV":      54,
		"LV":       55,
		"LVI":      56,
		"LVII":     57,
		"LVIII":    58,
		"LIX":      59,
		"LX":       60,
		"LXI":      61,
		"LXII":     62,
		"LXIII":    63,
		"LXIV":     64,
		"LXV":      65,
		"LXVI":     66,
		"LXVII":    67,
		"LXVIII":   68,
		"LXIX":     69,
		"LXX":      70,
		"LXXI":     71,
		"LXXII":    72,
		"LXXIII":   73,
		"LXXIV":    74,
		"LXXV":     75,
		"LXXVI":    76,
		"LXXVII":   77,
		"LXXVIII":  78,
		"LXXIX":    79,
		"LXXX":     80,
		"LXXXI":    81,
		"LXXXII":   82,
		"LXXXIII":  83,
		"LXXXIV":   84,
		"LXXXV":    85,
		"LXXXVI":   86,
		"LXXXVII":  87,
		"LXXXVIII": 88,
		"LXXXIX":   89,
		"XC":       90,
		"XCI":      91,
		"XCII":     92,
		"XCIII":    93,
		"XCIV":     94,
		"XCV":      95,
		"XCVI":     96,
		"XCVII":    97,
		"XCVIII":   98,
		"XCIX":     99,
		"C":        100,
	}
	return romanMap[romanNum]
}

func arabicToRoman(arabicNum int) string {
	romanMap := map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		22:  "XXII",
		23:  "XXIII",
		24:  "XXIV",
		25:  "XXV",
		26:  "XXVI",
		27:  "XXVII",
		28:  "XXVIII",
		29:  "XXIX",
		30:  "XXX",
		31:  "XXXI",
		32:  "XXXII",
		33:  "XXXIII",
		34:  "XXXIV",
		35:  "XXXV",
		36:  "XXXVI",
		37:  "XXXVII",
		38:  "XXXVIII",
		39:  "XXXIX",
		40:  "XL",
		41:  "XLI",
		42:  "XLII",
		43:  "XLIII",
		44:  "XLIV",
		45:  "XLV",
		46:  "XLVI",
		47:  "XLVII",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		51:  "LI",
		52:  "LII",
		53:  "LIII",
		54:  "LIV",
		55:  "LV",
		56:  "LVI",
		57:  "LVII",
		58:  "LVIII",
		59:  "LIX",
		60:  "LX",
		61:  "LXI",
		62:  "LXII",
		63:  "LXIII",
		64:  "LXIV",
		65:  "LXV",
		66:  "LXVI",
		67:  "LXVII",
		68:  "LXVIII",
		69:  "LXIX",
		70:  "LXX",
		71:  "LXXI",
		72:  "LXXII",
		73:  "LXXIII",
		74:  "LXXIV",
		75:  "LXXV",
		76:  "LXXVI",
		77:  "LXXVII",
		78:  "LXXVIII",
		79:  "LXXIX",
		80:  "LXXX",
		81:  "LXXXI",
		82:  "LXXXII",
		83:  "LXXXIII",
		84:  "LXXXIV",
		85:  "LXXXV",
		86:  "LXXXVI",
		87:  "LXXXVII",
		88:  "LXXXVIII",
		89:  "LXXXIX",
		90:  "XC",
		91:  "XCI",
		92:  "XCII",
		93:  "XCIII",
		94:  "XCIV",
		95:  "XCV",
		96:  "XCVI",
		97:  "XCVII",
		98:  "XCVIII",
		99:  "XCIX",
		100: "C",
	}
	return romanMap[arabicNum]
}

func calculate(operand1 int, operator string, operand2 int) int {
	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		result := operand1 - operand2
		if result <= 0 {
			panic("Отрицательный или нулевой результат для римских чисел")
		}
		return result
	case "*":
		return operand1 * operand2
	case "/":
		if operand1%operand2 != 0 {
			panic("Деление римских чисел должно быть без остатка")
		}
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

	var isRomanInput bool
	var isArabicInput bool

	operand1, err := strconv.Atoi(operand1Str)
	if err != nil {
		isRomanInput = true
		operand1 = romanToArabic(operand1Str)
	} else {
		isArabicInput = true
	}

	operand2, err := strconv.Atoi(operand2Str)
	if err != nil {
		if isRomanInput {
			operand2 = romanToArabic(operand2Str)
		} else {
			panic("Одновременный ввод римских и арабских чисел недопустим")
		}
	} else {
		if isArabicInput {
		} else {
			panic("Одновременный ввод римских и арабских чисел недопустим")
		}
	}

	if operand1 < 1 || operand1 > 10 || operand2 < 1 || operand2 > 10 {
		panic("Калькулятор принимает только числа от 1 до 10 включительно")
	}

	result := calculate(operand1, operator, operand2)

	if isRomanInput {
		if result <= 0 {
			panic("Отрицательный или нулевой результат для римских чисел")
		}

		fmt.Printf("Результат: %s\n", arabicToRoman(result))
	} else {
		fmt.Printf("Результат: %d\n", result)
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
