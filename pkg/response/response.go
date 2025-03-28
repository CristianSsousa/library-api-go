package response

import "github.com/gin-gonic/gin"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ErrorResponse(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

func Success(c *gin.Context, statusCode int, data interface{}) {
	responseJSon := Response{
		Status:  statusCode,
		Message: "Success",
		Data:    data,
	}

	c.JSON(statusCode, responseJSon)
}

func ErrorValidation(c *gin.Context, statusCode int, errors []map[string]string) {
	c.JSON(statusCode, gin.H{"errors": errors})
}



