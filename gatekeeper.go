package main

import (
    "github.com/taion809/gatekeeper/builder"
    "github.com/taion809/gatekeeper/config"
    "github.com/taion809/gatekeeper/server"
)

func main() {
    conf, err := config.ReadInConfigFile("config.toml")

    if err != nil {
        panic(err)
    }

    var apps = make(map[string]*builder.Application, len(conf.Applications))

    for k, v := range conf.Applications {
        apps[k] = &builder.Application{
            Name:  v.Name,
            Root:  v.Root,
            Key:   v.Key,
            Steps: v.Steps,
        }
    }

    server := &server.Server{
        Port:         conf.Server.Port,
        Bind:         conf.Server.Bind,
        BuildURI:     conf.Server.BuildURI,
        Protocol:     conf.Server.Protocol,
        Cert:         conf.Server.Cert,
        Key:          conf.Server.Key,
        Applications: apps,
    }

    server.StartServer()
}
