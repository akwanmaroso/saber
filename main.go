package main

import "net/http"
import "os"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("helo"))
	})

	if err := http.ListenAndServe(":9090", nil); err != nil {
		os.Exit(1)
	}
}
