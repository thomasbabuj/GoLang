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

func worker( urlCh chan  string, sizeCh chan string, workerId int) {
    for {
        url := <-urlCh
        length, err := getPage(url)
        if err == nil {
            sizeCh <- fmt.Sprintf("%s has length %d ( %d )", url ,length, workerId)
        } else {
            sizeCh <- fmt.Sprintf("Error getting %s : %s", url ,err)
        }
    }
}

func main() {

    urlCh := make(chan string)
    sizeCh := make(chan string)

    for i:=0; i < 10; i++ {
        go worker(urlCh, sizeCh, i)
    }

    urls := []string{
        "http://www.google.com",
        "http://www.yahoo.com",
        "http://www.bing.com",
        "http://www.bbc.co.uk",
        "http://www.oreilly.com/",
    }

    for _, url:= range urls {
        urlCh <- url
    }

    for i:=0; i <len(urls); i++ {
        fmt.Printf("%s \n", <-sizeCh )
    }


}
