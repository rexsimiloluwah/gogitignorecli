package gogitignore

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gogitignore",
	Short: "A simple CLI for generating .gitignore files",
	Long:  `A simple CLI for generating .gitignore files for any language, operating system, or input type`,
}

func Execute() {
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
