package conditionals

import "fmt"

func AreaTriangle() {
	area := calculate(6.2, 4.5)
	fmt.Println(area)
}

func calculate(b, h float64) float64 {
	return 0.5 * b * h
}
