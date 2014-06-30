package main

import (
    "github.com/taion809/gatekeeper/config"
    "github.com/taion809/gatekeeper/server"
)

func main() {
    conf, err := config.ReadInConfigFile("config.toml")

    if err != nil {
        panic(err)
    }

    server.StartServer(&server.Server{Conf: conf})
}
