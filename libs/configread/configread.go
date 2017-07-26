package configread

import (
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

type Graph struct {
    Command string
    Max string
}

type Stats []struct {
    Name string
    Condition string
    Status string
    Graphs []Graph
}


// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Read() Config{
    var c Config

    dat, err := ioutil.ReadFile("package.cfg")
    check(err)

    jerr := json.Unmarshal(dat, &c)
    check(jerr)

    return c
}
