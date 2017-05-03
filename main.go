package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HelloWorld Struct
type HelloWorld struct {
	Greeting string `json:"greeting"`
}

func main() {
	const (
		defaultPort      = ":8080"
		defaultPortUsage = "default server port, ':8080'"
	)

	// flags
	port := flag.String("port", defaultPort, defaultPortUsage)

	flag.Parse()

	fmt.Printf("server will run on : %s\n", *port)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello/{name}", index).Methods("GET")
	log.Fatal(http.ListenAndServe(*port, router))
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Responsing to /hello request")
	log.Println(r.UserAgent())

	vars := mux.Vars(r)
	name := vars["name"]

	greeting := "Hello, " + name + "!"
	helloWorld := HelloWorld{Greeting: greeting}
	json.NewEncoder(w).Encode(helloWorld)
	w.WriteHeader(http.StatusOK)
}
