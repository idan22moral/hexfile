/*
Copyright © 2022 Idan Moral idan22moral@gmail.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hexfile",
	Short: "Convert files to and from hexadecimal representation",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
