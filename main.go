package main

import (
	"embed"
	"io/fs"
	"math/rand"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed web/build
var build embed.FS

// OpenWeb 调用系统浏览器
func OpenWeb(url string) {
	exec.Command(`cmd`, `/c`, `start`, `http://`+url).Start()
}

// SPAIndex 载入静态文件
func SPAIndex() http.FileSystem {
	fsys := fs.FS(build)
	buildStatic, _ := fs.Sub(fsys, "web/build")
	return http.FS(buildStatic)
}

// StaticHand 静态文件
func StaticHand() gin.HandlerFunc {
	return func(c *gin.Context) {
		upath := c.Request.URL.Path
		if !strings.HasPrefix(upath, "/api") {
			content := SPAIndex()
			http.FileServer(content).ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}

// SeedAPI 获取一个随机种子
func SeedAPI(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	c.JSON(http.StatusOK, gin.H{
		"status":  1,
		"message": "This is a seed",
		"data":    rand.Float64(),
	})
}

func main() {
	listenAddr := "127.0.0.1:8080"
	OpenWeb(listenAddr)
	// 允许跨域访问
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "OPTION"}
	// 初始化
	router := gin.Default()
	router.Use(cors.New(config))

	router.Use(StaticHand())
	api := router.Group("/api")
	{
		api.GET("/seed", SeedAPI)
	}

	router.Run(listenAddr)
}
