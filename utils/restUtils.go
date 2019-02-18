package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 20 * time.Second}

//GetJSON makes a simple Get Request, returns a String with a JSON
func GetJSON(url string) string {

	spaceClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 10 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return string(body[:])
}
