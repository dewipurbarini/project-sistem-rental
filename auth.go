package handlers

import (
	"net/http"
	"rental_kendaraan/models"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func LoginForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Error": "",
	}
	err := Tmpl.ExecuteTemplate(w, "login.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := models.GetUserByUsername(username)
	if err != nil || user.Password != password {
		data := map[string]interface{}{
			"Error": "Username atau password salah",
		}
		Tmpl.ExecuteTemplate(w, "login.html", data)
		return
	}

	sess, _ := store.Get(r, "session")
	sess.Values["admin"] = true
	sess.Save(r, w)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	sess, _ := store.Get(r, "session")
	sess.Options.MaxAge = -1
	sess.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
