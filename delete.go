package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for index, a := range albums {
		if a.ID == id {
			albums = append(albums[:index], albums[index+1:]...)
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album delete failed"})
}
