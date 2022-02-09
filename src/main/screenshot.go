package main

import (
	"go_screenshot/src/middleware"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-vgo/robotgo"
)

var owo float32

func main() {
	ip := ":8080"
	// 初始化屏幕缩放率
	awa, _ := robotgo.GetScaleSize()
	qwq, _ := robotgo.GetScreenSize()
	owo = float32(awa) / float32(qwq)

	if len(os.Args) > 1 {
		ip = os.Args[1]
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.StaticFile("/", "./resources/index.html")
	r.StaticFile("/pc.html", "./resources/pc.html")
	r.StaticFile("/mobile.html", "./resources/mobile.html")
	// 	r.GET("/", func(c *gin.Context) {
	// 		c.Header("Content-Type", "text/html; charset=utf-8")
	// 		c.String(200, `<!DOCTYPE html>
	// <html lang="cn">
	//   <head>
	//     <meta charset="UTF-8" />
	//     <meta http-equiv="X-UA-Compatible" content="IE=edge" />
	//     <meta name="viewport" content="width=device-width, initial-scale=1.0" />
	//     <script src="https://cdn.bootcss.com/jquery/3.4.1/jquery.js"></script>
	//     <!-- <meta content="width=device-width, initial-scale=0.33, maximum-scale=1.0, user-scalable=0;" name="viewport" /> -->
	//     <title>Document</title>
	//   </head>
	//   <body>
	//     <div id="pw_body" style="width: 100%; height: 100%">
	//       <canvas
	//         id="canvas"
	//         style="position: absolute; left: 0px; top: 0px"
	//         width="1080"
	//         height="1920"
	//       ></canvas>
	//     </div>
	//     <script language="JavaScript">
	//       var canvas = document.getElementById("canvas");
	//       var ctx = canvas.getContext("2d");
	//       var img = new Image();
	//       img.src = "./screenshot?temp=" + Math.random();
	//       img.onload = () => {
	//         ctx.translate(1080, 0);
	//         ctx.rotate(Math.PI / -2);
	//         ctx.drawImage(img, -1920, -1080, 1920, 1080);
	//       };

	//       $("#pw_body").click(() => {
	//         img.src = "./screenshot?temp=" + Math.random();
	//         img.onload = () => {
	//           ctx.drawImage(img, -1920, -1080, 1920, 1080);
	//         };
	//       });
	//     </script>
	//   </body>
	// </html>
	// `)
	// 	})
	r.GET("screenshot/", func(c *gin.Context) {
		if c.Query("type") == "click" {
			clickX, _ := strconv.Atoi(c.Query("clickX"))
			clickY, err := strconv.Atoi(c.Query("clickY"))
			if err == nil {
				log.Printf("单击(%v,%v)\n", zoom(clickX), zoom(clickY))
				// robotgo.MoveMouseSmooth(zoom(clickX), zoom(clickY), 0.1, 0.1)
				robotgo.MoveClick(zoom(clickX), zoom(clickY), `left`, false)
			}
		}
		if c.Query("type") == "dbclick" {
			clickX, _ := strconv.Atoi(c.Query("clickX"))
			clickY, err := strconv.Atoi(c.Query("clickY"))
			if err == nil {
				log.Printf("双击(%v,%v)\n", zoom(clickX), zoom(clickY))
				// robotgo.MoveMouseSmooth(zoom(clickX), zoom(clickY), 0.1, 0.1)
				robotgo.MoveClick(zoom(clickX), zoom(clickY), `left`, true)
			}
		}
		if c.Query("type") == "rclick" {
			clickX, _ := strconv.Atoi(c.Query("clickX"))
			clickY, err := strconv.Atoi(c.Query("clickY"))
			if err == nil {
				log.Printf("右击(%v,%v)\n", zoom(clickX), zoom(clickY))
				// robotgo.MoveMouseSmooth(zoom(clickX), zoom(clickY), 0.1, 0.1)
				robotgo.MoveClick(zoom(clickX), zoom(clickY), `right`, false)
			}
		}

		middleware.ImageToByte(c.Writer, middleware.ScreenShotRGBA()[0])
	})
	r.GET("screenshot/:page", func(c *gin.Context) {
		page, err := strconv.Atoi(c.Param("page"))
		if err != nil || page < 0 {
			page = 0
		}
		middleware.ImageToByte(c.Writer, middleware.ScreenShotRGBA()[page])
	})
	r.GET("getinfo", func(c *gin.Context) {
		var x, y = robotgo.GetScaleSize()
		c.JSON(http.StatusOK, gin.H{"Resolution": strconv.Itoa(x) + "," + strconv.Itoa(y)})
	})
	r.Run(ip)
}

func zoom(z int) int {
	return int(float32(z) / owo)
}
