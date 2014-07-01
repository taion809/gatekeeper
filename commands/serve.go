package commands

import (
    "github.com/spf13/cobra"
    "github.com/taion809/gatekeeper/builder"
    "github.com/taion809/gatekeeper/server"
)

var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "Startup the build server with your favorite protocol.",
    Long: `Startup the build server with your favorite protocol.
    This allows you to use http(s) based build requests... so cool, much fun.
    `,
}

func init() {
    serveCmd.Run = serve
}

func serve(cmd *cobra.Command, args []string) {
    conf := InitializeConfig()

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
