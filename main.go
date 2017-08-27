package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// We want a statically linked binary. The net package seems to be the challenge,
	// so make sure we use that.
	resp, err := http.Get("https://example.org")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}
