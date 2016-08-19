package main

import (
	"os"

	"github.com/shinmyung0/autoscaler/cli"
)

func main() {

	args := os.Args[1:]
	os.Exit(cli.Run(args))

}
