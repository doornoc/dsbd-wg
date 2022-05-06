package api

import (
	"github.com/doornoc/dsbd-wg/pkg/core/peer"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RestAPI() error {
	router := gin.Default()
	router.Use(cors)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/peer", peer.Add)
			v1.GET("/peer", peer.Get)
			v1.DELETE("/peer", peer.Delete)
		}
	}

	log.Fatal(http.ListenAndServe(":8080", router))

	return nil
}

func cors(c *gin.Context) {
	//c.Header("Access-Control-Allow-Headers", "Accept, Content-ID, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-ID", "application/json")
	c.Header("Access-Control-Allow-Credentials", "true")
	//c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
