package learning

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func StartCounterServer() {
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/l", handleL)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func handleL(w http.ResponseWriter, r *http.Request) {
	lissajous(w)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w,"Count %d\n", count)
	mu.Unlock()
}
