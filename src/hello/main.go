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
    // Here we are using range to receive from channel
    for word := range wordChannel {
        fmt.Printf("%s ", word)
    }

    fmt.Printf("\n")
}
