package main

import "fmt"

func soma(a int, b int) int {

    return a + b
}

func main() {

    a := 5
    b := 5

    res := soma(a, b)
    fmt.Println("A soma de ", a, " + ", b, " = ", res)

}
