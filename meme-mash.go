package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var link = "https://www.reddit.com/r/doge.json"

func handler(w http.ResponseWriter, r *http.Request) {
	var url = link

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
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
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		return
	}

	name := r.PostFormValue("redditname")
	fmt.Print(name + "\n")

	//fmt.Println("handle is called")
	link = "https://www.reddit.com/r/" + name + ".json"
	handler(w, r)
}

func main() {
	fmt.Println("Initializing...")

	//enter := js.Global.Get("document").Call("getElementById", "redditid")
	//enter.Call("addEventListener", "keyup", handleInputKeyUp, false)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	http.HandleFunc("/", handler)
	http.HandleFunc("/post", postHandler)
	http.ListenAndServe(":80", nil)

}
