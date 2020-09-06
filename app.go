package main

import (
	"fmt"
	"log"

	"github.com/soma42/go-libchat-core/src/peer"
)

func main() {
	identity, error := peer.CreateIdentity(1)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(identity)
}
