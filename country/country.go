package country

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

//Api URL
const APIROOT = "https://restcountries.eu/rest/v2/alpha/"
const APIROOT2 = "http://api.gbif.org/v1/occurrence/"

type Country struct {
	Code        string `json:"alpha2Code"`
	CountryName string `json:"name"`
	CountryFlag string `json:"flag"`
}

type Result struct {
	Species    string `json:"species"`
	SpeciesKey int    `json:"speciesKey"`
}

type Occ struct {
	Results []Result //`json:"results"`

}

type Merge struct {
	Country Country
	Occ     []Result
}

func GetCountry() {
	http.HandleFunc("/conservation/v1/country/", handlerCountry)
}

func handlerCountry(w http.ResponseWriter, r *http.Request) {
	request := strings.Split(r.URL.Path, "/conservation/v1/country/")
	query := request[1]
	url := APIROOT + query
	url2 := APIROOT2 + "search?country=" + query // "/ landskode for info om landet"
	response, err := http.Get(url)
	response2, err2 := http.Get(url2)

	if err != nil { //error handler
		fmt.Println("Error: we have a problem", err)
	}
	if err2 != nil { //error handler
		fmt.Println("Error: we have a problem", err2)
	}
	var part1 Country
	var part2 Occ
	err = json.NewDecoder(response.Body).Decode(&part1)
	err2 = json.NewDecoder(response2.Body).Decode(&part2)
	var merged Merge
	merged.Country = part1
	merged.Occ = part2.Results
	var buffer = new(bytes.Buffer)
	encode := json.NewEncoder(buffer)
	encode.Encode(merged)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	io.Copy(w, buffer)

}
