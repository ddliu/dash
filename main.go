package main

import (
  "os"
  "github.com/codegangsta/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "dash"
    app.Usage = "start dash"

    app.Action = func(c *cli.Context) {
        println("Hello world!")
    }

    app.Run(os.Args)
}