package middleware

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"runtime"

	screenshot "github.com/kbinani/screenshot"
)

// 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// 妮可结构体
type Neko struct {
}

// 截图响应结构体
type NekoSreenshotResponse struct {
	QWQ  string
	Imgs []image.RGBA
}

// 妮可可以截图
func (*Neko) Screenshot(req ArithRequest, res *NekoSreenshotResponse) error {
	res.Imgs = ScreenShotRGBA()
	res.QWQ = "妮可截完图了qwq"
	return nil
}

// 获取不同操作系统的环境的截图临时文件夹
func getScreenshotFilename() string {
	var (
		filepath string
	)
	if runtime.GOOS == "windows" {
		filepath = os.Getenv("systemdrive") + "\\ProgramData\\"
	} else {
		filepath = "/tmp/"
	}
	return filepath
}

// 存储图片到本地临时文件夹
func SaveImages(imgs []image.RGBA) {
	for i := 0; i < len(imgs); i++ {
		file, _ := os.Create(fmt.Sprintf("%s/%d.png",getScreenshotFilename(),i))
		defer file.Close()
		png.Encode(file, &imgs[i])
	}
}

// 截图,返回图片切片
func ScreenShotRGBA() []image.RGBA {
	n := screenshot.NumActiveDisplays()
	ls := make([]image.RGBA, n)
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			// connect()
			fmt.Println(err)
		}
		ls[i] = *img
	}
	return ls
}

func ImageToByte(wer io.Writer,img image.RGBA){
	png.Encode(wer,&img)
}
