package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() (r *httprouter.Router) {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHandler)
	router.POST("/upload/:vid-id", uploadHandler)
	return router
}
func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8080", r)
}
