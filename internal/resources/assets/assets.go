package assets

import (
	"crypto/md5" //#nosec
	"embed"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	FlareState "github.com/soulteary/flare/internal/state"
)

//go:embed favicon.ico
var Favicon embed.FS

//go:embed android-chrome-192x192.png
var Png192 embed.FS

//go:embed android-chrome-512x512.png
var Png512 embed.FS

//go:embed manifest.json
var Manifest embed.FS

func RegisterRouting(router *gin.Engine) {

	router.Use(optimizeResourceCacheTime())

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=31536000")
		c.FileFromFS("favicon.ico", http.FS(Favicon))
	})

	router.GET("/android-chrome-192x192.png", func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=31536000")
		c.FileFromFS("android-chrome-192x192.png", http.FS(Png192))
	})

	router.GET("/android-chrome-512x512.png", func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=31536000")
		c.FileFromFS("android-chrome-512x512.png", http.FS(Png512))
	})

	router.GET("/manifest.json", func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=31536000")
		c.FileFromFS("manifest.json", http.FS(Manifest))
	})

	if FlareState.AppFlags.DebugMode {
		router.StaticFS("/assets/css", http.Dir("embed/assets/css"))
		return
	}
}

// ViewHandler support dist handler from UI
// https://github.com/gin-gonic/gin/issues/1222
func optimizeResourceCacheTime() gin.HandlerFunc {
	data := []byte(time.Now().String())
	/* #nosec */
	etag := fmt.Sprintf("W/%x", md5.Sum(data))
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, "/assets/") ||
			strings.HasPrefix(c.Request.RequestURI, "/favicon.ico") {
			c.Header("Cache-Control", "public, max-age=31536000")
			c.Header("ETag", etag)

			if match := c.GetHeader("If-None-Match"); match != "" {
				if strings.Contains(match, etag) {
					c.Status(http.StatusNotModified)
					return
				}
			}
		}
	}
}
