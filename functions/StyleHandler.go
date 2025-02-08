package functions

import (
	"net/http"
	"os"
	"strings"
)

func StyleHandler(w http.ResponseWriter, r *http.Request) {
	FilePath := strings.TrimPrefix(r.URL.Path, "/")
	// stores file information in rr
	rr, err := os.Stat(FilePath)
	if err != nil || rr.IsDir() {
		RenderPageNotFound(w)
		return
	}
	http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
}
