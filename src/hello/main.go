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
    // Implicitly receive channels. Since we call three times we get three words
    word := <- wordChannel
    fmt.Printf("%s ", word)

    word = <- wordChannel
    fmt.Printf("%s ", word)

    word = <- wordChannel
    fmt.Printf("%s ", word)

    word = <- wordChannel
    fmt.Printf("%s ", word)

    // This one won't receive anything since the channel is already closed because it only has 4 words
    word = <- wordChannel
    fmt.Printf("%s \n", word)
}
