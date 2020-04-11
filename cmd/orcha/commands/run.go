package commands

import (
	"github.com/aagea/orcha/internal/app/orcha"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Launch the Orcha service",
	Long:  `Launch the service`,
	Run: func(cmd *cobra.Command, args []string) {
		s := orcha.NewService(orchaConfig)
		s.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
