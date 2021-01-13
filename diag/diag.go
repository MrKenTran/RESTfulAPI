package diag


import (
	"fmt"
	"log"
	"net/http"
	"time"

)

const APIROOT = "http://api.gbif.org/v1/"
const APIROOT2 = "https://restcountries.eu/rest/v2/"
var startTime time.Time

func uptime() time.Duration {
	return time.Since(startTime)
}
func init() {
	startTime = time.Now()
}

func GetDiagHandler(){
	http.HandleFunc("/conservation/v1/diag/", diagHandler)
}

func diagHandler(w http.ResponseWriter, r *http.Request) {
	 resp, err := http.Get(APIROOT)
	 resp2, err2 := http.Get(APIROOT2)
	if err != nil {
		log.Fatal(err)
	}
	if err2 != nil {
		log.Fatal(err2)
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Fprintf(w,"gbif:", resp.StatusCode)
	} else {
		fmt.Fprintf(w,"Server is broken")
	}

	if resp2.StatusCode >= 200 && resp2.StatusCode <= 299 {
		fmt.Fprintf(w,"\nrestcountries:", resp2.StatusCode)
	} else {
		fmt.Fprintf(w,"Server is broken")
	}
	fmt.Fprintf(w, "\nVersion:", "v1")
	fmt.Fprintf(w,"\nuptime %s",uptime())










}




