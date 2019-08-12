package handlers

import (
    "net/http"
    "fmt"
    //"encoding/json"
)

func GetMembers(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("GET GetMembers")
}
