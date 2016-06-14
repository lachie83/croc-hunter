package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// grab hostname
	hostname, _ := os.Hostname()
	// construction html
	fmt.Fprintf(w, `<html xmlns="http://www.w3.org/1999/xhtml">
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <title>Tokyo Summit Croc Hunter</title>
        <link rel='stylesheet' href='/static/game.css'/>
    </head>
    <body>

    	<canvas id="canvasBg" width="800" height="490" ></canvas>
    	<canvas id="canvasEnemy" width="800" height="500" ></canvas>
    	<canvas id="canvasJet" width="800" height="500" ></canvas>
        <canvas id="canvasHud" width="800" height="500" ></canvas>
        <script src='/static/game2.js'></script>
        <div class="details">
				<strong>Version:</strong> 2.0<br>
				<strong>Hostname:</strong> %v<br>

        </div>
    </body>
</html>`, hostname)

	// log page hit to stdout
	log.Println("served web page")
}

func main() {
	// log start to stdout
	log.Println("Starting server...")
	// point / at the handler fuction
	http.HandleFunc("/", handler)
	// serve static content from /static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.ListenAndServe(":8080", nil)
}
