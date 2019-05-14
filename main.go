package main

import (
	"fmt"
	"github.com/HtLord/servapi/servapi"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main()  {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "Hello", "<h1>Hello world MTFK!</h1>")
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