package main

import (
	"fmt"

	"github.com/nextwavedevs/drop/api"
	"github.com/nextwavedevs/drop/database"
)

func main(){
	client := database.DB()
	
	fmt.Println("Starting drop ...")

	api.StartApi()
}