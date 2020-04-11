package commands

import (
	"fmt"
	"os"

	"github.com/aagea/orcha/internal/app/orcha"
	"github.com/spf13/cobra"
)

var orchaConfig orcha.Config

var rootCmd = &cobra.Command{
	Use:     "orcha",
	Short:   "Orcha service",
	Long:    `Orcha service`,
	Version: "NaN",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute the user command
func Execute(version string, commit string) {
	versionTemplate := fmt.Sprintf("%s [%s] ", version, commit)
	rootCmd.SetVersionTemplate(versionTemplate)
	orchaConfig.Version = version
	orchaConfig.Commit = commit
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
