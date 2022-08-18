package main

import (
	"fmt"
)

var lowNames = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

var tensNames = []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

var bigNames = []string{"thousand", "million", "billion"}

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

func main() string {
	fmt.Println("================================================")
	fmt.Println("WELCOME TO THE BANK OF STEPHEN :)")
	fmt.Println("================================================")

	var decision string
	fmt.Println("Do you want to run a transfer?")
	fmt.Scan(&decision)

	switch decision {
	case "n":
		fmt.Println("Kindly move to the appropraite section or meet the security man to help direct you. Thank you!!!")
	case "y":
		var typeOfAmount int
		fmt.Println("Choose the mode you want to fill (press 1 or 2) \n" +
			"1. I am typing the amount in numbers \n" +
			"2. I am typing the amount in words \n")
		fmt.Scan(&typeOfAmount)

		switch typeOfAmount {
		case 1:
			var amount int
			fmt.Print("Kindly type the amount you wish to transfer: ")
			fmt.Scan(&amount)

			fmt.Println("================================================")
			fmt.Println("The Amount in figure: \n", amount)
			fmt.Println("Amount in Words: \n", convertNum2Words(amount))
			fmt.Println("================================================")
		default:
			fmt.Print("Nice")
		}

	}
}
