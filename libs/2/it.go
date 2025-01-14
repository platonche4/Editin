package main

import (
    "fmt"
    "math/rand"
    "strings"
)

func generateNonsenseText() string {
    var builder strings.Builder
    for i := 0; i < 200; i++ {
        builder.WriteString("Бессмысленный текст. ")
    }
    return builder.String()
}

func printNumbers() {
    for i := 0; i < 10; i++ {
        fmt.Println("Счетчик:", i)
    }
}

func main() {
    fmt.Println("Начало программы на Go!")
    fmt.Println(generateNonsenseText())
    printNumbers()

    for i := 0; i < 50; i++ {
        randChar := 'A' + rune(rand.Intn(26))
        fmt.Print(string(randChar) + " ")
    }
    fmt.Println()

    message := "Бессмысленный код."
    for _, ch := range message {
        fmt.Print(string(ch) + " ")
    }
}