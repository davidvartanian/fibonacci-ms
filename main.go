package main

import (
	"fibonacci-ms/pkg/fibonacci"
	"fibonacci-ms/pkg/fibonacci/endpoints"
	"net/http"
	"os"
	"strconv"
)

func main() {
	svc := fibonacci.NewService()
	http.HandleFunc("/get", endpoints.GetRequestHandler(svc))
	http.HandleFunc("/list", endpoints.ListRequestHandler(svc))
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	println("Listening to 0.0.0.0:" + strconv.Itoa(port))
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		println(err)
	}
}
