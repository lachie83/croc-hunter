// The infamous "croc-hunter" game as featured at many a demo
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	httpListenAddr string
	hostname string
	release string
	powered string
)

func main() {
	flag.StringVar(&httpListenAddr, "port", "8080", "HTTP Listen address.")

	flag.Parse()

	log.Println("Starting server...")

	var err error
	hostname, err = os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	release = os.Getenv("WORKFLOW_RELEASE")
        if release == "" {
                release = "unknown"
        }
        powered = os.Getenv("POWERED_BY")
        if powered == "" {
                powered = "Deis"
        }

	// point / at the handler function
	http.HandleFunc("/", httpHandler)

	// serve static content from /static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	log.Println("Server started. Listening on port " + httpListenAddr)
	log.Fatal(http.ListenAndServe(":" + httpListenAddr, nil))
}