package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"rental_kendaraan/models"
)

func PelangganList(w http.ResponseWriter, r *http.Request) {
	list, err := models.GetAllPelanggan()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{"Pelanggan": list}
	Tmpl.ExecuteTemplate(w, "pelanggan_list.html", data)
}

func PelangganCreateForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"ID": 0, "Nama": "", "Alamat": "", "Telepon": ""}
	Tmpl.ExecuteTemplate(w, "pelanggan_form.html", data)
}

func PelangganCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p := &models.Pelanggan{
		Nama:    r.FormValue("nama"),
		Alamat:  r.FormValue("alamat"),
		Telepon: r.FormValue("telepon"),
	}
	if err := models.CreatePelanggan(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/pelanggan", http.StatusSeeOther)
}

func PelangganEditForm(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	p, err := models.GetPelangganByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Tmpl.ExecuteTemplate(w, "pelanggan_form.html", p)
}

func PelangganUpdate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	p := &models.Pelanggan{
		ID:      id,
		Nama:    r.FormValue("nama"),
		Alamat:  r.FormValue("alamat"),
		Telepon: r.FormValue("telepon"),
	}
	if err := models.UpdatePelanggan(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/pelanggan", http.StatusSeeOther)
}

func PelangganDelete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := models.DeletePelanggan(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/pelanggan", http.StatusSeeOther)
}
