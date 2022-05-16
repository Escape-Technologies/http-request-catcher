package methods

import (
	"github.com/Escape-Technologies/http-request-catcher/pkg/database"
	"github.com/gin-gonic/gin"
)

func FetchBucket(c *gin.Context) {
	bucket := database.GetBucket(c.Param("bucket_id"))

	c.JSON(200, gin.H{
		"message": "Ok",
		"data":    bucket,
	})
}
