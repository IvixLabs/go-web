package static

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed resources
var resources embed.FS

func GetFileHandler(developmentMode bool, staticDir string) http.Handler {

	resourceDir, err := fs.Sub(resources, "resources")
	if err != nil {
		panic(err)
	}

	var resourcesFS http.FileSystem

	if developmentMode {
		resourcesFS = http.Dir(staticDir)
	} else {
		resourcesFS = http.FS(resourceDir)
	}

	return http.FileServer(resourcesFS)
}
