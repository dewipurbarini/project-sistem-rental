package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func InitTemplates() {
	Tmpl = template.Must(template.New("").Funcs(template.FuncMap{
		"now": func() int { return time.Now().Year() },
	}).ParseGlob("templates/*.html"))

	log.Println("Templates parsed.")
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess, _ := store.Get(r, "session")
		if strings.HasPrefix(r.URL.Path, "/static/") || r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}
		if sess.Values["admin"] == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RegisterRoutes(r *mux.Router) {
	r.Use(AuthMiddleware)
	r.HandleFunc("/dashboard", Dashboard).Methods("GET")
	r.HandleFunc("/pelanggan", PelangganList).Methods("GET")
	r.HandleFunc("/kendaraan", KendaraanList).Methods("GET")
	r.HandleFunc("/transaksi", TransaksiList).Methods("GET")
	r.HandleFunc("/report", ReportDashboard).Methods("GET")

	// Static
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Auth
	r.HandleFunc("/login", LoginForm).Methods("GET")
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/logout", Logout).Methods("POST")

	// Root
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	})

	// Dashboard
	r.HandleFunc("/dashboard", Dashboard).Methods("GET")

	// Pelanggan
	r.HandleFunc("/pelanggan", PelangganList).Methods("GET")
	r.HandleFunc("/pelanggan/new", PelangganCreateForm).Methods("GET")
	r.HandleFunc("/pelanggan", PelangganCreate).Methods("POST")
	r.HandleFunc("/pelanggan/{id}/edit", PelangganEditForm).Methods("GET")
	r.HandleFunc("/pelanggan/{id}", PelangganUpdate).Methods("POST")
	r.HandleFunc("/pelanggan/{id}/delete", PelangganDelete).Methods("POST")

	// Kendaraan
	r.HandleFunc("/kendaraan", KendaraanList).Methods("GET")
	r.HandleFunc("/kendaraan/new", KendaraanCreateForm).Methods("GET")
	r.HandleFunc("/kendaraan", KendaraanCreate).Methods("POST")
	r.HandleFunc("/kendaraan/{id}/edit", KendaraanEditForm).Methods("GET")
	r.HandleFunc("/kendaraan/{id}", KendaraanUpdate).Methods("POST")
	r.HandleFunc("/kendaraan/{id}/delete", KendaraanDelete).Methods("POST")

	// Transaksi
	r.HandleFunc("/transaksi", TransaksiList).Methods("GET")
	r.HandleFunc("/transaksi/new", TransaksiCreateForm).Methods("GET")
	r.HandleFunc("/transaksi", TransaksiCreate).Methods("POST")
	r.HandleFunc("/transaksi/{id}/nota", TransaksiNotaPDF).Methods("GET")
	r.HandleFunc("/transaksi/{id}/selesai", TransaksiSelesai).Methods("POST")
	r.HandleFunc("/transaksi/{id}/delete", TransaksiDelete).Methods("POST")

	// Report
	r.HandleFunc("/report", ReportDashboard).Methods("GET")
	r.HandleFunc("/report/pdf", ReportPDF).Methods("GET")
	r.HandleFunc("/report/pdf/yearly", ReportPDFYearly).Methods("GET")
	r.HandleFunc("/report/excel", ReportExcel).Methods("GET")

}
