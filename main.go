package main

import (
	"fmt"
	"log"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Unable to parse request err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST Request succesfull\n")
	encData := r.FormValue("encData")
	fmt.Fprintf(w, "Got enc data as %s\n", encData)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Unable to parse request err: %v", err)
		return
	}

	fmt.Fprintf(w, "/get API called succesfull\n")

	encData := r.FormValue("key")
	fmt.Fprintf(w, "Got key as %s\n", encData)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/get", getHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
