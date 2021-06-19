package main

import (
	"fmt"

	"gitlab.com/nextwavedevs/drop/drop-api/router"
)

func main() {

	fmt.Println("Starting drop ...")
	router.StartApi()

}
