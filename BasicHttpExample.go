package main

import (
	serializers "./serializers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 3000

	http.HandleFunc("/", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := serializers.RootResponse{Message: "Hello World"}
	data, err := json.Marshal(response)

	if err != nil {
		panic("Bad Request")
	}

	fmt.Fprint(w, string(data))
}
