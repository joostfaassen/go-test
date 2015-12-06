package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	mux "github.com/gorilla/mux"
)

var greetingFlag string

func init() {
	flag.StringVar(&greetingFlag, "greeting", "Hello", "Type of greeting")
}

func main() {
	flag.Parse()
	fmt.Println(fmt.Sprintf("%s, gooo!?", greetingFlag))

	//http.HandleFunc("/", helloHandler)

	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler)
	r.HandleFunc("/hello/{name}", helloHandler)
	r.HandleFunc("/articles/{articleId}", helloHandler)
	http.Handle("/", r)

	http.ListenAndServe("127.0.0.1:8080", nil)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if strings.Compare(name, "") == 0 {
		name = "world"
	}
	w.Write([]byte(fmt.Sprintf("%s, %s!", greetingFlag, name)))
	log.WithFields(log.Fields{
		"name": name,
	}).Info("HelloHandler Request")

	//spew.Dump(r)

	/*
		b, err := json.Marshal(r)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(b))
		}
		w.Write(b)
	*/
}
