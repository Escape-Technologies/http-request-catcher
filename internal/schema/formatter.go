package schema

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Escape-Technologies/http-request-catcher/api"
	"github.com/gin-gonic/gin"
)

func _formatHeaders(headers http.Header) map[string]string {
	ret := make(map[string]string)
	for k, v := range headers {
		ret[k] = strings.Join(v, ",")
	}

	return ret
}

func FormatBucketEntry(c *gin.Context) *api.BucketEntry {
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
		Headers:  _formatHeaders(c.Request.Header),
		Body:     string(body),
	}
}
