package main

import (
	"github.com/freepai/hummer/core"
)

import (
	_ "github.com/freepai/hummer/plugins/servers/http"
)

func main() {
	hummer := core.NewHummer("./hummer.yaml")
	hummer.Start()
}
