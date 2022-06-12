package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/set-cookie", func(w http.ResponseWriter,
		r *http.Request) {

		cookie := &http.Cookie{
			Name:  "title",
			Value: "SPY x FAMILY",
		}
		http.SetCookie(w, cookie)

		fmt.Fprintf(w, "Cookieをセットしました")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
