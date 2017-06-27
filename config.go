package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type Config struct {
    Server Server
    Stats Stats
}

type Server struct {
    Domen string
}

type Stats []struct {
    Name string
    Command string
}


// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    var c Config

    dat, err := ioutil.ReadFile("wss.cfg")
    check(err)

    jerr := json.Unmarshal(dat, &c)
    if jerr != nil {
        fmt.Println("error:", jerr)
    }

    fmt.Println(c)
}
