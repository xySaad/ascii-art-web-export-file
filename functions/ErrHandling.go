package functions

import (
	"fmt"
	"net/http"
)

func ErrHandling(err error, j int, w http.ResponseWriter) {
	if err != nil {
		// Send HTML response with the error message for the pop-up
		if j == 1 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `<html>
			<head><script ">alert("⚠️ %s");</script></head>
			<body>
			</body>
		</html>`, err.Error())
			return
		} else if j == 2 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `<html>
			<head><script ">alert("⚠️ %s");</script></head>
			<body>
			</body>
		</html>`, err.Error())
			return
		}
	}
}
