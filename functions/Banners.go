package functions

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

func ListBanners() (banners []string, status int) {
	entries, err := os.ReadDir("./banners")
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		bannerName := strings.Split(entry.Name(), ".")
		banners = append(banners, bannerName[0])
	}
	return banners, http.StatusOK
}

func Banners(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RenderError(w, http.StatusMethodNotAllowed)
		return
	}
	banners, status := ListBanners()
	if status != 200 {
		RenderError(w, status)
		return
	}

	bannersByte, err := json.Marshal(banners)
	if err != nil {
		RenderError(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bannersByte)
}
