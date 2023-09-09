package console

import (
	"fmt"
	"os"

	"github.com/fajarachmadyusup13/product-management/internal/config"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "",
	Short: "Base Command",
	Long:  "Base Command",
}

// Execute :nodoc:
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	config.GetConf()
}
