package main

import (
	"awesomeProject2/routes"
	"awesomeProject2/utils"
	"flag"
)

var (
	addr = flag.String("addr", ":8000", "TCP address to listen to")
)

func main() {
	flag.Parse()
	utils.InitLogger()
	routes.Routes(addr)
}
