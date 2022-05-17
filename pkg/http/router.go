package server

import (
	"fmt"

	"github.com/Escape-Technologies/http-request-catcher/internal/methods"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	return r
}

func SetupRouter(r *gin.Engine) {

	r.GET("/:bucket_id", methods.FetchBucket)
	r.DELETE("/:bucket_id", methods.RemoveBucket)

	r.Any("/:bucket_id/:request_id", methods.CreateBucketEntry)
}

func StartRouter(r *gin.Engine, port int) {

	formattedPort := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting server on port: %s\n", formattedPort)

	if err := r.Run(formattedPort); err != nil {
		panic(err)
	}
}
