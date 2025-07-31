package conditionals

import "fmt"

func AreaRectangle() {
	area := rect(3.6, 6.3)
	fmt.Println(area)
}

func rect(w, l float64) float64 {
	return w * l
}
