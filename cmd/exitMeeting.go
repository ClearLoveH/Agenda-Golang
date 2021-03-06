// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"Agenda-Golang/service"
	"fmt"
	"github.com/spf13/cobra"
)

// exitMeetingCmd represents the exitMeeting command
var exitMeetingCmd = &cobra.Command{
	Use:   "exitMeeting",
	Short: "Exit a meeting",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		infoLog.Println("Exit Meeting Called.")
		title,_ := cmd.Flags().GetString("title")
		currentUser,flag := service.GetCurrentUser()
		if flag == false{
			fmt.Println("Please Sign in firstly")
		} else {
			if exitMeeting := service.QuitMeeting(currentUser.GetName(),title); exitMeeting == true{
				fmt.Println("Quit meeting successfully!")
			} else {
				fmt.Println("Fail to quit meeting.")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(exitMeetingCmd)

	// Here you will define your flags and configuration settings.
	exitMeetingCmd.Flags().StringP("title", "t", "", "the title of meeting")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exitMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exitMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
