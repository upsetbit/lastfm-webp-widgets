package webputil

import (
	"fmt"
	"log"

	"github.com/chai2010/webp"
)

func DoSomething(data []byte) {
	var width, height int
	var err error

	if width, height, _, err = webp.GetInfo(data); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("width = %d, height = %d\n", width, height)
}
