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

func getter( url string, size chan string) {
    length, err := getPage(url)
    if err == nil {
        size <- fmt.Sprintf("%s has length %d", url ,length)
    }
}

func main() {
    urls := []string{
        "http://www.google.com",
        "http://www.yahoo.com",
        "http://www.bing.com",
        "http://www.bbc.co.uk",
    }

    size := make(chan string)

    for _, url := range urls {
        go getter(url, size)
    }

    for i:=0; i< len(urls); i++ {
        fmt.Printf("%s \n", <-size )
    }

}
