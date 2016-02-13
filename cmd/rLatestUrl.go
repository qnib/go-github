// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/google/go-github/github"
)

// rLatestUrlCmd represents the rLatestUrl command
var rLatestUrlCmd = &cobra.Command{
	Use:   "rLatestUrl",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		rluClient := github.NewClient(nil)
		rluOpt := &github.ListOptions{}
		releases, _, err := rluClient.Repositories.ListReleases(ghOrg, ghRepo, rluOpt)
		if err != nil {
			fmt.Println(err)
		}
		for _, release := range releases {
			preRelease := *release.Prerelease
			draft := *release.Draft
			if ! (preRelease || draft) {
				for _, asset := range release.Assets {
					if strings.HasPrefix(*asset.Name, ghPrefix) {
						fmt.Println(*asset.BrowserDownloadURL)
					}
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(rLatestUrlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rLatestUrlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rLatestUrlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
