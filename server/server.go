package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("[ Server Started at: http://localhost:8080 ]")
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8080", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	log.Println("---- Start the handler ----")
	defer log.Println("---- Finish the handler ----")

	// fork the request context
	ctx := req.Context()

	// a long process :)
	select {
	// our main code
	case <-time.After(time.Second * 5):
		res.Write([]byte("Hello Go!"))
	// if the client abort the loading, so just exit! don't wast the resources for it! 💛
	case <-ctx.Done():
		log.Fatalln("---- Server Time Out! ----")
		http.Error(res, ctx.Err().Error(), http.StatusInternalServerError)
	}
}
