package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rhenanteix/goldscan/internal/scan"
)

func main() {

	fmt.Println("Bem vindo essa é minha API")
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	http.HandleFunc("/api/v1/scans", scan.CreateHandler)
	http.HandleFunc("/api/v1/scans/upload", scan.UploadHandler)
	fmt.Println("Servidor rodando na porta 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	// 	item := receivable.Receivable {
	// 	ID:            "rec-001",
	// 	PayerName:     "Empresa ABC",
	// 	PayerDocument: "12345678000190",
	// 	InvoiceNumber: "NF-100",
	// 	Amount:        1250000,
	// }

	// fmt.Println(item)

}
