package cmd

import (
	"github.com/spf13/cobra"
	"github.com/taylormonacelli/goodcow/embedded4"
)

// failCmd represents the fail command
var failCmd = &cobra.Command{
	Use:   "fail",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := embedded4.RunEmbedded()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(failCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// failCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// failCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
