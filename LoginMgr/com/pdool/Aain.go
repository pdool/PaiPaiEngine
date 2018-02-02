package Coremain

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("xxx")
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter your name:")
	input, err := inputReader.ReadString('\n')
	    if err != nil {
		         fmt.Println("There were errors reading, exiting program.")
		         return
		    }
	     fmt.Printf("Your name is %s", input)
}
