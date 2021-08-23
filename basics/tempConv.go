package main
import "fmt"
//import "strconv"
//import "encoding/json"

// aliasing type
type Celsius float32
type Fahrenheit float32

// Function to convert celsius to fahrenheit
func toFahrenheit(t Celsius) Fahrenheit {
	var temp Fahrenheit
	temp = Fahrenheit((t*9/5 )+ 32)
	return temp
}

func main() {
    var tempCelsius Celsius = 100

    tempFahr := toFahrenheit(tempCelsius)	// function call
    fmt.Printf("%f ˚C is equal to %f ˚F \n",tempCelsius,tempFahr)
}
