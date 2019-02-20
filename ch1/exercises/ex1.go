// Modify the echo program to aslo print os.Args[0], the name of the comand that invoked it

package main

import ( 
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[:], " "))
}
