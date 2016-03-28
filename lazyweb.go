package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var Menus []string

func init() {
	Menus = []string{"CoffeeCat", "ShortFilms", "Fun", "OpenSource", "About"}
}

func www_root(w http.ResponseWriter, r *http.Request) {
	www(w, r, "template/index.html", "Index")
}

func www_about(w http.ResponseWriter, r *http.Request) {
	www(w, r, "template/about.html", "About")
}

func www_opensource(w http.ResponseWriter, r *http.Request) {
	www(w, r, "template/opensource.html", "OpenSource")
}

func www_fun(w http.ResponseWriter, r *http.Request) {
	www(w, r, "template/fun.html", "Fun")
}

// www_coffeecat is little special, because it automatically put image files.
func www_coffeecat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	f, err := os.Open("images/coffeecat")
	if err != nil {
		log.Print(err)
		return
	}
	images, err := f.Readdirnames(-1)
	if err != nil {
		log.Print(err)
		return
	}
	sort.Strings(images)
	for i, img := range images {
		images[i] = filepath.Join("images/coffeecat", img)
	}

	fmap := template.FuncMap{
		"lower": strings.ToLower,
	}
	t, err := template.New("coffeecat.html").Funcs(fmap).ParseFiles("template/coffeecat.html", "template/head.html", "template/menu.html", "template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	info := struct {
		Menus  []string
		MenuOn string
		Images []string
	}{
		Menus:  Menus,
		MenuOn: "CoffeeCat",
		Images: images,
	}
	err = t.Execute(w, info)
	if err != nil {
		log.Fatal(err)
	}
}

func www(w http.ResponseWriter, r *http.Request, page, menuon string) {
	w.Header().Set("Content-Type", "text/html")
	fmap := template.FuncMap{
		"lower": strings.ToLower,
	}
	t, err := template.New(filepath.Base(page)).Funcs(fmap).ParseFiles(page, "template/head.html", "template/menu.html", "template/footer.html")
	if err != nil {
		log.Fatal(err)
		return
	}
	info := struct {
		Menus  []string
		MenuOn string
	}{
		Menus:  Menus,
		MenuOn: menuon,
	}
	err = t.Execute(w, info)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	portPtr := flag.String("http", "", "service port ex):80")
	flag.Parse()
	if *portPtr == "" {
		fmt.Println("lazyweb service")
		flag.PrintDefaults()
		os.Exit(1)
	}
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("template"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.HandleFunc("/", www_root)
	http.HandleFunc("/fun", www_fun)
	http.HandleFunc("/opensource", www_opensource)
	http.HandleFunc("/coffeecat", www_coffeecat)
	http.HandleFunc("/about", www_about)
	log.Fatal(http.ListenAndServe(*portPtr, nil))
}
