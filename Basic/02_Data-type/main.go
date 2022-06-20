package main

import "fmt"

func main() {
    // whole number
    // example : 2, 5, 10
    var wholeNumber int  = 10

    // real number 
    // example: 2.45, 10.22
    // we can use float32 or float64
    var realNumber float64 = 20.15


    // character
    // example : 's', 't', '#'
    var char rune = 'T'

    // string
    var msg string = "HI There"

    // boolean
    // true or false
    var isRaining bool = true

    fmt.Println(wholeNumber)
    fmt.Println(realNumber)
    fmt.Println(char)
    fmt.Println(msg)
    fmt.Println(isRaining)
}
