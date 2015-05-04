package main

import (
    "fmt"
    "time"
)

// select key word will allow go program to listen and communicate to multiple channel at the same time

func emit(wordChannel chan string, done chan bool) {
    words := []string{"the", "quick", "brown", "fox"}
    i := 0

    t := time.NewTimer(3 * time.Second)

    defer close(wordChannel)

    // infinite loop, if somebody is listening to receive  then they will get the next word
    // and if it receives done channel then it terminates the channel.
    for {
        select {
            case wordChannel <- words[i] :
                i += 1
                if i == len(words) {
                    i = 0
                }

            case <- done:
                fmt.Printf(" Go done \n")
                done<-true
                return
            case <- t.C :
                fmt.Printf(" \n Timer terminates the channel \n")
                return
        }
    }
}

func main() {
    wordCh := make(chan string)
    doneCh := make(chan bool)

    go emit(wordCh, doneCh)

    for word := range wordCh {
        fmt.Printf("%s ", word)
    }
}

