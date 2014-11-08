package main

import (
	"flag"
	"github.com/SlyMarbo/spdy"
	"log"
	"net/http"
)

func main() {
	cert := flag.String("cert", "server.crt", "ssl certificate")
	key := flag.String("key", "server.key", "ssl key")

	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir("./public")))
	var err error
	go func() {
		log.Print("Launching SPDY on :44300...")
		err = spdy.ListenAndServeTLS(":44300", *cert, *key, nil)
	}()
	go func() {
		log.Print("Launching No SPDY on :44301...")
		err = http.ListenAndServeTLS(":44301", *cert, *key, nil)
	}()
	select {}
}
