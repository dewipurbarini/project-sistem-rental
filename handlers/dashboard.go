package handlers

import (
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Dashboard",
	}
	err := Tmpl.ExecuteTemplate(w, "dashboard.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


