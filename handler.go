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
