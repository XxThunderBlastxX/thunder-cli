package cmd

import (
	"github.com/XxThunderBlastxX/helpers"
	"github.com/spf13/cobra"
)

// jokesCmd represents the jokes command
var jokesCmd = &cobra.Command{
	Use:     "joke",
	Aliases: []string{"j", "jok"},
	Short:   "Get random dad jokes in your terminal",
	Long:    `Dadjoke CLI is a tool that gives you a random dad joke`,
	Run: func(cmd *cobra.Command, args []string) {
		helpers.GetRandomJokes()
	},
}

//init func is executed when joke command is given to cli
func init() {
	rootCmd.AddCommand(jokesCmd)
}
