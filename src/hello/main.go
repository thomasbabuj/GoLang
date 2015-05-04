package main

import (
    "fmt"
    "time"
)

// Channel of channels
// Here the first param is channel of channels, so to inside this function we are creating a new channel for word and
// send the wordchannel back to the main function.

func emit(chanChannel chan chan string, done chan bool) {

    wordChannel := make(chan string)
    chanChannel <- wordChannel

    defer close(wordChannel)

    words := []string{"the", "quick", "brown", "fox"}
    i := 0

    t := time.NewTimer(3 * time.Second)

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
    channelCh:= make(chan chan string)
    doneCh := make(chan bool)

    go emit(channelCh, doneCh)

    wordCh := <-channelCh

    for word := range wordCh {
        fmt.Printf("%s ", word)
    }
}

