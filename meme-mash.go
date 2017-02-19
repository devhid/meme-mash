package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var link = "https://www.reddit.com/r/doge.json"
var hasPosted = false

func handler(w http.ResponseWriter, r *http.Request) {
	var url = link

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	link = "https://www.reddit.com/r/doge.json"

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	resp, err := client.Do(req)

	bytes, _ := ioutil.ReadAll(resp.Body)

	parse(bytes, w)
	resp.Body.Close()

	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(w, getMemeArray())

	if !hasPosted {
		http.HandleFunc("/post", postHandler)
		hasPosted = true
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		return
	}

	name := r.PostFormValue("redditname")

	if name == "" {
		link = "https://www.reddit.com/r/doge.json"
	} else {
		link = "https://www.reddit.com/r/" + name + ".json"
	}

	handler(w, r)
}

func main() {
	fmt.Println("Initializing...")

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
