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
    Links string
}

type Event struct{
    Name string
    Desc string
    Speakers string
    Date string
}

type Contact struct{
    Email string
    Phone string
    Message string
}
