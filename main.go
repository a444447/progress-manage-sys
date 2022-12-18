package main

import (
	"fmt"
	"progress-manage-system/model"
	"progress-manage-system/router"
)

func main() {
	if _, err := model.InitDB(); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("model.Db:%v\n", model.Db)

	router.InitRouter()
}
