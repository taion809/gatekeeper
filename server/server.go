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

func StartServer(conf *config.Config) {
    s := Server{Conf: conf}

    router := mux.NewRouter()

    build_str := fmt.Sprintf("%s/{application:([a-zA-Z0-9])\\w+}", conf.Server.BuildURI)
    router.HandleFunc(build_str, s.BuildHandler)

    http.Handle("/", router)

    conn_string := fmt.Sprintf("%s:%d", conf.Server.Bind, conf.Server.Port)
    http.ListenAndServe(conn_string, nil)
}

func (s *Server) BuildHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    app := s.Conf.Applications[vars["application"]]

    fmt.Println("Kicking off build for ", app.Name)
    builder.StartBuild(&app)
}
