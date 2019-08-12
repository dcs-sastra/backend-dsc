package handlers

import (
    "net/http"
    "fmt"
    //"encoding/json"
)

func GetImages(w http.ResponseWriter, r *http.Request)  {
    fmt.Println("GET GetImages")
}
