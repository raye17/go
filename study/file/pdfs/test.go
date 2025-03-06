package pdfs

import (
	"github.com/jung-kurt/gofpdf"
)

func GeneratePdf(filename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(190, 7, "Welcome to topgoer.com", "0", 0, "CM", false, 0, "")
	pdf.ImageOptions(
		"./image/mm.jpg",
		100, 100,
		100, 100,
		false,
		gofpdf.ImageOptions{ImageType: "jpeg", ReadDpi: true},
		0,
		"",
	)
	return pdf.OutputFileAndClose(filename)
}
