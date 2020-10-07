
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

var(
	resource string
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create model",
	Long: `Flag create for generate model`,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Create(resource)
		defer f.Close()

		if err != nil {
			color.Red(fmt.Sprintf("Error: %v", err))
		}

		color.Cyan(fmt.Sprintf("Success for create resource %s", resource))

	},
}

func init() {
	createCmd.Flags().StringVarP(&resource, "resource", "r", "", "create resource full")
	rootCmd.AddCommand(createCmd)
}
