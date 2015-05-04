package main

import (
    "fmt"
)

func emit(c chan string) {
    words := []string{"the", "quick", "brown", "fox"}

    for _, word := range words {
        // sending to channel
        c  <- word
    }

    close(c)
}

func main() {
    // Creating a channel
    wordChannel := make(chan string)

    go emit(wordChannel)

    // receive the things comes from the wordChannel
    // Implicitly receive channels and its only receive first word
    word := <- wordChannel
    fmt.Printf("%s\n", word)
}
