package art

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

const (
	asciiChars = "@B%8WM#*oahkbdpwmZO0QCJYXzcvnxrjft/|()1{}[]-_+~<>i!lI;:,^`. "
)

func Convert(path string, scale float64) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("couldn't open this file")
		return
	}
	defer file.Close()
	data, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("couldn't decode this image")
		return
	}
	//get the size of ASCII art
	width := int(float64(data.Bounds().Dx()) * scale)
	height := int(float64(data.Bounds().Dy()) * scale)

	//resize

	newImage := image.NewRGBA(image.Rect(0, 0, width, height)) //新建一个rgba图像

	//使用双线性插值法，改变原图片尺寸
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dx := int(float64(x) / scale)
			dy := int(float64(y) / scale)
			newImage.Set(x, y, data.At(dx, dy))
		}
	}
	// 最后的ascii art
	ASCII := ""
	//开始转换为ascii图像
	bounds := newImage.Bounds()
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			gray := getPixelGray(newImage.At(x, y))
			charIndex := int(float64(gray) / 255 * float64(len(asciiChars)-1))

			ASCII += string(asciiChars[charIndex])
		}
		ASCII += "\n" // 每行结束加一个换行符
	}
	fmt.Println(ASCII) //将图像打印出来
}

// getPixelGray 获取像素的灰度值 ！！！
func getPixelGray(pixel color.Color) uint8 {
	r, g, b, _ := pixel.RGBA()
	gray := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b) //灰度值公式
	return uint8(gray / 256)                                       //因为
}
