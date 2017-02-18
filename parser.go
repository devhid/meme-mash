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

var arr []string
var verifiedLinks []string

func getImageArray() []string {
	return verifiedLinks
}

func parse(jsonInput []byte, w http.ResponseWriter) {

	urls := ""

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
		tErr(err)
		//fmt.Print(sval)
		//fmt.Print("|")

		urls += sval + "|"
	}

	verifyLinks(urls)

}

func verifyLinks(links string) {
	arr = strings.Split(links, "|")

	for i := 0; i < len(arr); i++ {
		match, _ := regexp.MatchString("/\\.(jpe?g|png|bmp)$/i", arr[i])
		whitelistlinks := strings.Contains(arr[i], "reddituploads") || (strings.Contains(arr[i], "imgur") && match)

		if strings.Contains(arr[i], "imgur") && !match {
			arr[i] = arr[i] + ".png"
		}

		if match || whitelistlinks {
			fmt.Print(arr[i])
			fmt.Print("|")
			verifiedLinks = append(verifiedLinks, strings.Replace(arr[i], "&amp;", "&", -1))
		}

	}
}

func tErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
