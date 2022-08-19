package main

import (
	"fmt"
)

var lowNames = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

var tensNames = []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

var bigNames = []string{"thousand", "million", "billion"}

var NumWords = map[string]map[int]string{
	"lowNumbers": {
		0: "zero", 1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight",
		9: "nine", 10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen",
		17: "seventeen", 18: "eighteen", 19: "nineteen",
	},

	"tens": {20: "twenty", 30: "thirty", 40: "forty", 50: "fifty", 60: "sixty", 70: "seventy", 80: "eighty", 90: "ninety"},

	"hundred": {100: "hundred"},

	"thousand": {1000: "thousand"},

	"Million": {1000000: "million"},

	"billion": {1000000000: "billion"},
}

func convertWords(words string) int {
	for k, _ := range NumWords {
		if k != words {
			for k, v := range NumWords["tens"] {
				if v == words {
					fmt.Println("The Amount in figure is: ", k)
				}
				return k
			}
		} else if k != words {
			for k, v := range NumWords["thousand"] {
				if v == words {
					fmt.Println("The Amount in figure is: ", k)
				}
			}
		}
	}
	return convertWords(words)
}

func convert999(num int) string {
	s1 := lowNames[num/100] + " hundred"
	s2 := convert99(num % 100)
	if num <= 99 {
		return s2
	}
	if num == 0 {
		return s1
	} else {
		return s1 + " " + s2
	}
}

func convert99(num int) string {
	if num < 20 {
		return lowNames[num]
	}
	s := tensNames[num/10-2]
	if num == 0 {
		return s
	}
	return s + "-" + lowNames[num]
}

func convertNum2Words(num int) string {
	if num < 0 {
		return "negative " + convertNum2Words(-num)
	}

	if num <= 999 {
		return convert999(num)
	}

	s := ""
	t := 0
	for num > 0 {
		if num != 0 {
			s2 := convert999(num % 1000)
			if t > 0 {
				s2 = s2 + " " + bigNames[t-1]
			}
			if s == "" {
				s = s2
			} else {
				s = s2 + ", " + s
			}
		}
		num /= 1000
		t++
	}
	return s
}

func main() {
	fmt.Println("================================================")
	fmt.Println("WELCOME TO THE BANK OF STEPHEN :)")
	fmt.Println("================================================")

	var decision string
	fmt.Println("Do you want to run a transfer? (y/n)")
	fmt.Scan(&decision)

	switch decision {
	case "n":
		fmt.Println("Kindly move to the appropraite section or meet the security man to help direct you. Thank you!!!")
	case "y":
		var typeOfAmount int
		fmt.Println("Choose the mode you want to fill (press 1 or 2) \n" +
			"1. I am typing the amount in numbers \n" +
			"2. I am typing the amount in words")
		fmt.Scan(&typeOfAmount)

		switch typeOfAmount {
		case 1:
			var amount int
			fmt.Print("Kindly type the amount you wish to transfer: ")
			fmt.Scan(&amount)

			fmt.Println("================================================")
			fmt.Println("The Amount in figure: \n", amount)
			fmt.Println("Amount in Words: \n", convertNum2Words(amount), "NGN")
			fmt.Println("================================================")

		case 2:
			var words string
			fmt.Print("Kindly type the amount you wish to transfer: ")
			fmt.Scan(&words)

			fmt.Println(convertWords(words))

		default:
			fmt.Print("Nice")
		}

	}
}
