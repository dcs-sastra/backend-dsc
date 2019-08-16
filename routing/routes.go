package routing

import (
    "net/http"
    "backendDsc/handlers"
)

type Route struct{
    Name string
    Method string
    Pattern string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Get Date Time",
        "GET",
        "/",
        handlers.Index,
    },
    Route{
        "Get Image by ID",
        "GET",
        "/images/{id}",
        handlers.GetImage,
    },
    Route{
        "Get Team Info",
        "GET",
        "/members",
        handlers.GetMembers,
    },
    Route{
        "Get Member Info",
        "GET",
        "/members/{id}",
        handlers.GetMember,
    },
}
