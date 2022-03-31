package main

import (
	"blog/databases"
	"blog/router"
)

func main() {
	if err := databases.InitDb(); err != nil {
		panic(err)
	}

	router.InitRouter()
}
