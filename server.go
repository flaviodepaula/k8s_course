package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/secret", Secret)
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

func Secret(w http.ResponseWriter, r *http.Request) {

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	fmt.Fprintf(w, "<h1>User: %s   Password: %s . secrets funcionando</h1>", user, password)
}

func Healthz(w http.ResponseWriter, r *http.Request) {

	duration := time.Since(startedAt)
	if duration.Seconds() > 25 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %s", fmt.Sprintf("%.2f", duration.Seconds()))))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}
