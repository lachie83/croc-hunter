/*
    Copyright (c) 2016 Lachlan Evenson

    Permission is hereby granted, free of charge, to any person obtaining
    a copy of this software and associated documentation files (the "Software"),
    to deal in the Software without restriction, including without limitation
    the rights to use, copy, modify, merge, publish, distribute, sublicense,
    and/or sell copies of the Software, and to permit persons to whom the Software
    is furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in
    all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
    EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
    OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
    IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
    CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
    TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
    OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

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