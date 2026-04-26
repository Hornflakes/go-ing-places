package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Released int16  `json:"released"`
}

// albums slice to see record album data.
var albums = []album{
	{ID: 1, Title: "Going Under", Artist: "Evanescence", Released: 2003},
	{ID: 2, Title: "Now You re Gone", Artist: "Basshunter", Released: 2007},
	{ID: 3, Title: "GO BABY", Artist: "Justin Bieber", Released: 2025},
	{ID: 4, Title: "Gone with the Sin", Artist: "HIM", Released: 1999},
	{ID: 5, Title: "Iris", Artist: "Goo Goo Dolls", Released: 1998},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a reponse.
func getAlbumByID(c *gin.Context) {
	idParam := c.Param("id")

	// Parse string to int64 (base 10, 64-bit)
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a number: " + err.Error()})
		return
	}

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
