package main

import (
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"rental_kendaraan/handlers"
	"rental_kendaraan/models"
)

func main() {
	// Logging console output
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Inisialisasi koneksi database
	if err := models.Init(); err != nil {
		log.Fatal().Err(err).Msg("❌ Gagal koneksi ke database")
	}

	// Parse semua template dengan urutan benar dan tambahan fungsi
	handlers.Tmpl = template.Must(template.New("").Funcs(template.FuncMap{
		"now": func() int { return time.Now().Year() },
		"dict": func(values ...interface{}) map[string]interface{} {
			m := make(map[string]interface{})
			for i := 0; i < len(values); i += 2 {
				m[values[i].(string)] = values[i+1]
			}
			return m
		},
		"html": func(s interface{}) template.HTML {
			switch v := s.(type) {
			case string:
				return template.HTML(v)
			default:
				return ""
			}
		},
	}).ParseFiles(
		"templates/base.html",
		"templates/dashboard.html",
		"templates/pelanggan_list.html",
		"templates/pelanggan_form.html",
		"templates/kendaraan_list.html",
		"templates/kendaraan_form.html",
		"templates/transaksi_list.html",
		"templates/transaksi_form.html",
		"templates/report_filter.html",
		"templates/login.html",
	))

	log.Info().Msg("✅ Template berhasil diparsing")

	// Setup router utama
	r := mux.NewRouter()

	// Register semua route aplikasi
	handlers.RegisterRoutes(r)

	log.Info().Msg("✅ Semua route telah terdaftar")

	// Debug route (opsional)
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		log.Info().Msgf("➡️  %s %v", path, methods)
		return nil
	})

	// Jalankan server
	log.Info().Msg("🚀 Server berjalan di http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal().Err(err).Msg("❌ Server error")
	}
}
