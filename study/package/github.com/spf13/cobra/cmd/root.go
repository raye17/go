package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git",
	Short: "git short description",
	Long:  "git long description ,git is a cmd lalala,lalala",
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("unrecognized command"))
	},
}

func Execute() {
	rootCmd.Execute()
}
