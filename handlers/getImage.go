package handlers

import (
    "net/http"
    "fmt"
    "os"
    "io"
    "github.com/gorilla/mux"
    //"encoding/json"
)

func GetImage(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("GET GetImage")
    params := mux.Vars(r)
    fmt.Println(params["id"])
    myfile :=  fmt.Sprintf("./images/%s.jpg",params["id"])
    f,err := os.Open(myfile)
    if err !=nil {
        fmt.Println(err)
    }
    defer f.Close()
    w.Header().Set("Content-type","image/jpeg")
    io.Copy(w,f)
}
