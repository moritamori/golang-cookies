package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
)

func main() {
	hashKey := []byte("hash-key")
	// blockKey := []byte("bloooooooock-key")
	s := securecookie.New(hashKey, nil)
	// s := securecookie.New(hashKey, blockKey)

	http.HandleFunc("/set-cookie", func(w http.ResponseWriter,
		r *http.Request) {

		values := map[string]string{
			"title": "SPY x FAMILY",
		}
		encoded, err := s.Encode("cookie-name", values)
		if err == nil {
			cookie := &http.Cookie{
				Name:  "cookie-name",
				Value: encoded,
			}
			http.SetCookie(w, cookie)

			fmt.Fprintf(w, "Cookieをセットしました")
		}
	})

	http.HandleFunc("/show-cookie", func(w http.ResponseWriter,
		r *http.Request) {

		if cookie, err := r.Cookie("cookie-name"); err == nil {
			value := make(map[string]string)

			err = s.Decode("cookie-name", cookie.Value, &value)
			if err == nil {
				fmt.Fprintf(w, "値は%qです。", value["title"])
			} else {
				fmt.Fprintf(w, err.Error())
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
