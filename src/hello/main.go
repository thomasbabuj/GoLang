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

    // this code got deadlock error mainly because, this range wordchannel is waiting for
    // response from emit channel, even though the channel finished running.
    for word := range wordChannel {
        fmt.Printf("%s ", word)
    }

}
