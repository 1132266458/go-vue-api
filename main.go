package main

import (
	"go-vue-api/model"
	"go-vue-api/routes"
)

func main() {

	model.InitDb()

	routes.InitRouter()
}
