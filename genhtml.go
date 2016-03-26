package main

import (
	"fmt"
)

func head(title string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<head>
<title>Lazypic %s</title>
<meta charset="utf-8">
<link rel="stylesheet" type="text/css" href="/template/lazyweb.css">
<link rel="icon" type="image/png" href="/images/icon/lazyweb.png">
</head>
<body>`, title)
}

func menu(sel string) string {
	var rstring string
	rstring += `<div class="header"><div class="logo"><a href="/"><img src="/images/lazypic_logo.png" width=180 height=60></a></div>`
	menus := []string{"coffeecat","shortfilms","fun","opensource","about"}
	for _, i := range menus {
		if i == sel {
			rstring += fmt.Sprintf(`<div class="menuon">%s</div>`, i)
		} else {
			rstring += fmt.Sprintf(`<div class="menu"><a href="/%s">%s</a></div>`, i, i)
		}
	}
	return rstring + "</div>"
}

func tail() string {
	return`
<div id="footer">
Copyright © 2016 All rights Reserved by Lazypictures.
</div></body></html>
`
}
