package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    MakeRequest()
}

func MakeRequest() {
req, err := http.NewRequest("GET", "https://api.newrelic.com/v2/alerts_policies.json", nil)
if err != nil {
    // handle err
}

req.Header.Set("X-Api-Key", os.Getenv("API"))
resp, err := http.DefaultClient.Do(req)
if err != nil {
    // handle err
}
defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}
