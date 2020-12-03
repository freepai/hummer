package controller

import (
	"fmt"
	"github.com/freepai/hummer/core/shorturl/service"
	"net/http"
)

type ShortUrlController struct {
	manager *service.Manager
}

func NewShortUrlController(manager *service.Manager) *ShortUrlController {
	return &ShortUrlController{
		manager: manager,
	}
}

func (c *ShortUrlController) PostShortUrl(w http.ResponseWriter, req *http.Request) {
	ns := req.Header.Get("X-HUMMER-NS")
	longUrl := req.Form.Get("long_url")

	c.manager.Post(ns, longUrl)

	fmt.Fprintf(w, "OK")
}

func (c *ShortUrlController) GetShortUrl(w http.ResponseWriter, req *http.Request) {
	ns := req.Header.Get("X-HUMMER-NS")
	code := req.Form.Get("code")

	su,_ := c.manager.Get(ns, code)

	fmt.Fprintf(w, "%+v", su)
}