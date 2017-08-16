package main

import (
	"net/http"

	"github.com/g0tzer0/coyote/util"
	"github.com/g0tzer0/coyote/web/controller"
)

func main() {
	controller.Setup()

	http.ListenAndServeTLS(":8443", "../certs/server.crt", "../certs/server.key", new(util.GzipHandler))
}
