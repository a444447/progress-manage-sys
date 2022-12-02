package main

import (
	"progress-manage-system/model"
	"progress-manage-system/router"
)

func main() {
	model.InitDB()
	router.InitRouter()
}
