package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "bookstore",
	Short: "A simple bookstore application",
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
