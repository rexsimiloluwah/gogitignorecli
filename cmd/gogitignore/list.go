package gogitignore

import (
	"fmt"
	"strings"

	"github.com/rexsimiloluwah/gogitignorecli/pkg/gogitignore"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Lists the gitignore input types",
	Run: func(cmd *cobra.Command, args []string) {
		char, _ := cmd.Flags().GetString("char")
		gitignoreInputTypes := gogitignore.FetchGitignoreList()
		if char == "" {
			fmt.Printf("All .gitignore input types\n\n")
			fmt.Println(strings.Join(gitignoreInputTypes, ", "))
			return
		}
		inputTypes := gogitignore.GetInputTypesStartsWith(gitignoreInputTypes, char)
		fmt.Printf("All .gitignore input types starting with '%s':\n\n", char)
		fmt.Println(strings.Join(inputTypes, ", "))
	},
}

func init() {
	listCmd.PersistentFlags().StringP("char", "c", "", "Lists all the gitignore input types starting with the passed character")

	rootCmd.AddCommand(listCmd)
}
