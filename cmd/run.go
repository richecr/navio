package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/container"
)

var (
	// Used for containerame flag.
	containerName string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&containerName, "name", "", "The name of the container")
	rootCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(createContainer())
}

func createContainer() *cobra.Command {
	return &cobra.Command{
		Use: "run",
		RunE: func(cmd *cobra.Command, args []string) error {
			// navio run IMAGE COMMAND PARAMS...
			image := args[0]
			command := args[1]
			params := args[2:]

			if containerName == "" {
				// TODO: generate a random container name
				containerName = "XPTO"
			}

			l.Log("INFO", fmt.Sprintf("Image: %s, Command: %s, Params: %v", image, command, params))
			args = append([]string{"run", image, command, containerName}, params...)
			container.CreateContainer(args)

			return nil
		},
	}
}