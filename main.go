package main

import (
	"fmt"
	"log"

	c "github.com/WatchJani/Consistent-Hashing/consistent_hashing"
)

func main() {
	ch := c.NewConsistentHashing()
	if err := ch.Load("./config.json"); err != nil {
		log.Println(err)
	}

	s := ch.FindServer("janko")
	fmt.Println(s.Addr)
}
