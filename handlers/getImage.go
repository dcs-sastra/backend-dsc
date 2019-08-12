package handlers

import (
    "net/http"
    "fmt"
    //"encoding/json"
)

func GetImage(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("GET GetImage")
}
