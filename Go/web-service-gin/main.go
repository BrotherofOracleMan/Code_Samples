package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID:"1", Title:"Blue Train",Artist: "John Coltrane",Price: 56.99},
	{ID:"2", Title: "Jeru", Artist:"Gerry Mulligan", Price: 17.99},
	{ID:"3", Title: "Sarah Vaughn and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func postAlbums(c * gin.Context){
	var newAlbum album
	//creation of newAlbum
	if err := c.BindJSON(&newAlbum); err != nil{
		return
	}
	albums = append(albums,newAlbum)
	c.IndentedJSON(http.StatusCreated,newAlbum)
}


func getAlbumByID(c *gin.Context){
	id := c.Param("id")
	for _,a := range albums {
		if a.ID == id{
			c.IndentedJSON(http.StatusOK,a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"album not found"})
}
*/

func DivideTwoOps(c *gin.Context) {
	print("Test")
	op1, err1 := strconv.Atoi(c.Query("op1"))
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(2)
	}
	op2, err2 := strconv.Atoi(c.Query("op2"))
	if op2 == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"val": "Cannot divide by zero"})
		return
	}
	if err2 != nil {
		fmt.Println(err1)
		os.Exit(2)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"val": float64(op1 / op2)})
}

func MultiplyTwoOps(c *gin.Context) {
	print("Test")
	op1, err1 := strconv.Atoi(c.Query("op1"))
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(2)
	}
	op2, err2 := strconv.Atoi(c.Query("op2"))
	if err2 != nil {
		fmt.Println(err1)
		os.Exit(2)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"val": op1 * op2})
}

func SubtractTwoOps(c *gin.Context) {
	print("Test")
	op1, err1 := strconv.Atoi(c.Query("op1"))
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(2)
	}
	op2, err2 := strconv.Atoi(c.Query("op2"))
	if err2 != nil {
		fmt.Println(err1)
		os.Exit(2)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"val": op1 - op2})
}

func addTwoOps(c *gin.Context) {
	print("Test")
	op1, err1 := strconv.Atoi(c.Query("op1"))
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(2)
	}
	op2, err2 := strconv.Atoi(c.Query("op2"))
	if err2 != nil {
		fmt.Println(err1)
		os.Exit(2)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"val": op1 + op2})
}

func main() {
	router := gin.Default()
	router.GET("/add", addTwoOps)
	router.GET("/subtract", SubtractTwoOps)
	router.GET("/multiply", MultiplyTwoOps)
	router.GET("/divide", DivideTwoOps)
	router.Run(":8080")
}
