package functions

import (
	"fmt"
	"html/template"
	"net/http"

)

func RenderPageNotFound(w http.ResponseWriter, Status int) {
	tmpl, err := template.ParseFiles("templates/page-not-found.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		fmt.Println("Error parsing template:", err)
		return
	}
	w.WriteHeader(Status)
	type msgStatus struct {
		Messages map[int]string
	}

	statusMessages := msgStatus{
		Messages: map[int]string{
			404: "We can't find the page you're looking for.",
			500: "Please try again later.",
			400: "Bad request.",
		},
	}
	data := map[string]interface{}{
		"Status":        Status,
		"RollBackerror":statusMessages.Messages[Status],
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Could not execute template", http.StatusInternalServerError)
		fmt.Println("Error executing template:", err)
	}
}
