package main

import (
    "fmt"
)

func makeId(idChan chan int){
    var id int
    id = 0

    for {
        idChan <- id
        id += 1
    }
}

func main() {
    idChan := make(chan int)

    go makeId(idChan)

    fmt.Printf("%d \n", <-idChan)
    fmt.Printf("%d \n", <-idChan)
    fmt.Printf("%d \n", <-idChan)
}
