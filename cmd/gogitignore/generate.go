package gogitignore

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/rexsimiloluwah/gogitignorecli/pkg/gogitignore"
	"github.com/rexsimiloluwah/gogitignorecli/pkg/utils"

	"github.com/spf13/cobra"
)

const (
	inputTypeArgName  = "input"
	outputDirArgName  = "outdir"
	gitignoreFileName = ".gitignore"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generates the .gitignore file for the passed input type",
	Run: func(cmd *cobra.Command, args []string) {
		inputType, _ := cmd.Flags().GetString(inputTypeArgName)
		outputDir, _ := cmd.Flags().GetString(outputDirArgName)

		gitignoreInputTypesList := gogitignore.FetchGitignoreList()
		var gitignoreFileContent string

		inputTypeSplit := strings.Split(inputType, ",")

		// check if the input types are valid
		for _, i := range inputTypeSplit {
			valid := gogitignore.CheckGitignoreInputTypeExists(i, gitignoreInputTypesList)
			if !valid {
				// check for the closest matching input types
				closestMatches := gogitignore.GetClosestInputTypeMatch(i, gitignoreInputTypesList)
				fmt.Printf("`%s` is not a valid gitignore input type\n\n", i)
				fmt.Printf("Did you mean: %s?\n", strings.Join(closestMatches, ", "))
				return
			}
		}

		if len(inputTypeSplit) == 1 {
			gitignoreFileContent = gogitignore.FetchGitignoreFileContent(inputType)
		} else {
			gitignoreFileContent = gogitignore.MergeGitignoreFileContent(inputTypeSplit...)
		}

		err := utils.WriteContentToFile(gitignoreFileName, outputDir, gitignoreFileContent)
		filePath := filepath.Join(outputDir, gitignoreFileName)
		if err != nil {
			fmt.Printf("an error occurred while writing to %s: %s", filePath, err.Error())
			return
		}
		fmt.Printf("\nYipee! .gitignore file successfully generated at: %s 🎉🎉\n", filePath)
	},
}

func init() {
	generateCmd.PersistentFlags().StringP(inputTypeArgName, "i", "", "Input type i.e. python")
	generateCmd.PersistentFlags().StringP(outputDirArgName, "o", "./", "Output directory")

	generateCmd.MarkPersistentFlagRequired("input")

	rootCmd.AddCommand(generateCmd)
}
