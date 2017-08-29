package main

import (
	"fmt"
	"net/http"

	"github.com/FirusGuo/Golang-MVC/books"
	"github.com/julienschmidt/httprouter"
)

func main() {

	// NEW ROUTER
	router := httprouter.New()

	router.GET("/index", books.Index)

	// LOCAL SERVER
	fmt.Println("Listening on Local Server...")
	panic(http.ListenAndServe(":3000", router))

}
