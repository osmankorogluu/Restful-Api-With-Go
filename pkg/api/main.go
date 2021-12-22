package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	Date   float64 `json:"date"`
}

var albums = []album{
	{ID: "1", Title: "Depresyon Oteli", Artist: "Norm Ender", Price: 12.99, Date: 2010},
	{ID: "2", Title: "Dünyanın Sonunda Doğmuşum", Artist: "maNga", Price: 10, Date: 2009},
	{ID: "3", Title: "Lose Yourself", Artist: "Eminem", Price: 8, Date: 2002},
	{ID: "4", Title: "Bu Son Olsun", Artist: "Cem Karaca", Price: 9.99, Date: 1990},
	{ID: "5", Title: "Urfalıyım Ezelden", Artist: "Kazancı Bedih", Price: 20, Date: 2014},
	{ID: "6", Title: "Anıları Sakla", Artist: "Batuhan Kordel", Price: 4.99, Date: 2021},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
