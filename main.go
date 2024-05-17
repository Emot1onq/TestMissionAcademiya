package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var result int
	fmt.Println("Введите строку")
	vvod, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	r := regexp.MustCompile("\\s+")
	vvod = r.ReplaceAllString(vvod, "")

	firstOperand, secondOperand, doit, flagRim := readString(vvod)
	if doit == "Alarm!!!" {
		fmt.Println("ВЫвожу панику и останавливаю программу!!!")

	} else {
		result = calc(firstOperand, secondOperand, doit)

		if flagRim {
			fmt.Println(translateTo(result))
		} else {
			fmt.Println(result)
		}

	}
}

func calc(firstOperand, secondOperand int, doit string) (result int) {

	switch doit {
	case "+":
		result = firstOperand + secondOperand
	case "-":
		result = firstOperand - secondOperand
	case "*":
		result = firstOperand * secondOperand
	case "/":
		result = firstOperand / secondOperand

	}

	return
}
func readString(vvod string) (firstOperand, secondOperand int, doit string, flagRim bool) {
	var flagArab bool

	r := regexp.MustCompile(`\+|\-\*|\/`)
	arr := r.Split(vvod, -1)
	doit = r.FindString(vvod)

	if strings.Contains(arr[0], "I") || strings.Contains(arr[0], "X") || strings.Contains(arr[0], "V") {
		flagRim = true
		firstOperand = translateFrom(arr[0])
	} else {
		flagArab = true
		firstOperand, _ = strconv.Atoi(arr[0])
	}

	if strings.Contains(arr[1], "I") || strings.Contains(arr[1], "X") || strings.Contains(arr[1], "V") {
		flagRim = true
		secondOperand = translateFrom(arr[1])
	} else {
		flagArab = true
		secondOperand, _ = strconv.Atoi(arr[1])
	}

	if flagRim && flagArab {
		return 0, 0, "Alarm!!!", flagRim
	}

	return
}
func translateFrom(item string) (result int) {
	switch item {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10

	}
	return 0
}

func translateTo(number int) string {
	romanValueList := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanCharList := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var res string
	for i := 0; i < len(romanValueList); i++ {
		for number >= romanValueList[i] {
			number -= romanValueList[i]
			res += romanCharList[i]
		}
	}
	return res
}

//func add(x, y int) int {
//    return x + y
//}
