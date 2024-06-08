package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	podName := os.Getenv("POD_NAME")

	fmt.Fprintf(w, "<h1> Name: %s.    Age: %s.   PodeName:  %s. </h1>", name, age, podName)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {

	data, err := os.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatal("error reading file: ", err)
	}

	fmt.Fprintf(w, "<h1>My Family: %s </h1>", string(data))
}
