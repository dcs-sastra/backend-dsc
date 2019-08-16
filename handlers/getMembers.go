package handlers

import (
    "net/http"
    "fmt"
    "os"
    "encoding/json"
    "encoding/csv"
)

func GetMembers(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("GET GetMembers")
    f,erro := os.Open("./data/members.csv")
    if erro!=nil{
        fmt.Println(erro)
    }
    defer f.Close()
    lines,errc := csv.NewReader(f).ReadAll()
    if errc!=nil{
        fmt.Println(errc)
    }
    members := make([]Member,0)
    for _,line := range lines{
        data := Member{
            Id: line[0],
            Team: line[1],
            Name: line[2],
            Branch: line[3],
            Links: line[4],
        }
        members = append(members,data)
    }
    dataJson,errj := json.Marshal(members)
    if errj!=nil{
        panic(errj)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(dataJson)
}
