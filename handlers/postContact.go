package handlers

import(
    "net/http"
    "fmt"
    "os"
    "encoding/json"
    "encoding/csv"
)

func PostContact(w http.ResponseWriter, r *http.Request){
    form := Contact{}
    response := Confirm{
        Status: "Success",
    }
    err := json.NewDecoder(r.Body).Decode(&form);
    if(err!=nil){
        fmt.Println(err)
        response.Status = "Error"
    }
    fmt.Println(form.Email,form.Phone,form.Message)
    f,erro := os.OpenFile("./data/contactus.csv",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0666)
    if(erro!=nil){
        fmt.Println(erro)
        response.Status = "Error"
    }
    defer f.Close()
    writer := csv.NewWriter(f)
    defer writer.Flush()
    data := []string{form.Email,form.Phone,form.Message}
    fmt.Println(data)
    errw := writer.Write(data)
    if(errw !=nil){
        fmt.Println(errw)
        response.Status = "Error"
    }
    statusJson,errj := json.Marshal(response)
    if errj!=nil{
        panic(errj)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}
