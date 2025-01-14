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


    message := "Бессмысленный код."
    for _, ch := range message {
        fmt.Print(string(ch) + " ")
    }
}