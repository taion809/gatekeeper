package server

import (
    "fmt"
    "github.com/gorilla/mux"
    "github.com/taion809/gatekeeper/builder"
    "github.com/taion809/gatekeeper/config"
    "net/http"
)

type Server struct {
    Conf *config.Config
}

func StartServer(server *Server) {
    router := mux.NewRouter()

    build_str := fmt.Sprintf("%s/{application:([a-zA-Z0-9])\\w+}", server.Conf.Server.BuildURI)
    router.HandleFunc(build_str, server.BuildHandler)

    http.Handle("/", router)
    conn_string := fmt.Sprintf("%s:%d", server.Conf.Server.Bind, server.Conf.Server.Port)

    if server.Conf.Server.Protocol == "http" {
        http.ListenAndServe(conn_string, nil)
    } else {
        http.ListenAndServeTLS(conn_string, server.Conf.Server.Cert, server.Conf.Server.Key, nil)
    }
}

func (s *Server) BuildHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    app := s.Conf.Applications[vars["application"]]

    // lol security?
    auth := vars["auth_key"]
    switch {
    case auth == nil && app.Key != nil:
        log.Fatal("Authentication key required")
    case auth != app.Key:
        log.Fatal("Authentication key mismatch")
    }

    fmt.Println("Kicking off build for ", app.Name)
    builder.StartBuild(&app)
}
