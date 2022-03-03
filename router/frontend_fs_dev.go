//go:build !prod

package router

import (
	"fmt"
	"io/fs"
	"os"
)

func getFrontendAssets() fs.FS {
	fmt.Println(os.Getwd())
	return os.DirFS("frontend/dist")
}
