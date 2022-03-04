package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql/testutil"
	"github.com/rajesh4295/graphql-server/gql"
)

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "server is running"})
	})

	router.StaticFS("/index", http.Dir("./public"))

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/index")
	})

	router.POST("/graphql", func(c *gin.Context) {
		var objmap map[string]interface{}
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
		err = json.Unmarshal(body, &objmap)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
		r := gql.ExecuteQuery(objmap["query"].(string), gql.Schema)
		c.JSON(http.StatusOK, r)
	})

	router.POST("/introspection", func(c *gin.Context) {
		r := gql.ExecuteQuery(testutil.IntrospectionQuery, gql.Schema)
		c.JSON(http.StatusOK, r)
	})

	router.Run(":9000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
