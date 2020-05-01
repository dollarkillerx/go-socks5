package main

import (
	"flag"
	"fmt"
	"log"
)


func main() {
	port := flag.Int("port",7010,"port for proxy")
	flag.Parse()

	server := &SocksServer{
		Version: 5,
	}

	if err := server.ListenAndServer(fmt.Sprintf(":%d",port));err != nil {
		log.Fatalln(err)
	}
}