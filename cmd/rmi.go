package cmd

import (
	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/images"
)

func init() {
	rootCmd.AddCommand(remove())
}

func remove() *cobra.Command {
	return &cobra.Command{
		Use:   "rmi",
		Short: "Remove a Image",
		Long:  "ex: navio remove image <image_name> remove a downloaded images located in the ./images directory.",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) == 0 {
				l.Log("WARNING", "You must insert a image name!")
			} else {
				if args[0] != "" {
					images.DeleteImage(args[0])
				}
			}

			return nil
		},
	}
}