package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/s-li1/interpreter/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
      panic(err)
    }

    fmt.Printf("Hello %s! Test this out!\n", user.Username)
    repl.Start(os.Stdin, os.Stdout)
    fmt.Printf("Type something\n")
}
