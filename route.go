package main

import (
	"encoding/json"
	"fmt"

	"net/http"
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
)
type responseHealth struct {
	Status   uint
	Msg string
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	//r.HandleFunc("/products/{id}", ProductsHandler)
	s := r.PathPrefix("/products").Subrouter()
	// "news/"

	s.HandleFunc("/", ProductHandler)
	// "/products/{key}/"
	s.HandleFunc("/{id}/", ProductHandler)
	// "/products/{key}/details"
	s.HandleFunc("/health", ProductHealthHandler)
	http.ListenAndServe(":8083", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	logging("We entered HomeHandler")
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: Home %v\n", vars["category"])
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ids := vars["id"]
	// rest call to downstream service -- Shopify API call Header
	endpoint := os.Getenv("ENDPOINT")
	fmt.Println(endpoint)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: Product with id  %v %v %v \n", ids, vars["category"], endpoint)
}

func ProductHealthHandler(w http.ResponseWriter, r *http.Request) {

	resph := responseHealth{}


		resph.Msg = "OK"
		resph.Status = http.StatusOK
		res2B, _ := json.Marshal(resph)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, string(res2B))
		logging(string(res2B))

	} /*  */
func logging(message string) (error) {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print(message)
	return nil
}