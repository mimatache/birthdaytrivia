package web

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed build
var app embed.FS

func GetUIHandler() (http.Handler, error) {
	fsys, err := fs.Sub(app, "build")
	if err != nil {
		return nil, fmt.Errorf("failure loading UI: %w", err)
	}

	return http.FileServer(http.FS(fsys)), nil
}
