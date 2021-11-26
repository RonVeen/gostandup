package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gostandup/model"
	"gostandup/persist"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use: `create <Key> <Name>
	<Key> unique Key
	<Name> client Name`,
	Args:  cobra.ExactArgs(2),
	Short: "Creates a new client",
	Long: `Creates a new client,

gostandup client create <Key> <Name>
	<Key> unique Key for the new client
	<Name> for the new client

Example: gostandup client create KEY01 ACME`,
	Run: func(cmd *cobra.Command, args []string) {
		createClient(*cmd, args)
	},
}

func init() {
	clientCmd.AddCommand(createCmd)
}

func createClient(cmd cobra.Command, args []string) {
	if len(args) != 2 {
		cmd.Usage()
		return
	}
	key := args[0]
	name := args[1]

	c := model.Client{
		Key:  key,
		Name: name,
	}

	success, err := persist.AppendClient(c)
	if success {
		fmt.Printf("Created Client as %s.	", c.String())
	} else {
		fmt.Println(err.Error())
	}
}
