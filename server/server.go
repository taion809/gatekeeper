package server

import (
    "fmt"
    "github.com/gorilla/mux"
    "github.com/taion809/gatekeeper/builder"
    "log"
    "net/http"
)

type Server struct {
    Port         int
    Bind         string
    BuildURI     string
    Protocol     string
    Cert         string
    Key          string
    Applications map[string]*builder.Application
}

func (server *Server) StartServer() {
    router := mux.NewRouter()

    build_str := fmt.Sprintf("%s/{application:([a-zA-Z0-9])\\w+}", server.BuildURI)
    router.HandleFunc(build_str, server.BuildHandler)

    http.Handle("/", router)
    conn_string := fmt.Sprintf("%s:%d", server.Bind, server.Port)

    if server.Protocol == "http" {
        http.ListenAndServe(conn_string, nil)
    } else {
        http.ListenAndServeTLS(conn_string, server.Cert, server.Key, nil)
    }
}

func (s *Server) BuildHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var app = s.Applications[vars["application"]]

    // lol security?
    auth := r.FormValue("auth_key")
    switch {
    case auth == "" && app.Key != "":
        fmt.Printf("%v \n", auth)
        log.Fatal("Authentication key required")
    case auth != app.Key:
        log.Fatal("Authentication key mismatch")
    }

    fmt.Println("Kicking off build for ", app.Name)
    app.StartBuild()
}
