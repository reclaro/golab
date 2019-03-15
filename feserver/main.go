package main

import (
	"fmt"
	"github.com/hoisie/web"
)

func hello() string {
	return "hello"
}

func main() {
	web.Get("/(.*)", hello)
	//web.Run("0.0.0.0:9999")
	fmt.Println("Started")
}
