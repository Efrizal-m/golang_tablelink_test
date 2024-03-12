package response

import "github.com/gin-gonic/gin"

func Error(c *gin.Context, httpCode int, err error) {
	c.JSON(httpCode, gin.H{
		"status":  false,
		"message": err.Error(),
	})
}

func Success(c *gin.Context, httpCode int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"status":  true,
		"message": "Successfully",
		"data":    data,
	})
}
