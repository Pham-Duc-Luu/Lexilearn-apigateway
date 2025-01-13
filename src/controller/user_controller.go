package controller

import (
	"api-gateway/src/config"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	ServiceConfig config.ServiceConfig
}

func (uc *UserController) ProxyRequest(c *gin.Context) {
	// Construct the target URL
	targetURL := uc.ServiceConfig.BaseURL + c.Param("proxyPath")
	log.Println(targetURL)
	// Create a new HTTP request
	req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	// Copy headers and add the API key
	for k, v := range c.Request.Header {
		req.Header[k] = v
	}

	req.Header.Set("x-api-key", uc.ServiceConfig.APIKey)

	// Send the request to the target service
	client := &http.Client{Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reach service"})
		return
	}
	defer resp.Body.Close()

	// Copy the response back to the client
	c.Status(resp.StatusCode)
	for k, v := range resp.Header {
		c.Writer.Header()[k] = v
	}
	io.Copy(c.Writer, resp.Body)
}
