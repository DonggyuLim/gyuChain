/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gyuChain/client"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Client Start before Set port",
	Long:  "Client is gyuChain client",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		port := args[0]
		client.Start(port)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
