/**
 * Copyright 2020 Orcha Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
