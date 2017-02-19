package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type MemeData struct {
	Image    string
	Upvotes  string
	Comments string
}

var arr []MemeData
var verifiedLinks []MemeData

func getMemeArray() []MemeData {
	return verifiedLinks
}

func parse(jsonInput []byte, w http.ResponseWriter) {

	var urls []MemeData

	arr = nil
	verifiedLinks = nil

	fmt.Print("Parsing...")

	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(string(jsonInput)))
	err := dec.Decode(&data)
	tErr(err)
	q := NewQuery(data)
	a, err := q.ArrayOfObjects("data", "children")

	for i := 0; i < len(a); i++ {
		sval, err := q.String("data", "children", strconv.Itoa(i), "data", "url")
		ups, err := q.Int("data", "children", strconv.Itoa(i), "data", "ups")
		comments, err := q.Int("data", "children", strconv.Itoa(i), "data", "num_comments")
		tErr(err)
		//fmt.Print(sval)
		//fmt.Print("|")

		urls = append(urls, MemeData{sval, strconv.Itoa(ups), strconv.Itoa(comments)})
	}

	verifyLinks(urls)

}

func verifyLinks(links []MemeData) {
	arr = links

	for i := 0; i < len(arr); i++ {
		match, _ := regexp.MatchString("/\\.(jpe?g|png|bmp|gif|gifv)$/i", arr[i].Image)
		whitelistlinks := strings.Contains(arr[i].Image, "reddituploads") || strings.Contains(arr[i].Image, "redd.it") || (strings.Contains(arr[i].Image, "imgur") && match)

		if strings.Contains(arr[i].Image, "imgur") && !match {
			arr[i].Image = arr[i].Image + ".png"
		}

		if strings.Contains(arr[i].Image, "&amp;") {
			arr[i].Image = strings.Replace(arr[i].Image, "&amp;", "&", -1)
			//fmt.Println(arr[i].Image)
		}

		if match || whitelistlinks {
			//fmt.Print(arr[i])
			//fmt.Print("|")
			verifiedLinks = append(verifiedLinks, arr[i])
		}

	}
}

func tErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
