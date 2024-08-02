package methods

import (
	"strings"

	"github.com/Escape-Technologies/http-request-catcher/internal/schema"
	"github.com/Escape-Technologies/http-request-catcher/pkg/database"
	"github.com/gin-gonic/gin"
)

const (
	DefaultMessage = "H@k3d!"
	HTMLMessage    = "<html><body><h1>H@k3d!</h1></body></html>"
)

// Create an entry for the current context.
func CreateBucketEntry(c *gin.Context) {
	requestData := schema.FormatBucketEntry(c)
	if requestData == nil {
		respondBasedOnAcceptHeader(c, 500, DefaultMessage)
		return
	}

	go database.StoreRequest(requestData)

	respondBasedOnAcceptHeader(c, 200, DefaultMessage)
}

// respondBasedOnAcceptHeader responds based on the Accept header
func respondBasedOnAcceptHeader(c *gin.Context, statusCode int, message string) {
	acceptHeader := c.GetHeader("Accept")

	if strings.Contains(acceptHeader, "application/json") {
		c.JSON(statusCode, gin.H{
			"message": message,
		})
	} else if strings.Contains(acceptHeader, "text/html") {
		c.Data(statusCode, "text/html", []byte(HTMLMessage))
	} else {
		// Default to JSON if no specific type is requested
		c.JSON(statusCode, gin.H{
			"message": message,
		})
	}
}
