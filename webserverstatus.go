package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"webserverstatus/libs/configread"
	"webserverstatus/libs/imggen"
)

type Main struct {
	Domeninfo string
}

type Status struct {
	Status string
}

type Sites struct {
	Domen  string
	Status string
	Data   []Stat
}

type Stat struct {
	Name   string
	Graph  string
	Stats  string
	Status string
}

var Conf configread.Config

func getConfig() {
	Conf = configread.Read()
}

func executeOne(command string) string {
	res := strings.Split(command, " ")

	out, err := exec.Command(res[0], res[1:]...).Output()
	if err != nil {
		panic(err)
	}

	return string(out)
}

func toJson(i interface{}) []byte {
	res, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		panic(err)
	}
	return res
}

func main() {
	getConfig()

	myMux := http.NewServeMux()
	myMux.HandleFunc("/api/main", srvMain)
	myMux.HandleFunc("/api/status", srvStatus)
	myMux.HandleFunc("/api/stat", srvStat)
	myMux.HandleFunc("/api/sites", srvSites)
	myMux.HandleFunc("/api/sites-status", srvSitesStatus)
	myMux.HandleFunc("/api/sites-stats", srvSitesStats)
	http.ListenAndServe(":1888", myMux)
}

func srvMain(w http.ResponseWriter, r *http.Request) {
	data := Main{Conf.Server.Ip}
	w.Write(toJson(data))
}

func srvStatus(w http.ResponseWriter, r *http.Request) {
	var data Status
	tmp := executeOne("bash scripts/pingstatus.sh " + Conf.Server.Ip)
	if tmp == "fail" {
		data = Status{"FAIL"}
	} else {
		data = Status{"OK"}
	}
	w.Write(toJson(data))
}

var ImgList = make(map[string][]int, len(Conf.Stats))

func srvStat(w http.ResponseWriter, r *http.Request) {
	var statdata []Stat
	for _, item := range Conf.Stats {
		var buffer bytes.Buffer
		buffer.WriteString("/img/")
		buffer.WriteString(item.Name)
		buffer.WriteString(".png")

		gval, err := strconv.ParseFloat(executeOne(item.Graphs[0].Command), 64)
		if err != nil {
			panic(err)
		}

		gmax, err := strconv.ParseFloat(executeOne(item.Graphs[0].Max), 64)
		if err != nil {
			panic(err)
		}

		if _, ok := ImgList[item.Name]; !ok {
			ImgList[item.Name] = make([]int, 20)
		}

		ImgList[item.Name] = imggen.Gen(
			ImgList[item.Name],
			gval,
			gmax,
			"web"+buffer.String(),
		)

		statdata = append(statdata, Stat{
			item.Name,
			buffer.String() + "?time=" + time.Now().Format("20060102150405"),
			executeOne(item.Condition),
			executeOne(item.Status),
		})
	}

	w.Write(toJson(statdata))
}

func srvSites(w http.ResponseWriter, r *http.Request) {
	var sitedata []Sites
	for _, site := range Conf.Sites {
		var item = Sites{
			Domen: site.Domen,
		}
		sitedata = append(sitedata, item)
	}
	w.Write(toJson(sitedata))
}

func srvSitesStatus(w http.ResponseWriter, r *http.Request) {
	//index := r.URL.Query().Get("index")
	//id, err := strconv.Atoi(index)
	//if err != nil {
	//panic(err)
	//}
	data := Status{"OK"} //TODO add status calculation
	w.Write(toJson(data))
}

func srvSitesStats(w http.ResponseWriter, r *http.Request) {
	var statdata []Stat
	index := r.URL.Query().Get("index")
	id, err := strconv.Atoi(index)
	if err != nil {
		panic(err)
	}
	for _, item := range Conf.Sites[id].Stats {
		var buffer bytes.Buffer
		buffer.WriteString("/img/")
		buffer.WriteString(item.Name)
		buffer.WriteString(".png")

		gval, err := strconv.ParseFloat(executeOne(item.Graphs[0].Command), 64)
		if err != nil {
			panic(err)
		}

		gmax, err := strconv.ParseFloat(executeOne(item.Graphs[0].Max), 64)
		if err != nil {
			panic(err)
		}

		if _, ok := ImgList[item.Name]; !ok {
			ImgList[item.Name] = make([]int, 20)
		}

		ImgList[item.Name] = imggen.Gen(
			ImgList[item.Name],
			gval,
			gmax,
			"web"+buffer.String(),
		)

		statdata = append(statdata, Stat{
			item.Name,
			buffer.String() + "?time=" + time.Now().Format("20060102150405"),
			executeOne(item.Condition),
			executeOne(item.Status),
		})
	}

	w.Write(toJson(statdata))
}
