package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Client related actions",
	Long:  "Allows you to add, update, remove and list client information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Invalid parameters specified")
		cmd.Help()
	},
}

func init() {
	AddToRoot(clientCmd)
}
