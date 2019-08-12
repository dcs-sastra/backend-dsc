package handlers

import (
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
    //"encoding/json"
)

func GetImage(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("GET GetImage")
    params := mux.Vars(r)
    fmt.Println(params["id"])
}
