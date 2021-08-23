package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main() {
    var str string = "This is an example of a string"
    fmt.Printf("T/F? \nDoes the string \"%s\" have prefix %s? ", str, "Th")
    fmt.Printf("\n%t\n\n", strings.HasPrefix(str, "Th"))    // Finding prefix

    fmt.Printf("Does the string \"%s\" have suffix %s? ", str, "ting")
    fmt.Printf("\n%t\n\n", strings.HasSuffix(str, "ting"))  // Finding suffix

    str = "Hi, I'm Marc, Hi."
    fmt.Printf("The position of the first instance of\"Marc\" is: ")
    fmt.Printf("%d\n", strings.Index(str, "Marc"))      // Finding first occurence
    fmt.Printf("The position of the first instance of \"Hi\" is: ")
    fmt.Printf("%d\n", strings.Index(str, "Hi"))        // Finding first occurence
    fmt.Printf("The position of the last instance of \"Hi\" is: ")
    fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))    // Finding last occurence
    fmt.Printf("The position of the first instance of\"Burger\" is: ")
    fmt.Printf("%d\n", strings.Index(str, "Burger"))    // Finding first occurence

    str = "Hello, how is it going, Hugo?"
    var manyG = "gggggggggg"
    fmt.Printf("Number of H's in %s is: ", str)
    fmt.Printf("%d\n", strings.Count(str, "H"))         // count occurences
    fmt.Printf("Number of double g's in %s is: ", manyG)
    fmt.Printf("%d\n", strings.Count(manyG, "gg"))      // count occurences


    var origS string = "Hi there! "
    var newS string
    newS = strings.Repeat(origS, 3)     // Repeating origS 3 times
    fmt.Printf("The new repeated string is: %s\n", newS)

    var orig string = "666"
    var an int
    fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
    an, _ = strconv.Atoi(orig)  // converting to number
    fmt.Printf("The integer is: %d\n", an)
    an = an + 5
    newS = strconv.Itoa(an)     // converting to string
    fmt.Printf("The new string is: %s\n", newS)
}
