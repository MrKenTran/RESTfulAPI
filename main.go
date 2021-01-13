// handler
// request
// mapt

package main

import (
	"fmt"
	"log"
	"net/http"
	B "oblig1/country"
	C "oblig1/diag"
	A "oblig1/sp1"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Velkommen til hovedsiden")
}

func handleRequest() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	A.GetSpecies()
	B.GetCountry()
	C.GetDiagHandler()
	handleRequest()
}
