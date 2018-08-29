package learning

import (
	"fmt"
	"log"
	"net/http"
)

func RunSimpleServer() {
	http.HandleFunc("/", handle);
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
