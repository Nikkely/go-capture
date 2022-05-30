package main

import (
	"log"
	"os"

	"github.com/Nikkely/go-capture/capturer"
)

func main() {
	if err := capturer.Run("./conf.yaml"); err != nil {
		log.Fatalln(err)
	}
	os.Exit(0)
}
