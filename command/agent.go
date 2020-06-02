package command

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/imdario/mergo"
	env "github.com/joeshaw/envdecode"
	"github.com/seashell/drago/agent"
	"github.com/seashell/drago/client"
	"github.com/seashell/drago/server"
	"github.com/spf13/cobra"

	"github.com/dimiro1/banner"
)

func NewAgentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agent",
		Short: "Runs a Drago agent",
		Long: `Usage: drago agent [options]
	
		Starts the Drago agent and runs until an interrupt is received.
		The agent may be a client and/or server.
	  
		The Drago agent's configuration primarily comes from the config
		files used, but a subset of the options may also be passed directly
		as CLI arguments.
	  `,
		Run: func(cmd *cobra.Command, args []string) {

			banner.Init(os.Stdout, true, true, strings.NewReader(Banner))

			fmt.Println("==> Starting drago agent...")

			// TODO: move this to prerun hook into root cmd, and pass config through context
			config := LoadConfigFromFile("./demo/server.hcl")
			mergo.Merge(&config, DefaultConfig)

			err := env.Decode(&config)

			agconfig := agent.Config{
				UI:      config.UI,
				DataDir: config.DataDir,
				Server: server.Config{
					Enabled: true,
				},
				Client: client.Config{
					Enabled: false,
				},
			}

			agent, err := agent.New(agconfig)
			if err != nil {
				fmt.Printf("==> Error starting Drago agent: %s\n", err)
			}

			agent.Run()

			var wait time.Duration

			c := make(chan os.Signal, 1)

			// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
			// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
			signal.Notify(c, os.Interrupt)

			// Block until we receive our signal.
			<-c

			// Create a deadline to wait for.
			_, cancel := context.WithTimeout(context.Background(), wait)
			defer cancel()

			os.Exit(0)
		},
	}

	cmd.Flags().BoolP("server", "s", false, "Start agent in server mode")
	cmd.Flags().BoolP("client", "c", true, "Start agent in client mode")
	cmd.Flags().BoolP("dev", "d", false, "Start agent in dev mode")

	return cmd
}