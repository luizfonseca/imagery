package handlers

import (
	"fmt"
	"io"

	"github.com/luizfonseca/imagery/pkg/middleware"
)

// Returns an image
//
// Returns
func ImageHandler(ctx middleware.RouterContext) {
	fetchUrl := ctx.Request.URL.Query()["url"][0]
	image := ctx.Fetch(middleware.FetchOptions{Method: "GET", Url: fetchUrl})

	body, err := io.ReadAll(image.Body)
	if err != nil {
		ctx.Logger.Error(fmt.Sprintf("Failed to fetch: %v", err))
	}

	ctx.Response.Header().Add("Content-Type", "image/jpg")
	ctx.Response.Header().Add("Cache-Control", "public, max-age=604800, immutable")

	defer ctx.Response.Write(body)

}
