package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var roman = [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

func main() {
	var exp string
	fmt.Println("Please type your expression")
	fmt.Scanf("%s", &exp)
	n1, sig, n2 := SplitExpression(exp)
	n12, n22, t := Transformation(n1, n2)
	res := Calculator(n12, sig, n22)
	if t == 2 {
		res1 := RomanResultConversion(res)
		fmt.Println(res1)
	} else {
		fmt.Println(res)
	}

}

func Calculator(a int, c string, b int) int {
	var res int
	switch {
	case c == "+":
		res = a + b
	case c == "-":
		res = a - b
	case c == "/":
		res = a / b
	case c == "*":
		res = a * b
	}

	return res
}

func SplitExpression(s string) (string, string, string) {
	var split []string
	var sig, n1, n2 string
	if strings.Count(s, "+") == 1 && strings.ContainsAny(s, "-*/") == false {
		split = strings.Split(s, "+")
		sig = "+"
	} else if strings.Count(s, "-") == 1 && strings.ContainsAny(s, "+*/") == false {
		split = strings.Split(s, "-")
		sig = "-"
	} else if strings.Count(s, "/") == 1 && strings.ContainsAny(s, "-+*") == false {
		split = strings.Split(s, "/")
		sig = "/"
	} else if strings.Count(s, "*") == 1 && strings.ContainsAny(s, "-+/") == false {
		split = strings.Split(s, "*")
		sig = "*"
	} else {
		fmt.Println("ERROR: unaviable operation")
		os.Exit(1)
	}
	n1 = split[0]
	n2 = split[1]
	return n1, sig, n2
}

func Transformation(a string, b string) (int, int, int) {
	t := 0
	var a1, b1 int
	for i := 0; i < 10; i++ {
		if a == roman[i] {
			t++
			a1 = i + 1

		}
		if b == roman[i] {
			t++
			b1 = i + 1
		}
	}
	if t == 0 {
		k := TestGreaterThanX(a, b)
		if k == false {
			a2, _ := strconv.Atoi(a)
			b2, _ := strconv.Atoi(b)
			if a2 <= 10 && b2 <= 10 {
				return a2, b2, t
			} else {
				fmt.Println("ERROR: cannot use numbers greater than 10")
				os.Exit(2)
			}
		} else if k == true {
			fmt.Println("ERROR: cannot use numbers greater than X")
			os.Exit(2)
		}
	} else if t == 1 {
		fmt.Println("ERROR: cannot use numbers of different systems")
		os.Exit(3)
	}
	return a1, b1, t
}

func RomanResultConversion(n int) (s string) {
	var n1 string
	if n <= 0 {
		fmt.Println("ERROR: roman numbers can only be positive, the result of calculation does nit match the condition")
		os.Exit(4)
	} else if n > 0 && n <= 10 {
		n1 = roman[n-1]
	} else if n > 10 && n <= 20 {
		n1 = "X" + roman[n-11]
	} else if n > 20 && n <= 30 {
		n1 = "XX" + roman[n-21]
	} else if n > 30 && n < 40 {
		n1 = "XXX" + roman[n-31]
	} else if n >= 40 && n < 50 {
		n1 = "XL" + roman[n-41]
	} else if n >= 50 && n <= 60 {
		n1 = "L" + roman[n-51]
	} else if n > 60 && n <= 70 {
		n1 = "LX" + roman[n-61]
	} else if n > 70 && n <= 80 {
		n1 = "LXX" + roman[n-71]
	} else if n > 80 && n < 90 {
		n1 = "LXXX" + roman[n-81]
	} else if n >= 90 && n < 100 {
		n1 = "XC" + roman[n-91]
	} else if n == 100 {
		n1 = "C"
	}

	return n1

}

func TestGreaterThanX(s1 string, s2 string) bool {
	t := true
	switch {
	case t == strings.Contains(s1, "X"):
		return t
	case t == strings.Contains(s1, "I"):
		return t
	case t == strings.Contains(s1, "V"):
		return t
	case t == strings.Contains(s1, "L"):
		return t
	case t == strings.Contains(s1, "C"):
		return t
	case t == strings.Contains(s2, "I"):
		return t
	case t == strings.Contains(s1, "X"):
		return t
	case t == strings.Contains(s1, "V"):
		return t
	case t == strings.Contains(s1, "L"):
		return t
	case t == strings.Contains(s1, "C"):
		return t
	default:
		t = false
	}
	return t
}
