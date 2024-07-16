package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SlayerAnsh/vfs-resolver-api/gql"
	"github.com/SlayerAnsh/vfs-resolver-api/path"
	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func init() {
	app = gin.Default()
	app.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
		return
	})
	app.GET("/protocol/*uri", func(c *gin.Context) {

		host := c.Request.Host
		chainId := strings.Split(strings.Split(host, ":")[0], ".")[0]

		chain, err := gql.GetChain(chainId)

		if err != nil {
			chainId = "galileo-4"
			chain, err = gql.GetChain(chainId)
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		vfs, err := gql.GetVfs(fmt.Sprint(chain["kernelAddress"]))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		uri := c.Param("uri")
		resolvedPath, err := path.ResolvePath(uri, vfs)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, resolvedPath)
		return
	})
}

func main() {
	app.Run(":8080")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
