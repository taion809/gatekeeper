package commands

import (
    "github.com/spf13/cobra"
    "github.com/taion809/gatekeeper/config"
    "log"
)

var initCommand = &cobra.Command{Use: "gatekeeper"}
var confFile string

func AddCommands() {
    initCommand.AddCommand(serveCmd)
}

func Execute() {
    AddCommands()
    initCommand.Execute()
}

func InitializeConfig() *config.Config {
    conf, err := config.ReadInConfigFile(confFile)

    if err != nil {
        log.Fatal(err)
    }

    return conf
}

func init() {
    initCommand.PersistentFlags().StringVar(&confFile, "config", "config.toml", "config file (default is path/config.toml)")
}
