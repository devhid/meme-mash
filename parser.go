package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func parse(jsonInput []byte, w http.ResponseWriter) {
	// fmt.Print("Parsing...")
	// var data = new(Data)
	// var children []Children = data.children

	// if err := json.Unmarshal(jsonInput, &children); err == nil {
	// 	fmt.Fprint(w, children[0].data.url)
	// } else {
	// 	fmt.Fprint(w, "whoops:", err)
	// }

    urls := ""

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
		// fmt.Print(sval)
		// fmt.Print("|")

        urls += sval + "|"
	}

    verifyLinks(sval)

}

func verifyLinks(links string) {
    arr := strings.split(links, "|")

    for i := 0; i < len(arr); i++ {
        match, _ := regexp.Compile("(?i)^.*\.(jpg|png|gif)$", arr[i])
        whitelistlinks := strings.Contains(arr[i], "reddituploads") || strings.Contains(arr[i], "imgur")

        if (!match || !whitelistlinks) {
            arr[i] = nil
        } else {
            fmt.Print(arr[i])
		    fmt.Print("|")
        }

    }
}

func tErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
