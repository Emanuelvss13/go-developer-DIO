package main

import "fmt"

const kelvin = 373.2

func main() {

	celcius := kelvin - 273.2

	fmt.Printf("A temperatura de ebulição da agua em Kelvin é: %g, convertido para celcius fica: %g", kelvin, celcius)

}
