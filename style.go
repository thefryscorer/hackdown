package main

const css = `
html {
   margin: 0px;
   height: 100%;
   width: 100%;
}

body {
  padding: 10px;
  font-family: sans-serif;
  font-size: 16pt;
  font-weight:300;
  color: #2d2d2d;
  background-color: #fefefe;
}

.dark_theme {
	color: #fefefe;
	background-color: #2d2d2d;
}

h1 {
  font-size: 22pt;
  font-weight:bold;
}

h2 {
  font-size: 20pt;
}

h3 {
  font-size: 18pt;
}

img {
  width: 50%;
}

blockquote {
  padding: 5px;
  background-color: whitesmoke;
  font-family: monospace;
  font-size: 14pt;
}

.blank_page  {
	width: 100%;
	height: 100%;
	background-color: grey;
	min-height: 100%;
}
`

var css_dark = `
body {
	color: #fefefe;
	background-color: #2d2d2d;
}

blockquote {
	background-color: #3d3d3d;
	color: white;
}

a {
	color: #8080ff;
}

`
