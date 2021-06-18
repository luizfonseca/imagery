package handlers

import (
	"net/http"

	"github.com/luizfonseca/imagery/pkg/middleware"
)

// Returns an image
//
// Returns
func ImageHandler(ctx middleware.RouterContext) {
	ctx.Response.WriteHeader(http.StatusCreated)
}
