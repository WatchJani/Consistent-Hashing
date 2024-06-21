package main

import (
	"fmt"
	"log"
	c "root/consistent_hashing"
)

func main() {
	ch := c.NewConsistentHashing()
	if err := ch.Load("./config.json"); err != nil {
		log.Println(err)
	}

	s := ch.FindServer("key")
	fmt.Println(s.Addr)
}
