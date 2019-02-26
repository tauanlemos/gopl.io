// Mo dif y fetch to add the prefix http:// to each argument URL if it is missing.
// Yo u mig ht want to use strings.HasPrefix .

package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") == false {
			url = "http://" + url
		} 
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		c, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", c)
		
	}
}
