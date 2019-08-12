package handlers

import (
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
    //"encoding/json"
)

func GetMember(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("GET GetMember")
    params := mux.Vars(r)
    fmt.Println(params["id"])
}
