package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
    "webserverstatus/libs/configread"
    "webserverstatus/libs/imggen"
    "math/rand"
)

var imgList = make([]int, 20)

type Data struct {
    Users users
    Count int
}

type users map[string]user

type user struct {
    Name, Lastname, Email string
}

func getJsonResp() ([]byte, error) {
    var conf = configread.Read()
    user1 := user{"Yana","Greeschenko","bkacklsdjls@"+conf.Server.Domen}
    user2 := user{"Aleksey","Greeschenko","greeschenko@"+conf.Server.Domen}
    users := users{
        "1":user1,
        "2":user2,
    }
    data := Data{
        users,
        len(users),
    }

    return json.MarshalIndent(data, "", "  ")
}

func srvGetUsers(w http.ResponseWriter, r *http.Request) {
    response, err := getJsonResp()
    if err != nil {
        panic(err)
    }

    imgList = imggen.Gen(imgList, rand.Intn(50));

    w.Write(response)
}

//func main() {
    //myMux := http.NewServeMux()
    //myMux.HandleFunc("/api/users", srvHome)
    //http.ListenAndServe(":1888", myMux)
//}

func main() {
    r := mux.NewRouter().StrictSlash(false)
    r.HandleFunc("/api/users", srvGetUsers).Methods("GET")
    /*r.HandleFunc("/api/users/add", srvAddUsers).Methods("POST")
    r.HandleFunc("/api/users/update/{id}", srvUpdateUsers).Methods("PUT")
    r.HandleFunc("/api/users/delete/{id}", srvDeleteGetUsers).Methods("DELETE")*/

    server := &http.Server{
        Addr: ":1888",
        Handler: r,
    }

    fmt.Println("Listening...")
    server.ListenAndServe()
}

