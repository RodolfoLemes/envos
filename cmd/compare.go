/*
Copyright © 2023 Rodolfo Lemes rodolfo_fero@hotmail.com
*/
package cmd

import (
	"github.com/RodolfoLemes/envos/internal"

	"github.com/spf13/cobra"
)

// compareCmd represents the compare command
var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "compare and change your env to catch mismatch",
	Long:  `This command will compare the env to a .go file, search all occurrences of get envs in this go file, open the .env and compare one another. Those missing envs on .env will be created empty`,
	RunE:  runCommand,
}

var (
	filename string
	filepath string
)

func init() {
	rootCmd.AddCommand(compareCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	compareCmd.PersistentFlags().StringVarP(&filename, "filename", "n", ".env", "env filename")
	compareCmd.PersistentFlags().StringVarP(&filepath, "filepath", "p", "config.go", "config.go file where there are all the envs calls")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runCommand(cmd *cobra.Command, args []string) error {
	return internal.Compare(filename, filepath)
}
