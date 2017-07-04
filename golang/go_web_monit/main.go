package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ashwanthkumar/slack-go-webhook"
)

const okStatus int = 200

func main() {

	for {

		mapURL := map[string]string{
			"url1":          "https://www.test.com/1,
			"url2":          "https://www.test.com/2",
		}

		mapStatusCode := map[string]int{}

		for _, v := range mapURL {
			_, statusCode, err := getStatusCode(v)
			if err != nil {
				payload("impossible to get status code" + err.Error())
			}
			mapStatusCode[v] = statusCode
		}

		for k, v := range mapStatusCode {
			err := checkStatusCode(v, k)
			if err != nil {
				payload(err.Error())
			}
		}

		time.Sleep(60 * time.Second)

	}
}

func payload(text string) {

	webhookURL := "https://hooks.slack.com/services/WEBHOOK"

	payload := slack.Payload{
		Text:     text,
		Username: "test,
		Channel:  "#infra",
	}
	err := slack.Send(webhookURL, "", payload)
	if len(err) > 0 {
		fmt.Printf("error : %s\n", err)
	}

}

func getStatusCode(url string) (site string, statusCode int, err error) {
	site = url
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	statusCode = res.StatusCode
	return
}

func checkStatusCode(code int, url string) (err error) {
	c := code
	if c != okStatus {
		err = fmt.Errorf("Error code is %v responding from %s", c, url)
	}
	return err
}
