package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"rental_kendaraan/models"
)

// ✅ List semua kendaraan
func KendaraanList(w http.ResponseWriter, r *http.Request) {
	list, err := models.GetAllKendaraan()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Kendaraan": list,
	}

	if err := Tmpl.ExecuteTemplate(w, "kendaraan_list.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ✅ Tampilkan form tambah kendaraan
func KendaraanCreateForm(w http.ResponseWriter, r *http.Request) {
	// Data kosong untuk form baru
	data := map[string]interface{}{
		"ID":          0,
		"Tipe":        "",
		"Merk":        "",
		"NomorPolisi": "",
		"HargaSewa":   0,
		"Status":      "tersedia",
	}
	if err := Tmpl.ExecuteTemplate(w, "kendaraan_form.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ✅ Proses simpan kendaraan baru
func KendaraanCreate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	k := &models.Kendaraan{
		Tipe:        r.FormValue("tipe"),
		Merk:        r.FormValue("merk"),
		NomorPolisi: r.FormValue("nomor_polisi"),
		HargaSewa:   parseFloat(r.FormValue("harga_sewa")),
		Status:      r.FormValue("status"),
	}

	if err := models.CreateKendaraan(k); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/kendaraan", http.StatusSeeOther)
}

// ✅ Tampilkan form edit kendaraan
func KendaraanEditForm(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	k, err := models.GetKendaraanByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if k == nil {
		http.NotFound(w, r)
		return
	}

	// Kirim data kendaraan ke form
	if err := Tmpl.ExecuteTemplate(w, "kendaraan_form.html", k); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ✅ Proses update kendaraan
func KendaraanUpdate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	k := &models.Kendaraan{
		ID:          id,
		Tipe:        r.FormValue("tipe"),
		Merk:        r.FormValue("merk"),
		NomorPolisi: r.FormValue("nomor_polisi"),
		HargaSewa:   parseFloat(r.FormValue("harga_sewa")),
		Status:      r.FormValue("status"),
	}

	if err := models.UpdateKendaraan(k); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/kendaraan", http.StatusSeeOther)
}

// ✅ Hapus kendaraan
func KendaraanDelete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := models.DeleteKendaraan(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/kendaraan", http.StatusSeeOther)
}

// ✅ Helper konversi string ke float
func parseFloat(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}
