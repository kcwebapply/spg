package main

import (
	"os"

	app "github.com/kcwebapply/spg/app"
)

func main() {
	app := app.AppInit()
	app.Run(os.Args)
}
