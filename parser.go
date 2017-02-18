package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	child []Children `json:"children"`
}

type Children struct {
	data Data2 `json:"data"`
}

type Data2 struct {
	url string `json:"url"`
}

func parse(j []byte) (Data, error) {
	fmt.Print("Parsing...")
	//var m Message
	var d Data
	err := json.Unmarshal(j, &d)

	if err != nil {
		fmt.Println("whoops:", err)
	}
	return d, err
}
