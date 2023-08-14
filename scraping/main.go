package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://api.chucknorris.io/jokes/random"
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR!")
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("ERROR!")
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var raw json.RawMessage
	err = json.Unmarshal(data, &raw)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	fmt.Println(string(raw))
}
