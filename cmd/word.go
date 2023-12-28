package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"long/internal/word"
	"strings"
)

var desc = strings.Join([]string{
	"以下是支持的单词转换模式：",
	"1.全部转换为大写",
	"2.全部转换为小写",
	"3.下划线转换为大写开头",
	"4.下划线转换为小写开头",
	"5.驼峰转换为下划线",
}, "\n")
var str string
var mode int8

// 子命令word的设置
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderScoreToUpperCamelCase:
			content = word.UnderScoreToUpperCamelCase(str)
		case ModeUnderScoreToLowerCamelCase:
			content = word.UnderScoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderScore:
			content = word.CamelCaseToUnderScore(str)
		default:
			log.Fatalln("暂不支持这种单词转换格式,请 使用help mode 查看帮助文档")
		}
		log.Printf("输出结果为: %s", content)
	},
}

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderScoreToUpperCamelCase
	ModeUnderScoreToLowerCamelCase
	ModeCamelCaseToUnderScore
)

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词的内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词的转换格式")

}
