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
	"github.com/hmsayem/jetstream-backup-restore-cli/client"
	"github.com/spf13/cobra"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore Data",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client.Restore(stream, auth)
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
	restoreCmd.Flags().StringVar(&stream.Name, "target", "", "Specify target stream")
	restoreCmd.MarkFlagRequired("target")
	restoreCmd.Flags().StringVar(&stream.Dir, "dir", "", "Specify backup directory")
	restoreCmd.MarkFlagRequired("dir")
	restoreCmd.Flags().BoolVar(&auth.TLS, "tls", true, "enables TLS")
	restoreCmd.Flags().StringVar(&auth.Cert, "tlscert", "", "specifies client certificate")
	restoreCmd.Flags().StringVar(&auth.Key, "tlskey", "", "specifies TLS key")
	restoreCmd.Flags().StringVar(&auth.CA, "tlsca", "", "specifies CA certificate")
	restoreCmd.Flags().StringVar(&auth.User, "user", "", "specifies a username")
	restoreCmd.Flags().StringVar(&auth.Password, "password", "", "specifies a password")
	restoreCmd.Flags().StringVar(&auth.Creds, "creds", "", "specifies credential file")
}
