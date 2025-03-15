package functions

import (
	"net/http"
	"os"
)

func StyleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RenderError(w,http.StatusMethodNotAllowed)
		return
	}
	rr, err := os.Stat(r.URL.Path[1:])
	if err != nil || rr.IsDir() {
		RenderError(w, http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
