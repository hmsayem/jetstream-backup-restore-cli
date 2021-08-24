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

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup Data",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client.Backup(stream, auth)
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.Flags().StringVar(&stream.Name, "target", "", "specifies target stream")
	backupCmd.MarkFlagRequired("target")
	backupCmd.Flags().StringVar(&stream.Dir, "dir", "", "specifies backup directory")
	backupCmd.MarkFlagRequired("dir")
	backupCmd.Flags().BoolVar(&auth.TLS, "tls", true, "enables TLS")
	backupCmd.Flags().StringVar(&auth.Cert, "tlscert", "", "specifies client certificate")
	backupCmd.Flags().StringVar(&auth.Key, "tlskey", "", "specifies TLS key")
	backupCmd.Flags().StringVar(&auth.CA, "tlsca", "", "specifies CA certificate")
	backupCmd.Flags().StringVar(&auth.User, "user", "", "specifies a username")
	backupCmd.Flags().StringVar(&auth.Password, "password", "", "specifies a password")
	backupCmd.Flags().StringVar(&auth.Creds, "creds", "", "specifies credential file")
}
