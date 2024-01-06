package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"long/internal/art"
)

var (
	path  string
	scale float64 // 缩放因子，调整输出的 ASCII 图像大小
)

var artCmd = &cobra.Command{
	Use:   "art ",
	Short: "image_ascii_converter",
	Long:  "convert the image  into ascii art , and you should provide the path of file",
	Run: func(cmd *cobra.Command, args []string) {
		if path == "" {
			fmt.Println("you should enter the path of the image")
			return
		}

		art.Convert(path, scale)
	},
}

func init() {
	artCmd.Flags().StringVarP(&path, "path", "p", "", "the path of the image you want to convert")
	artCmd.Flags().Float64VarP(&scale, "scale", "s", 0.3, "the scale of the ASCII art ,the default is 0.3,and the range is from 0 to 1")

}
