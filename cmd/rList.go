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
	"context"
)

// rListCmd represents the rList command
var rListCmd = &cobra.Command{
	Use:   "rList",
	Short: "Lists releases of a given github repo",
	Long: `Provides a ordered list of the releases for a give github repository`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("rList called")
		client := github.NewClient(nil)
		opt := &github.ListOptions{}
		releases, _, err := client.Repositories.ListReleases(context.Background(), ghOrg, ghRepo, opt)
		if err != nil {
	   		fmt.Println(err)
   		}
		for id, release := range releases {
			preRelease := *release.Prerelease
			if ! preRelease {
				fmt.Printf("########## >>%d<<, Ver:%s, Draft:%t, Pre:%t\n", id, *release.Name, *release.Draft, *release.Prerelease)
				for aid, asset := range release.Assets {
					if strings.HasPrefix(*asset.Name, ghPrefix) {
						fmt.Printf("%d> %s (%s)\n", aid, *asset.Name, *asset.BrowserDownloadURL)
					}
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(rListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
