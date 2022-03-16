package main

import (
	"fmt"
	"gy-serv-init/src/core"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 2 {
		panic("Too many arguments")
	} else if len(os.Args) < 2 {
		panic("Too few arguments")
	}

	fmt.Printf("Creating a project named %s\n", os.Args[1])
	core.Main(&pwd, &os.Args[1])
}
