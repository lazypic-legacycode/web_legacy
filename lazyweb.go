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
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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

	if r.URL.Path == "/coffeecat/" {
		www_coffeecat_root(w, r)
	} else {
		subURL := r.URL.Path[len("/coffeecat/"):]
		// sub url should convertable to int
		// ex) r.URL.Path == "/coffeecat/0"
		i, err := strconv.Atoi(subURL)
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		www_coffeecat_page(w, r, i)
	}
}

func www_coffeecat_root(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("images/coffeecat")
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
	http.Redirect(w, r, fmt.Sprintf("/coffeecat/%d", last), http.StatusFound)
}

func www_coffeecat_page(w http.ResponseWriter, r *http.Request, i int) {
	img := fmt.Sprintf("images/coffeecat/%02d.png", i)
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
	prev := fmt.Sprintf("coffeecat/%d", i-1)
	_, err = os.Stat(fmt.Sprintf("images/coffeecat/%02d.png", i-1))
	if err != nil {
		if os.IsNotExist(err) {
			prev = ""
		} else {
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	next := fmt.Sprintf("coffeecat/%d", i+1)
	_, err = os.Stat(fmt.Sprintf("images/coffeecat/%02d.png", i+1))
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
	t, err := template.New("coffeecat_page.html").Funcs(fmap).ParseFiles("template/coffeecat_page.html", "template/head.html", "template/menu.html", "template/footer.html")
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
		MenuOn: "CoffeeCat",
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
	http.HandleFunc("/", www_root)
	http.HandleFunc("/fun", www_fun)
	http.HandleFunc("/opensource", www_opensource)
	http.HandleFunc("/coffeecat/", www_coffeecat)
	http.HandleFunc("/about", www_about)
	log.Fatal(http.ListenAndServe(*portPtr, nil))
}
