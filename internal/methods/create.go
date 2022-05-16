package methods

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Escape-Technologies/http-request-catcher/api"
	"github.com/Escape-Technologies/http-request-catcher/pkg/database"
	"github.com/gin-gonic/gin"
)

var (
	DefaultMessage = "H@k3d!"
)

func formatHeaders(headers http.Header) map[string]string {
	ret := make(map[string]string)
	for k, v := range headers {
		ret[k] = strings.Join(v, ",")
	}

	return ret
}

func formatBucketEntry(c *gin.Context) *api.BucketEntry {
	body, err := c.GetRawData()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}

	return &api.BucketEntry{
		Id:       c.Param("request_id"),
		BucketId: c.Param("bucket_id"),
		Method:   c.Request.Method,
		Path:     c.Request.URL.Path,
		IP:       c.ClientIP(),
		Time:     time.Now().UTC().Format(time.RFC3339),
		Headers:  formatHeaders(c.Request.Header),
		Body:     string(body),
	}
}

func CreateBucketEntry(c *gin.Context) {
	requestData := formatBucketEntry(c)
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
