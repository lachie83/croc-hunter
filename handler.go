package main

import (
	"fmt"
	"net/http"
)

var html = `
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <title>Croc Hunter</title>
        <link rel='stylesheet' href='/static/game.css'/>
    </head>
    <body>
    	<canvas id="canvasBg" width="800" height="490" ></canvas>
    	<canvas id="canvasEnemy" width="800" height="500" ></canvas>
    	<canvas id="canvasJet" width="800" height="500" ></canvas>
        <canvas id="canvasHud" width="800" height="500" ></canvas>
        <script src='/static/game.js'></script>
        <div class="details">
		<strong>Hostname:</strong> %s<br>
		<strong>Version:</strong> %s<br>
		<strong>Powered By:</strong> %s<br>
        </div>
    </body>
</html>
`

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, html, hostname, release, powered)
}
