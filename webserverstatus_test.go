package main

import (
    "testing"
)

//test prop read config file
//test '/data/status.json',
//test '/data/stat.json',
//test '/data/main.json'


func TestPropReadConfigFile(t *testing.T) {
    getConfig()
    if Domen != "polonex.com.ua" {
        t.Fatal("not prop server domen")
    }
    if len(Commands) == 0 {
        t.Fatal("0 commands in config")
    }
    return
}

