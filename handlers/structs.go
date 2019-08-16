package handlers

type Confirm struct {
    Status string `json: status`
    Text string `json: text`
}

type Member struct{
    Id string
    Team string
    Name string
    Branch string
}
