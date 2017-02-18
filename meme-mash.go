package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("handle is called")
	url := "https://www.reddit.com/r/doge.json"
	//resp, _ := http.Get(url)
	//bytes, _ := ioutil.ReadAll(resp.Body)

	//fmt.Fprint(w, string(bytes))

	//resp.Body.Close()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	resp, err := client.Do(req)

	bytes, _ := ioutil.ReadAll(resp.Body)
	//fmt.Fprint(w, string(bytes))
	//fmt.Println(string(bytes))
	parse(bytes, w)
	resp.Body.Close()

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	tpl, err := template.ParseFiles("view.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(w, getImageArray())
}

func main() {
	fmt.Println("Initializing...")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
