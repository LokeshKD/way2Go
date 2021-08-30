package main
import "fmt"

func main() {
    Print(1, 3, 2, 0)
}

func Print(a...int) { // variable number of parameters
    if len(a) == 0 {
        return
    }

    for _, v := range a {
        fmt.Printf("The number is: %d\n", v)
    }
    return
}
