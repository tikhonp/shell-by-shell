package main

import (
	"fmt"
	"os"

	shellbyshell "github.com/tikhonp/shell-by-shell"
)

func main() {
    cfg, err := shellbyshell.ParseFlags()
    if err != nil {
        fmt.Println("ERROR: ", err.Error())
        os.Exit(1)
    }
    r, err := shellbyshell.Download(cfg.Url)
    if err != nil {
        fmt.Println("ERROR: ", err.Error())
        os.Exit(1)
    }
    steps, err := shellbyshell.ParseFile(r)
    if err != nil {
        fmt.Println("ERROR: ", err.Error())
        os.Exit(1)
    }
    fmt.Printf("%+v\n", steps)
}
