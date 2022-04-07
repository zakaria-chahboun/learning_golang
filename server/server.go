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
	http.HandleFunc("/2", handler2)
	http.ListenAndServe("localhost:8080", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	log.Println("---- Start the handler [1] ----")
	defer log.Println("---- Finish the handler [1] ----")

	// fork the request context
	ctx := req.Context()

	// a long process :)
	select {
	// our main code
	case <-time.After(time.Second * 5):
		res.Write([]byte("Hello Go! [1]"))
	// if the client abort the loading, so just exit! don't wast the resources for it! 💛
	case <-ctx.Done():
		fmt.Println("---- Server Time Out! [1] ----")
		// log.Fatalln("---- Server Time Out! [1] ----")
		// http.Error(res, ctx.Err().Error(), http.StatusInternalServerError)
	}
}

func handler2(res http.ResponseWriter, req *http.Request) {
	log.Println("---- Start the handler [2] ----")
	defer log.Println("---- Finish the handler [2] ----")

	// fork the request context
	ctx := req.Context()

	// a long process :)
	select {
	// our main code
	case <-time.After(time.Second * 4):
		res.Write([]byte("Hello Go! [2]"))
	// if the client abort the loading, so just exit! don't wast the resources for it! 💛
	case <-ctx.Done():
		fmt.Println("---- Server Time Out! [2] ----")
		//log.Fatalln("---- Server Time Out! [2] ----")
		//http.Error(res, ctx.Err().Error(), http.StatusInternalServerError)
	}
}
