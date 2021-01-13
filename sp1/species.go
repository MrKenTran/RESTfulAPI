package sp1
// Hvordan lagre data fra api kilde til struct
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)
	const APIROOT = "http://api.gbif.org/v1/species/"   // Rootpath til nettsiden


// setter opp mapt/array
   type Species struct {
	Key 		   int     `json:"key"`
	Kingdom        string `json:"kingdom"`
	Phylum         string `json:"phylum"`
	Family         string `json:"family"`
	Genus          string `json:"genus"`
	ScientificName string `json:"scientificName"`
	CanonicalName  string `json:"canonicalName"`
}
  type Year struct {
		Year string `json:"year"`
	}

	type MergedY struct {
		Year Year
		Species Species
	}



func GetSpecies(){
	http.HandleFunc("/conservation/v1/species/", handlerSpecies)
}

func handlerSpecies(w http.ResponseWriter, r *http.Request) {
	request := strings.Split(r.URL.Path, "/conservation/v1/species/") // split the url path
	query := request[1]                                               // entry path from user after last "/"
	url := APIROOT + query
	url2 := url + "/name"
	response, err := http.Get(url) // respons the url from user
	response2, err2 := http.Get(url2)

	if err != nil { //error handler
		fmt.Println("Error: we have a problem", err)
	}

	if err2 != nil {
		fmt.Println("Error: we have a problem", err2)
	}


	var content Species
	var content2 Year
	var merged MergedY
	err = json.NewDecoder(response.Body).Decode(&content)
	err2 = json.NewDecoder(response2.Body).Decode(&content2)
	merged.Species = content
	merged.Year = content2
	var buffer = new(bytes.Buffer)
	encode := json.NewEncoder(buffer)
	encode.Encode(merged)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	io.Copy(w, buffer)

}


