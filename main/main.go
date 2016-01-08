package main

import (
	"fmt"

	"me.car/network"
)

func main() {
	network.Listen()
	fmt.Println("The progress end")
	network.Close()
}
