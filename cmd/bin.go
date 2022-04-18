/*
Copyright © 2022 Idan Moral idan22moral@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/idan22moral/hexfile/convert"
	"github.com/spf13/cobra"
)

var binCmd = &cobra.Command{
	Use:   "bin",
	Short: "A brief description of your command",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("input file must be specified")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		outputFilename, err := cmd.Flags().GetString("output")

		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		data, err := os.ReadFile(args[0])

		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		err2 := convert.HexToFile(string(data), outputFilename, 0666)

		if err2 != nil {
			cmd.PrintErrln(err2)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(binCmd)

	binCmd.Flags().StringP("output", "o", "", "Name of the file to write to")
}
