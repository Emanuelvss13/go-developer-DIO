package calculadora

import "fmt"

func sum(values ...int) int {

	total := 0

	for _, value := range values {

		total += value

	}

	return total
}

func subtration(baseValue int, values ...int) int {

	for _, value := range values {

		baseValue -= value

	}

	return baseValue
}

func multiplication(values ...int) int {

	total := 1

	for _, value := range values {

		total *= value

	}

	return total
}

func division(a int, b int) int {
	return a / b
}

func main() {

	fmt.Println(sum(5, 5, 5, 5))           //20
	fmt.Println(subtration(100, 20, 40))   //40
	fmt.Println(multiplication(10, 10, 5)) //500
	fmt.Println(division(9, 3))            //3

}
