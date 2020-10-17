package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"signed-url-service/aws/s3"
)

func main() {
	router := gin.Default()
	router.GET(":bucket/*key", func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		url := s3.GenerateSignedURL(context.Param("bucket"), context.Param("key"))
		context.String(http.StatusOK, url)
	})

	var port string
	flag.StringVar(&port, "port", "8081", "The port you want the gin-gonic server to expose.")
	flag.Parse()
	log.Println(port)

	router.Run(":"+port)

}