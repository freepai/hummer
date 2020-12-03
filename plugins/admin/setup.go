package health

import (
	"fmt"
	"github.com/freepai/hummer/core/plugin"
	"github.com/freepai/hummer/core/server"
	"net/http"
)

func setup(ctx *plugin.Context) error {
	mgr := server.GetManager(ctx)
	mgr.HandleFunc("/admin/v1", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "admin info")
	})

	return nil
}

func init() {
	plugin.Register("admin", setup)
}
