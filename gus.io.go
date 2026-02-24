package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Print("Listening on port 80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func thetimenow() string {
	nownow := time.Now()
	returntime := nownow.Format("2006-01-02 15:04:05")
	return returntime
}

func handler(w http.ResponseWriter, r *http.Request) {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return
	}

	w.Write([]byte("{\n"))
	w.Write([]byte("    \"api.gus.io\" : {\n"))
	w.Write([]byte("        \"now_utc\" : \""))
	w.Write([]byte(thetimenow()))
	w.Write([]byte("\",\n"))
	w.Write([]byte("        \"Your_IP\" : \""))
	w.Write([]byte(host))
	w.Write([]byte("\",\n"))
	w.Write([]byte("        \"built with\" : \"golang net/http\",\n"))
	w.Write([]byte("        \"more info\" : \"http://www.gus.io/\"\n"))
	w.Write([]byte("    }\n"))
	w.Write([]byte("}\n"))
}
