package methods

import (
	"github.com/Escape-Technologies/http-request-catcher/internal/schema"
	"github.com/Escape-Technologies/http-request-catcher/pkg/database"
	"github.com/gin-gonic/gin"
)

const (
	DefaultMessage = "H@k3d!"
)

// Create an entry for the current context.
func CreateBucketEntry(c *gin.Context) {
	requestData := schema.FormatBucketEntry(c)
	if requestData == nil {
		c.JSON(500, gin.H{
			"message": DefaultMessage,
		})

		return
	}

	go database.StoreRequest(requestData)

	c.JSON(200, gin.H{
		"message": DefaultMessage,
	})
}
