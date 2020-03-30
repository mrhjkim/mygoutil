package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is ...
var RootCmd = &cobra.Command{
	Use: "mygoutil <flags> ...",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("test")
	},
}

// Execute run RootCmd.Execute
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
