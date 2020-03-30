package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "mygoutil <flags> ...",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("test")
	},
}
