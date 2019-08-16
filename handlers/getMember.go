package handlers

import (
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
    "encoding/csv"
    "encoding/json"
    "os"
)

func GetMember(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("GET GetMember")
    params := mux.Vars(r)
    fmt.Println(params["id"])
    id := params["id"]
    f,erro := os.Open("./data/members.csv")
    if erro!=nil{
        fmt.Println(erro)
    }
    defer f.Close()
    lines,errc := csv.NewReader(f).ReadAll()
    if errc!=nil{
        fmt.Println(errc)
    }
    for _,line := range lines{
        if(line[0] == id){
            data := Member{
                Id: line[0],
                Team: line[1],
                Name: line[2],
                Branch: line[3],
            }
            dataJson,errj := json.Marshal(data)
            if errj!=nil{
                panic(errj)
            }
            w.Header().Set("Content-Type","application/json")
            w.WriteHeader(http.StatusOK)
            w.Write(dataJson)
            break
        }
    }
}
