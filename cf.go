package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/cardigann/go-cloudflare-scraper"
)

func CS() {
	scraper, err := scraper.NewTransport(http.DefaultTransport)
	if err != nil {
		log.Fatal(err)
	}

	c := http.Client{Transport: scraper}

	res, err := c.Get("https://csgostats.gg/match")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if err == nil {
		fmt.Print(string(body))
	}

	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}
