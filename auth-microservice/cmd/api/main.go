package main

import (
	internal "auth-microservice/internal"
	log "log"
)

func main() {
	// nil != nil?
	// no nil == nil
	// => no error
	if err := internal.Run(); err != nil {
		log.Fatal(err)
	}
}
