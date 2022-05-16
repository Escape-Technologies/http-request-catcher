package methods

import (
	"github.com/Escape-Technologies/http-request-catcher/pkg/database"
	"github.com/gin-gonic/gin"
)

func RemoveBucket(c *gin.Context) {
	count := database.DeleteBucket(c.Param("bucket_id"))

	c.JSON(200, gin.H{
		"message": "Ok",
		"data":    count,
	})
}
