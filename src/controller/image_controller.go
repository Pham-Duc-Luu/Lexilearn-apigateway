package controller

import (
	"api-gateway/src/config"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
	ServiceConfig config.ServiceConfig
}

func (ic *ImageController) ProxyRequest(c *gin.Context) {
	// Similar implementation as UserController
	targetURL := ic.ServiceConfig.BaseURL + c.Param("proxyPath")

	req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	for k, v := range c.Request.Header {
		req.Header[k] = v
	}
	req.Header.Set("x-api-key", ic.ServiceConfig.APIKey)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reach service"})
		return
	}
	defer resp.Body.Close()

	c.Status(resp.StatusCode)
	for k, v := range resp.Header {
		c.Writer.Header()[k] = v
	}
	io.Copy(c.Writer, resp.Body)
}
