package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	//r.HandleFunc("/products/{id}", ProductsHandler)
	s := r.PathPrefix("/news").Subrouter()
	// "news/"

	s.HandleFunc("/", NewsHandler)
	// "/products/{key}/"
	s.HandleFunc("/{id}/", NewsHandler)
	// "/products/{key}/details"
	s.HandleFunc("/health", NewsHealthHandler)
	http.ListenAndServe(":8083", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: Home %v\n", vars["category"])
}

func NewsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids := vars["id"]
	// rest call to downstream service
	endpoint := os.Getenv("ENDPOINT")
	fmt.Println(endpoint)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: Product wit id  %v %v %v \n", ids, vars["category"], endpoint)
}

func NewsHealthHandler(w http.ResponseWriter, r *http.Request) {
	// Get calling endpoint
	endpoint := os.Getenv("ENDPOINT")
	callback := fmt.Sprintf("%v%v", endpoint, r.RequestURI)

	response, err := http.Get(callback)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Println(string(data))
		io.WriteString(w, string(data))
	}

}/*  */
