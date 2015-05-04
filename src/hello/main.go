/*
   Multiple Readers, Multiple Writers
   In this program we will fetch list of webpages and returns the size
 */

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func getPage(url string) (int, error) {
    // Using GET method from http
    resp, err := http.Get(url)

    if err != nil {
        return 0, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return 0, err
    }

    return len(body), nil
}

// multiple go routines to receive url

func worker( urlCh chan  string, sizeCh chan string) {
    for {
        url := <-urlCh
            length, err := getPage(url)
            if err == nil {
                sizeCh <- fmt.Sprintf("%s has length %d", url ,length)
            } else {
                sizeCh <- fmt.Sprintf("Error getting %s : %s", url ,err)
            }
    }
}

func main() {

    urlCh := make(chan string)
    sizeCh := make(chan string)

    go worker(urlCh, sizeCh)

    urlCh <- "http://www.oreilly.com/"
    fmt.Printf("%s \n", <-sizeCh)

}
