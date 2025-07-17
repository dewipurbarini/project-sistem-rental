package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	"rental_kendaraan/models"
)

func TransaksiList(w http.ResponseWriter, r *http.Request) {
	list, err := models.GetAllTransaksi()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Tmpl.ExecuteTemplate(w, "transaksi_list.html", map[string]interface{}{"Transaksi": list})
}

func TransaksiCreateForm(w http.ResponseWriter, r *http.Request) {
	pelanggan, _ := models.GetAllPelanggan()
	kendaraan, _ := models.GetAllKendaraan()
	data := map[string]interface{}{
		"ID":             0,
		"Pelanggan":      pelanggan,
		"Kendaraan":      kendaraan,
		"TanggalMulai":   "",
		"TanggalSelesai": "",
		"Total":          0,
		"Status":         "baru",
	}
	Tmpl.ExecuteTemplate(w, "transaksi_form.html", data)
}

func TransaksiCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pid, _ := strconv.Atoi(r.FormValue("id_pelanggan"))
	kid, _ := strconv.Atoi(r.FormValue("id_kendaraan"))
	total := parseFloat(r.FormValue("total"))
	t := &models.Transaksi{
		IDPelanggan:    pid,
		IDKendaraan:    kid,
		TanggalPinjam:   r.FormValue("tanggal_pinjam"),
		TanggalKembali: r.FormValue("tanggal_kembali"),
		Total:          total,
		Status:         "baru",
	}
	if err := models.CreateTransaksi(t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/transaksi", http.StatusSeeOther)
}

func TransaksiEditForm(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	t, err := models.GetTransaksiByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if t == nil {
		http.NotFound(w, r)
		return
	}
	pelanggan, _ := models.GetAllPelanggan()
	kendaraan, _ := models.GetAllKendaraan()
	data := map[string]interface{}{
		"ID":             t.ID,
		"Pelanggan":      pelanggan,
		"Kendaraan":      kendaraan,
		"TanggalPinjam":   t.TanggalPinjam,
		"TanggalKembali": t.TanggalKembali,
		"Total":          t.Total,
		"Status":         t.Status,
	}
	Tmpl.ExecuteTemplate(w, "transaksi_form.html", data)
}

func TransaksiUpdate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	pid, _ := strconv.Atoi(r.FormValue("id_pelanggan"))
	kid, _ := strconv.Atoi(r.FormValue("id_kendaraan"))
	total := parseFloat(r.FormValue("total"))
	t := &models.Transaksi{
		ID:             id,
		IDPelanggan:    pid,
		IDKendaraan:    kid,
		TanggalPinjam:   r.FormValue("tanggal_pinjam"),
		TanggalKembali: r.FormValue("tanggal_kembali"),
		Total:          total,
		Status:         r.FormValue("status"),
	}
	if err := models.UpdateTransaksi(t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/transaksi", http.StatusSeeOther)
}

func TransaksiDelete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := models.DeleteTransaksi(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/transaksi", http.StatusSeeOther)
}

func TransaksiSelesai(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	tr, err := models.GetTransaksiByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if tr == nil {
		http.NotFound(w, r)
		return
	}
	tr.Status = "selesai"
	if err := models.UpdateTransaksi(tr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/transaksi", http.StatusSeeOther)
}

func TransaksiNotaPDF(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	tr, err := models.GetTransaksiByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if tr == nil {
		http.NotFound(w, r)
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(190, 10, "STRUK TRANSAKSI - LANMADJA SEJAHTERA")
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(50, 8, fmt.Sprintf("ID Transaksi: %d", tr.ID))
	pdf.Ln(8)
	pdf.Cell(50, 8, fmt.Sprintf("ID Pelanggan: %d", tr.IDPelanggan))
	pdf.Ln(8)
	pdf.Cell(50, 8, fmt.Sprintf("ID Kendaraan: %d", tr.IDKendaraan))
	pdf.Ln(8)
	pdf.Cell(50, 8, fmt.Sprintf("Tanggal Pinjam: %s", tr.TanggalPinjam))
	pdf.Ln(8)
	pdf.Cell(50, 8, fmt.Sprintf("Tanggal Kembali: %s", tr.TanggalKembali))
	pdf.Ln(8)
	pdf.Cell(50, 8, fmt.Sprintf("Total: Rp %.0f", tr.Total))
	pdf.Ln(8)
	pdf.Cell(50, 8, fmt.Sprintf("Status: %s", tr.Status))

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=nota.pdf")
	_ = pdf.Output(w)
}
