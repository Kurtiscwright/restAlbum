package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// function to iterate through all albums searching for album with matching supplied id.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// loop through list of albums looking for matching id
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	// if no match is found return friendly album not found message
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
