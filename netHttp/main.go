package main

import (
	"fmt"
	"log"
	"net/http"
)

// But du jeu
// Trouver les codes HTTP des URLS passées en params

type site struct {
	Name string
}

func main() {
	s := site{Name: "https://golang.org/pkg/net/http/eiu/"}
	gHTTPCode(s)
}

func gHTTPCode(s site) {
	resp, err := http.Get(s.Name)
	if err != nil {
		log.Fatalf("error while getting resp from %v %v", s, err)
	}
	code := resp.StatusCode
	fmt.Println(code, http.StatusText(code))
}
