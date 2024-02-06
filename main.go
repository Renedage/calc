package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romeNumbers = map[string]int{
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
	"XX":   20,
	"XXX":  30,
	"XL":   40,
	"L":    50,
	"LX":   60,
	"LXX":  70,
	"LXXX": 80,
	"XC":   90,
	"C":    100,
}

var romeNumbersReverse = map[int]string{
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
	20:  "XX",
	30:  "XXX",
	40:  "XL",
	50:  "L",
	60:  "LX",
	70:  "LXX",
	80:  "LXXX",
	90:  "XC",
	100: "C",
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите математическое выражение (В виде: 1 + 1 или I + II)")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		res1 := strings.Trim(text, "\r\n")

		fmt.Println(calculate(res1))

	}

}

func calculate(text string) string {
	s := strings.Split(text, " ")

	if len(s) != 3 {
		panic("Неверное количество параметров")
	}

	arab := 0
	rome := 0

	isArabic := false
	firstNumber := 0
	secondNumber := 0
	isArabic, firstNumber = getArabicNumber(s[0])

	if isArabic {
		arab += 1
	} else {
		rome += 1
	}

	isArabic, secondNumber = getArabicNumber(s[2])

	if isArabic {
		arab += 1
	} else {
		rome += 1
	}

	if firstNumber < 0 || secondNumber < 0 {
		panic("Строка не является математической операцией")
	} else if firstNumber > 10 || secondNumber > 10 {
		panic("Нельзя использовать числа больше 10")
	} else if arab == 2 {
		isArabic = true
	} else if rome == 2 {
		isArabic = false
	}

	if s[1] != "+" && s[1] != "-" && s[1] != "*" && s[1] != "/" {
		panic("Строка не является математической операцией")
	}

	if arab != 2 && rome != 2 {
		panic("Используются одновременно разные системы счисления.")
	}

	if s[1] == "+" {
		return formatNumber(firstNumber+secondNumber, isArabic)
	} else if s[1] == "-" {
		return formatNumber(firstNumber-secondNumber, isArabic)
	} else if s[1] == "*" {
		return formatNumber(firstNumber*secondNumber, isArabic)
	} else if s[1] == "/" {
		return formatNumber(firstNumber/secondNumber, isArabic)
	}
	return ""
}

func formatNumber(number int, isArabic bool) string {
	if isArabic {
		return strconv.Itoa(number)
	} else {
		i, ok := romeNumbersReverse[number]
		if ok {
			return i
		} else {
			firstNumber := number % 10
			secondNumber := number - firstNumber
			return romeNumbersReverse[secondNumber] + romeNumbersReverse[firstNumber]
		}
	}
}

func getArabicNumber(s string) (bool, int) {
	if number, err := strconv.Atoi(s); err == nil {
		return true, number
	} else {
		i, ok := romeNumbers[s]
		if !ok {
			return false, -1
		}
		return false, i
	}
}
