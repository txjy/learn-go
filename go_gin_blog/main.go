package main

import (
	"go_bin_blog/model"
	"go_bin_blog/routes"
)

func main() {

	model.InitDb()
	routes.InitRouter()
}
