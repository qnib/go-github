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
	//"strings"

	"github.com/spf13/cobra"
	"github.com/google/go-github/github"
	"context"
)

// tListCmd represents the tList command
var tListCmd = &cobra.Command{
	Use:   "tList",
	Short: "Lists tags of a given github repo",
	Long: `Provides a ordered list of the tags for a give github repository.

	Not really functional yet, since it does not provide a download url...`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		client := github.NewClient(nil)
		opt := &github.ListOptions{}
		tags, _, err := client.Repositories.ListTags(context.Background(), ghOrg, ghRepo, opt)
		if err != nil {
	   		fmt.Println(err)
   		}
		for id, tag := range tags {
			fmt.Printf("########## >>%d<<, Ver:%s\n", id, *tag.Name)
		}
	},
}

func init() {
	RootCmd.AddCommand(tListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
