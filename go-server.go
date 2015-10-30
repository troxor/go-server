package main

import "fmt"
import "net/http"
import "time"
import "html"
import "log"
import "github.com/fvbock/endless"

func main() {
	log.Print("Starting net/http server on 8000")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2*time.Second)
		fmt.Fprintf(w, "Hello %q\n", html.EscapeString(r.URL.Path))
	})
	err := endless.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


