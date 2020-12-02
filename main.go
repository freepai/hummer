package main

import (
	"github.com/freepai/hummer/core"
	_ "github.com/freepai/hummer/plugins"
)

func main() {
	hummer := core.NewHummer("./hummer.yaml")
	hummer.Start()
}
