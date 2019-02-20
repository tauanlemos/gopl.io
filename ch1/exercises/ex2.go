
// Modify the echo program to aslo print os.Args[0], the name of the comand that invoked it

package main

import ( 
	"fmt"
	"os"
	"strconv"
)

func main() {

	for i, arg := range os.Args[1:] {
		fmt.Println(strconv.Itoa(i))
		fmt.Println(arg)
	}
}
