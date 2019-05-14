package main

import (
	"fmt"
	"github.com/HtLord/servapi/servapi"
	"log"
	"net/http"
	"os"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>My API will work on Heroku. But not sure about db connection.</h1>")
	})
	http.HandleFunc("/create/keeper", servapi.CreateKeeper)
	fmt.Println("Start serving")
	fmt.Println(os.Getenv("MONGOSECRET"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}