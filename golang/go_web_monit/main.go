// This program have one main
// - Monitore website with http get status code answer
package main

import (
	"net/http"
	"log"
	"fmt"
)

// Define good answer from get
const okStatus int = 200

func main() {

	mapSites := map[string]string{
		"youtube": "http://youtube.com",
		"google": "http://google.com",
		"yahoo": "http://yahoo.com",
	}

	mapSatusCode := map[string]int{}

	for k, v := range mapSites{
		_, statusCode, err := getStatusCode(v)
		if err != nil {
		log.Fatal(err)
		}
		fmt.Printf("URL : %v is %v [OK]\n", k, statusCode)
		mapSatusCode[v] = statusCode
	}

	for k, v := range mapSatusCode {
		err := checkStatusCode(v, k)
		if err != nil {
			log.Println(err)
		}
	}
}

func getStatusCode(url string) (site string, statusCode int, err error){
	site = url
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	statusCode = res.StatusCode
	return
}

func checkStatusCode(code int, url string) (err error){
	c := code
	if c != okStatus {
		err = fmt.Errorf("Error code is %v responding from %s", c, url)
	}
	return err
}