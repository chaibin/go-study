package main

import "fmt"

func main() {
	a, b := 12, 13
	c := eval(a, b, "+")
	fmt.Println(c)
	if d, err := eval1(a, b, "22"); err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println(d)
	}
}

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("invalid op")
	}

}

func eval1(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a + b, nil
	case "*":
		return a + b, nil
	case "/":
		return a + b, nil
	default:
		return 0, fmt.Errorf("invalid op")
	}
}
