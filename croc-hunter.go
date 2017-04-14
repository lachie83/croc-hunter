// The infamous "croc-hunter" game as featured at many a demo
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	httpListenAddr := flag.String("port", "8080", "HTTP Listen address.")

	flag.Parse()

	log.Println("Starting server...")

	// point / at the handler function
	http.HandleFunc("/", handler)

	// serve static content from /static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	log.Println("Server started. Listening on port " + *httpListenAddr)
	log.Fatal(http.ListenAndServe(":"+*httpListenAddr, nil))
}

const (
	html = `
		<html>
			<head>
				<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
				<title>Croc Hunter</title>
				<link rel='stylesheet' href='/static/game.css'/>
				<link rel="icon" type="image/png" href="/static/favicon-16x16.png" sizes="16x16" />
				<link rel="icon" type="image/png" href="/static/favicon-32x32.png" sizes="32x32" />
			</head>
			<body>
				<canvas id="canvasBg" width="800" height="490" ></canvas>
				<canvas id="canvasEnemy" width="800" height="500" ></canvas>
				<canvas id="canvasJet" width="800" height="500" ></canvas>
				<canvas id="canvasHud" width="800" height="500" ></canvas>
				<script src='/static/game.js'></script>
				<div class="details">
				<strong>Hostname: </strong>%s<br>
				<strong>Release: </strong>%s<br>
				<strong>Commit: </strong>%s<br>
				<strong>Powered By: </strong>%s<br>
				</div>
			</body>
		</html>
		`
)

func handler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/healthz" {
		w.WriteHeader(http.StatusOK)
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("could not get hostname: %s", err)
	}

	release := os.Getenv("WORKFLOW_RELEASE")
	commit := os.Getenv("GIT_SHA")
	powered := os.Getenv("POWERED_BY")

	if release == "" {
		release = "unknown"
	}
	if commit == "" {
		commit = "not present"
	}
	if powered == "" {
		powered = "deis"
	}

	fmt.Fprintf(w, html, hostname, release, commit, powered)
}
