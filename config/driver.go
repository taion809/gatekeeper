package config

import (
    "github.com/BurntSushi/toml"
)

type Config struct {
    Server       server
    Applications map[string]Application
}

type server struct {
    Port     int
    Bind     string `toml:"bind_address"`
    BuildURI string `toml:"build_uri"`
    Protocol string
    Cert     string `toml:"ssl_cert_file,omitempty"`
    Key      string `toml:"ssl_key_file,omitempty"`
}

type Application struct {
    Name  string
    URI   string
    Root  string
    Key   string   `toml:"api_key"`
    Steps []string `toml:"build_steps"`
}

func ReadInConfigFile(filename string) (*Config, error) {
    var conf Config
    _, err := toml.DecodeFile(filename, &conf)

    return &conf, err
}
