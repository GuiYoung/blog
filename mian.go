package main

import (
	"blog/models"
	"blog/router"
)

func main() {
	if err := models.InitDb(); err != nil {
		panic(err)
	}

	router.InitRouter()
}
