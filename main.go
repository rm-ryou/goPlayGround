package main

import (
	"log"

	"github.com/rm-ryou/goPlayGround/cmd"
)

func main() {
	objectList := cmd.CreateObjectWorld()
	data := cmd.CreateData(objectList)

	err := cmd.OutputToFile("dist/result.ppm", data)
	if err != nil {
		log.Fatal(err)
	}
}
