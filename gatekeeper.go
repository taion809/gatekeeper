package main

import (
    "github.com/taion809/gatekeeper/config"
)

func main() {
    conf, err := config.ReadInConfigFile("config.toml")

    if err != nil {
        panic(err)
    }
}
