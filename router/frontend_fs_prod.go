//go:build prod

package router

import (
	"embed"
	"io/fs"
)

//go:generate cp -r ../frontend/dist ./dist
//go:embed dist
var frontend embed.FS

func getFrontendAssets() fs.FS {
	f, err := fs.Sub(frontend, "dist")
	if err != nil {
		panic(err)
	}

	return f
}
