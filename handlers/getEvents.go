package handlers

import (
    "net/http"
    "fmt"
    "os"
    "encoding/json"
    "encoding/csv"
)

func GetEvents(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("GET GetEvents")
    f,erro := os.Open("./data/events.csv")
    if erro!=nil{
        fmt.Println(erro)
    }
    defer f.Close()
    lines,errc := csv.NewReader(f).ReadAll()
    if errc!=nil{
        fmt.Println(errc)
    }
    events := make([]Event,0)
    for _,line := range lines{
        event := Event{
            Name: line[0],
            Desc: line[1],
            Speakers: line[2],
            Date: line[3],
        }
        events = append(events,event)
    }
    dataJson,errj := json.Marshal(events)
    if errj!=nil{
        panic(errj)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(dataJson)
}
