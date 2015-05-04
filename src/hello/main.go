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

}

func main() {
    // Creating a channel
    wordChannel := make(chan string)

    // starting multiple channels
    go emit(wordChannel)
    go emit(wordChannel)

    for word := range wordChannel {
        fmt.Printf("%s ", word)
    }

}
