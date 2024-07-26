package main

import (
	"log"

	"github.com/rm-ryou/goPlayGround/cmd"
)

func main() {
	data := cmd.CreateData(256, 256)
	err := cmd.OutputToFile("dist/result.ppm", data)
	if err != nil {
		log.Fatal(err)
	}
}
