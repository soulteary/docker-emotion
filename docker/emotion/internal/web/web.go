package web

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

//go:embed assets/favicon.png
var Favicon embed.FS

//go:embed index.html
var HomePage []byte

//go:embed assets
var Assets embed.FS

func API(pythonAPI string, port string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(optimizeResourceCacheTime([]string{"/favicon.png", "/assets/"}))

	r.Any("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", HomePage)
	})

	favicon, _ := fs.Sub(Favicon, "assets")
	r.Any("/favicon.png", func(c *gin.Context) {
		c.FileFromFS("favicon.png", http.FS(favicon))
	})

	static, _ := fs.Sub(Assets, "assets")
	r.StaticFS("/assets", http.FS(static))

	r.GET("/show-version.js", func(c *gin.Context) {
		version := `project <a href="https://github.com/soulteary/docker-emotion">soulteary/docker-emotion</a>, version 2022.09.30`
		c.Data(http.StatusOK, "application/javascript; charset=utf-8", []byte(`document.getElementById('version').innerHTML = '`+version+`';`))
		c.Abort()
	})

	r.Any("/api/*proxyPath", createProxy(pythonAPI))

	r.Run(":" + port)
}
