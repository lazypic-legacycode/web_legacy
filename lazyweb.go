package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var Menus []string

func init() {
	Menus = []string{"CoffeeCat", "CatNDog", "ShortFilms", "About"}
}

func www_root(w http.ResponseWriter, r *http.Request) {
	www(w, r, "template/index.html", "Index")
}

func www_shortfilms(w http.ResponseWriter, r *http.Request) {
	www(w, r, "template/shortfilms.html", "ShortFilms")
}

func www_about(w http.ResponseWriter, r *http.Request) {
	www(w, r, "template/about.html", "About")
}

// www_coffeecat shows coffeecat toon pages.
func www_coffeecat(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Path[len("/coffeecat/"):]
	if page == "" {
		www_toon_root(w, r, "CoffeeCat")
		return
	}
	// page should convertable to int
	i, err := strconv.Atoi(page)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	www_toon_page(w, r, "CoffeeCat", i)
}

// www_catndog shows catndog toon pages.
func www_catndog(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Path[len("/catndog/"):]
	if page == "" {
		www_toon_root(w, r, "CatNDog")
		return
	}
	// page should convertable to int
	i, err := strconv.Atoi(page)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	www_toon_page(w, r, "CatNDog", i)
}

// www_toon_root redirects to the last page of the toon.
func www_toon_root(w http.ResponseWriter, r *http.Request, title string) {
	ltitle := strings.ToLower(title)
	f, err := os.Open("toon/" + ltitle)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	images, err := f.Readdirnames(-1)
	if err != nil {
		log.Print(err)
		return
	}
	sort.Sort(sort.Reverse(sort.StringSlice(images)))
	last := -1
	re := regexp.MustCompile("([0-9]+)[.]png$")
	for i := 0; i < len(images); i++ {
		m := re.FindStringSubmatch(images[i])
		if len(m) == 2 {
			last, _ = strconv.Atoi(m[1])
			break
		}
	}
	if last == -1 {
		log.Print("no images matches to regexp.")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/"+ltitle+"/%d", last), http.StatusFound)
}

// www_toon_page shows the page of the toon.
func www_toon_page(w http.ResponseWriter, r *http.Request, title string, i int) {
	ltitle := strings.ToLower(title)
	w.Header().Set("Content-Type", "text/html")

	img := fmt.Sprintf("toon/"+ltitle+"/%02d.png", i)
	_, err := os.Stat(img)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		} else {
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	prev := fmt.Sprintf(ltitle+"/%d", i-1)
	_, err = os.Stat(fmt.Sprintf("toon/"+ltitle+"/%02d.png", i-1))
	if err != nil {
		if os.IsNotExist(err) {
			prev = ""
		} else {
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	next := fmt.Sprintf(ltitle+"/%d", i+1)
	_, err = os.Stat(fmt.Sprintf("toon/"+ltitle+"/%02d.png", i+1))
	if err != nil {
		if os.IsNotExist(err) {
			next = ""
		} else {
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	fmap := template.FuncMap{
		"lower": strings.ToLower,
	}
	t, err := template.New("toon_page.html").Funcs(fmap).ParseFiles("template/toon_page.html", "template/head.html", "template/menu.html", "template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	info := struct {
		Menus  []string
		MenuOn string
		Image  string
		Prev   string
		Next   string
	}{
		Menus:  Menus,
		MenuOn: title,
		Image:  img,
		Prev:   prev,
		Next:   next,
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
	http.Handle("/toon/", http.StripPrefix("/toon/", http.FileServer(http.Dir("toon"))))
	http.HandleFunc("/", www_root)
	http.HandleFunc("/shortfilms", www_shortfilms)
	http.HandleFunc("/coffeecat/", www_coffeecat)
	http.HandleFunc("/catndog/", www_catndog)
	http.HandleFunc("/about", www_about)
	log.Fatal(http.ListenAndServe(*portPtr, nil))
}
