package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/parfaire/marvelapi/channels/marvel"
	charCont "github.com/parfaire/marvelapi/controllers/character_controller"
	"github.com/parfaire/marvelapi/util"
)

func main() {
	redisClient := InitResources()

	// initiate controllers with DI here
	characterController := charCont.New(redisClient, marvel.New())

	// set router
	router := httprouter.New()
	router.GET("/characters", characterController.Perform(characterController.GetAllIDs))
	router.GET("/characters/:id", characterController.Perform(characterController.GetById))

	address := fmt.Sprintf(":%v", util.MARVEL_API_PORT)
	Run(address, router)
}

func Run(address string, router *httprouter.Router) {
	http.ListenAndServe(address, router)
}
