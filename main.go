package main

import (
	"fmt"
	"github.com/HtLord/servapi/servapi"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main()  {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/create/keeper", servapi.GinCreateTest)
	fmt.Println("Start serving")
	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080


}

func notGin(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>My API will work on Heroku. But not sure about db connection.</h1>")
	})
	http.HandleFunc("/create/keeper", servapi.CreateKeeper)
}