package config

import (
    "github.com/BurntSushi/toml"
)

type Config struct {
    Server       server
    Applications map[string]application
}

type server struct {
    Port int
    Bind string `toml:"bind_address"`
}

type application struct {
    Name  string
    URI   string
    Root  string
    Steps []string `toml:"build_steps"`
}

func ReadInConfigFile(filename string) (Config, error) {
    var conf Config
    _, err := toml.DecodeFile(filename, &conf)

    return conf, err
}
