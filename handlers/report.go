package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"
	"rental_kendaraan/models"
)

// ✅ Halaman filter laporan
func ReportDashboard(w http.ResponseWriter, r *http.Request) {
	Tmpl.ExecuteTemplate(w, "report_filter.html", nil)
}

// ✅ Laporan PDF Bulanan
func ReportPDF(w http.ResponseWriter, r *http.Request) {
	month, _ := strconv.Atoi(r.URL.Query().Get("month"))
	year, _ := strconv.Atoi(r.URL.Query().Get("year"))
	if month == 0 || year == 0 {
		now := time.Now()
		month = int(now.Month())
		year = now.Year()
	}

	data, err := models.GetTransaksiBulanan(month, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(190, 10, fmt.Sprintf("Laporan Transaksi %02d/%d", month, year))
	pdf.Ln(12)

	// Header
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(20, 8, "ID", "1", 0, "C", false, 0, "")
	pdf.CellFormat(50, 8, "Mulai", "1", 0, "C", false, 0, "")
	pdf.CellFormat(50, 8, "Selesai", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 8, "Total", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 8, "Status", "1", 0, "C", false, 0, "")
	pdf.Ln(8)

	// Data
	pdf.SetFont("Arial", "", 11)
	for _, t := range data {
		pdf.CellFormat(20, 8, fmt.Sprintf("%d", t.ID), "1", 0, "C", false, 0, "")
		pdf.CellFormat(50, 8, t.TanggalPinjam, "1", 0, "C", false, 0, "")
		pdf.CellFormat(50, 8, t.TanggalKembali, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 8, fmt.Sprintf("Rp %.0f", t.Total), "1", 0, "R", false, 0, "")
		pdf.CellFormat(30, 8, t.Status, "1", 0, "C", false, 0, "")
		pdf.Ln(8)
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=laporan.pdf")
	_ = pdf.Output(w)
}

// ✅ Laporan PDF Tahunan
func ReportPDFYearly(w http.ResponseWriter, r *http.Request) {
	year, _ := strconv.Atoi(r.URL.Query().Get("year"))
	if year == 0 {
		year = time.Now().Year()
	}

	data, err := models.GetTransaksiYearly(year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(190, 10, fmt.Sprintf("Laporan Transaksi Tahun %d", year))
	pdf.Ln(12)

	// Header
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(20, 8, "ID", "1", 0, "C", false, 0, "")
	pdf.CellFormat(50, 8, "Mulai", "1", 0, "C", false, 0, "")
	pdf.CellFormat(50, 8, "Selesai", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 8, "Total", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 8, "Status", "1", 0, "C", false, 0, "")
	pdf.Ln(8)

	// Data
	pdf.SetFont("Arial", "", 11)
	for _, t := range data {
		pdf.CellFormat(20, 8, fmt.Sprintf("%d", t.ID), "1", 0, "C", false, 0, "")
		pdf.CellFormat(50, 8, t.TanggalPinjam, "1", 0, "C", false, 0, "")
		pdf.CellFormat(50, 8, t.TanggalKembali, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 8, fmt.Sprintf("Rp %.0f", t.Total), "1", 0, "R", false, 0, "")
		pdf.CellFormat(30, 8, t.Status, "1", 0, "C", false, 0, "")
		pdf.Ln(8)
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=laporan_tahunan.pdf")
	_ = pdf.Output(w)
}

// ✅ Laporan Excel Bulanan
func ReportExcel(w http.ResponseWriter, r *http.Request) {
	month, _ := strconv.Atoi(r.URL.Query().Get("month"))
	year, _ := strconv.Atoi(r.URL.Query().Get("year"))
	if month == 0 || year == 0 {
		now := time.Now()
		month = int(now.Month())
		year = now.Year()
	}

	data, err := models.GetTransaksiBulanan(month, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	f := excelize.NewFile()
	sheetName := "Laporan"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Header
	headers := []string{"ID", "ID Pelanggan", "ID Kendaraan", "Tanggal Mulai", "Tanggal Selesai", "Total", "Status"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, h)
	}

	// Data
	row := 2
	for _, t := range data {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), t.ID)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), t.IDPelanggan)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), t.IDKendaraan)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), t.TanggalPinjam)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), t.TanggalKembali)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), t.Total)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), t.Status)
		row++
	}

	// Atur sheet aktif
	f.SetActiveSheet(index)

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"laporan_%02d_%d.xlsx\"", month, year))

	if err := f.Write(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
