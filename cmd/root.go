package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/viniciusbds/navio/src/logger"
)

var l = logger.New(time.Kitchen, true)

var (
	rootCmd = &cobra.Command{
		Use:          "navio",
		Short:        "|___/ Navio is an extremely simple runtime container",
		SilenceUsage: true,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("Root cmd")
		// },
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
