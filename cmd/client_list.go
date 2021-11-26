package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gostandup/persist"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all clients",
	Long: `Lists all clients. 
Results can be filtered`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cmd.Usage()
			return
		}
		listCommand(cmd, args)
	},
}

func init() {
	clientCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listCommand(cmd *cobra.Command, args []string) {
	c, err := persist.ReadClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	maxLen := 0
	for _, l := range c {
		maxLen = max(maxLen, len(l.Key))
	}

	fmt.Println("List of all clients")
	fmt.Printf("\n%10s %s\n", "Key", "Name")
	fmt.Printf("%10s %s\n", "===", "====")
	for _, l := range c {
		fmt.Printf("%10s %s\n", l.Key, l.Name)
	}
	fmt.Println()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
