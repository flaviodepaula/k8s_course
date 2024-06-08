package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	podName := os.Getenv("POD_NAME")

	fmt.Fprintf(w, "<h1> Name: %s.    Age: %s.   PodeName:  %s. </h1>", name, age, podName)
}
