package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	en := gin.New()

	en.GET("/api/v1/message", func(c *gin.Context) {
		c.Writer.Write([]byte(""))
		c.JSON(http.StatusOK, Result{
			Code:    2000,
			Message: "success",
			Data:    "ok",
		})

	})

	en.Run("0.0.0.0:8081")
}

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
