/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"github.com/spf13/cobra"
)

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "to log in a pic-bed's server ",
	Run: func(cmd *cobra.Command, args []string) {
		url := "http://localhost:8080/api/auth/login"
		name, password := args[0], args[1]
		fmt.Println(name, password)
		resp, err := HttpRequest.Post(url,
			map[string]interface{}{
				"username": name,
				"password": password,
			})
		if err != nil {
			panic(err)
		}
		defer func(resp *HttpRequest.Response) {
			err := resp.Close()
			if err != nil {

			}
		}(resp)

		body, err := resp.Body()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))

	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
}
