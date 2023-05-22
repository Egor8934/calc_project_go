package calcHelpers

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	INCORRECT    = "SYNTAX ERROR: The data is incorrect.You can enter only Arabic or only Roman numerals and only one operator(+|-|*|/)!Try again!"
	SCALE        = "SYNTAX ERROR: You are trying to use different number systems at the same time!!!"
	EXPRESSION   = "EXPRESSION IS INCORRECT!Use two operands and one operator, and only positive numbers!"
	RANGE        = "This Calculator operates only on positive integers from 1 to 10 inclusive"
	ROMANOPERAND = "One of the operands is invalid or greater than the allowed value"
	NEGATIVENUM  = "Value is ZERO or negative number! In Roman numeral system there are no negative value and null!"
)

var romanNums = map[string]int{
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}
var intToRoman10To100 = map[rune]string{
	'1': "X",
	'2': "XX",
	'3': "XXX",
	'4': "XL",
	'5': "L",
	'6': "LX",
	'7': "LXX",
	'8': "LXX",
	'9': "XC",
}
var intToRoman1To10 = map[rune]string{
	'1': "I",
	'2': "II",
	'3': "III",
	'4': "IV",
	'5': "V",
	'6': "VI",
	'7': "VII",
	'8': "VIII",
	'9': "IX",
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func calculate(left int, right int, operator string) int {
	switch operator {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	default:
		panic(INCORRECT)
	}
}

func convertResult(n int, isRoman bool) {
	if !isRoman {
		fmt.Println("Result: ", n)
		return
	}
	if n <= 0 {
		panic(NEGATIVENUM)
	}

	str := strconv.Itoa(n)

	var result string
	if len(str) == 3 {
		result = "C"
		fmt.Println("Result: ", result)
		return
	}
	if len(str) == 1 {
		result = intToRoman1To10[rune(str[0])]
		fmt.Println("Result: ", result)
		return
	}
	if len(str) == 2 {
		if rune(str[1]) == '0' {
			result = intToRoman10To100[rune(str[0])]
		} else {
			result = intToRoman10To100[rune(str[0])] + intToRoman1To10[rune(str[1])]
		}
		fmt.Println("Result: ", result)
		return
	}
}

func parseStr(s string, isRoman bool) {
	var pattern string
	if isRoman {
		pattern = `^[XVI]*[-+/*][XVI]*$`
	} else {
		pattern = `^\d*[-+/*]\d*$`
	}

	matchedExp, err := regexp.Match(pattern, []byte(s))
	checkErr(err)
	if !matchedExp {
		panic(EXPRESSION)
	}

	var operator string
	for _, val := range s {
		if val == '-' || val == '+' || val == '*' || val == '/' {
			operator = string(val)
			break
		}
	}
	parseExpr := strings.Split(s, operator)

	if isRoman {
		left, ok := romanNums[parseExpr[0]]
		if !ok {
			panic(ROMANOPERAND)
		}
		right, ok := romanNums[parseExpr[1]]
		if !ok {
			panic(ROMANOPERAND)
		}
		convertResult(calculate(left, right, operator), isRoman)
		return

	} else {
		left, err := strconv.Atoi(parseExpr[0])
		checkErr(err)
		if left > 10 || left == 0 {
			panic(RANGE)
		}

		right, err := strconv.Atoi(parseExpr[1])
		checkErr(err)
		if right == 0 || right > 10 {
			panic(RANGE)
		}

		convertResult(calculate(left, right, operator), isRoman)
		return
	}
}

func checkString(s string) {
	patternAll := `[^-*/+\dXVI]`
	matchAll, err := regexp.Match(patternAll, []byte(s))
	checkErr(err)

	if matchAll {
		panic(INCORRECT)
	}
	patternDigit := `^[+-/*\d]*$`
	patternRoman := `^[+-/*XVI]*$`

	matchedDig, err := regexp.Match(patternDigit, []byte(s))
	checkErr(err)
	if matchedDig {
		parseStr(s, false)
		return
	}

	matchedRom, err := regexp.Match(patternRoman, []byte(s))
	checkErr(err)
	if matchedRom {
		parseStr(s, true)
		return
	}
	panic(SCALE)
}

func Welcome() {
	fmt.Println("Welcome to my-calculator-app! If you want to exit, please enter `q|Q`")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		s := strings.ReplaceAll(input, " ", "")
		s = strings.TrimSpace(s)
		s = strings.ToUpper(s)

		if s == "Q" {
			fmt.Println("Bye...")
			return
		}
		checkString(s)
	}
}
