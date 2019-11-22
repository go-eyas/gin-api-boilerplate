package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
)

// Assets 静态资源服务器, prefix 路由前缀
func Assets(prefix string, box *packr.Box) gin.HandlerFunc {
	fileserver := http.FileServer(box)
	fileserver = http.StripPrefix(prefix, fileserver)
	return func(c *gin.Context) {
		p := strings.TrimPrefix(c.Request.URL.Path, prefix)
		if box.Has(p) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}

// AssetsNoRoute 没有匹配到任何路由时，将静态资源服务器的 index.html 返回，通常用于前端 History API 路由
func AssetsNoRoute(box *packr.Box) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if c.Request.Method == "GET" && c.Writer.Status() == 404 && strings.Index(c.Request.URL.Path, "/api/") != 0 {
				h, e := box.FindString("index.html")
				if e != nil {
					c.JSON(500, gin.H{
						"error": e.Error(),
					})
					return
				}
				c.Header("content-type", "text/html")
				c.String(200, h)
			}
		}()
		c.Next()
	}
}
