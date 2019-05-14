package main

import (
	"fmt"
	"github.com/HtLord/servapi/servapi"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/create/keeper", servapi.CreateKeeper)
	fmt.Println("Start serving")
	log.Fatal(http.ListenAndServe(":80", nil))
}