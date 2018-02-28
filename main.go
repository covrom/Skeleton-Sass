//go:generate go-bindata-assetfs -pkg main -o files.go  index.html images/ css/

package main

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
)

func main() {
	http.Handle("/",
		http.FileServer(
			&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: ""}))
	http.ListenAndServe(":8000", nil)
}
