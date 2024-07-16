package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SlayerAnsh/vfs-resolver-api/gql"
	"github.com/SlayerAnsh/vfs-resolver-api/path"
	"github.com/gin-gonic/gin"
)

var App *gin.Engine

func init() {
	App = gin.Default()
	App.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
		return
	})
	App.GET("/protocol/*uri", func(c *gin.Context) {

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

func Handler(w http.ResponseWriter, r *http.Request) {
	App.ServeHTTP(w, r)
}
