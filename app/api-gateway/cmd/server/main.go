package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func main() {
	r := gin.Default()
	r.POST("/user/register", ReverseProxy("user-edge-service:8000", "/v1/register"))
	r.POST("/user/login", ReverseProxy("user-edge-service:8000", "/v1/login"))
	r.POST("/auth/code", ReverseProxy("user-edge-service:8000", "/v1/auth/code"))
	r.POST("/auth/verify", ReverseProxy("user-edge-service:8000", "/v1/auth/verify"))
	r.GET("/course", ReverseProxy("course-edge-service:8001", "/v1/course"))
	r.Run(":80")
}

func ReverseProxy(target string, path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = target
			req.Host = target
			req.URL.Path = path
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
