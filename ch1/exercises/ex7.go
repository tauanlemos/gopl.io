// The function cal l io.Copy(dst, src) re ads fro m src and writes to dst . Use it
// instead of ioutil.ReadAll to copy the respons e body to os.Stdout without requir ing a
// buffer large enoug h to hold the ent ire stream. Be sure to che ck the error result of io.Copy .
// 
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
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
