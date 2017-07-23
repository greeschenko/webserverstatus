package main

import (
    //"fmt"
    "webserverstatus/libs/configread"
    "os/exec"
    "strings"
    "encoding/json"
    "net/http"
)

type Main struct {
    Domeninfo string
}

type Status struct {
    Status string
}

type Stat struct {
    Name   string
    Graph  string
    Stats  string
    Status string
}

var Conf configread.Config

func executeOne(command string) string {
    res := strings.Split(command," ")

    out, err := exec.Command(res[0], res[1:]...).Output()
    if err != nil {
        panic("error")
    }

    return string(out)
}

func toJson(i interface{}) []byte {
    res, err := json.MarshalIndent(i, "", "  ")
    if err != nil {
        panic (err)
    }
    return res
}

func main() {
    Conf = configread.Read()

    myMux := http.NewServeMux()
    myMux.HandleFunc("/api/main", srvMain)
    myMux.HandleFunc("/api/status", srvStatus)
    myMux.HandleFunc("/api/stat", srvStat)
    http.ListenAndServe(":1888", myMux)
}

func srvMain(w http.ResponseWriter, r *http.Request) {
    data := Main{Conf.Server.Domen}
    w.Write(toJson(data))
}

func srvStatus(w http.ResponseWriter, r *http.Request) {
    data := Status{"OK"}
    w.Write(toJson(data))
}

func srvStat(w http.ResponseWriter, r *http.Request) {
    var statdata []Stat
    for s := range Conf.Stats {
        item := Conf.Stats[s]
        statdata = append(statdata, Stat{
            item.Name,
            "test",
            executeOne(item.Condition),
            executeOne(item.Status),
        })
    }
    w.Write(toJson(statdata))
}
