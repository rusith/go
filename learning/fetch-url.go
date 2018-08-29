package learning

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func FetchUrl() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		fmt.Printf("%s", b)
	}
}