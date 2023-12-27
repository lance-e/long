package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"simpleCliApplication/internal/translate"
)

var (
	query string
	to    string
)

var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "translate the word",
	Long:  "translate the word you had entered ,and return the result ",
	Run: func(cmd *cobra.Command, args []string) {
		if query == "" || to == "" {
			fmt.Println("not enough params")
			return
		}
		translate.Translate(query, to)

	},
}

func init() {
	translateCmd.Flags().StringVarP(&query, "word", "w", "hello", "the word you want to translate")
	translateCmd.Flags().StringVarP(&to, "to", "t", "zh", "the type of which you want to translate")
}
