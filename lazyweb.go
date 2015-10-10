package main

import (
	"fmt"
	"flag"
	"os"
	"net"
	"net/http"
	"html/template"
	)

func localIP() string {
	var iplist []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Error: " + err.Error() + "\n")
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				iplist = append(iplist, ipnet.IP.String())
			}
		}
	}
	return iplist[0]
}

type Web struct {
	Title string
}

func www_root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, _ := template.ParseFiles("template/index.html")
	p := Web{Title: "Lazypic"}
	t.Execute(w, p)
}

func www_about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, _ := template.ParseFiles("template/about.html")
	p := Web{Title: "Lazypic:about"}
	t.Execute(w, p)
}

func www_opensource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, _ := template.ParseFiles("template/opensource.html")
	p := Web{Title: "Lazypic:Opensource"}
	t.Execute(w, p)
}


func www_fun(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, _ := template.ParseFiles("template/fun.html")
	p := Web{Title: "Lazypic:Fun"}
	t.Execute(w, p)
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
	http.HandleFunc("/about", www_about)

	http.ListenAndServe(*portPtr,nil)
}
