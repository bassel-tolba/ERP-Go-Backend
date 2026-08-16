package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/bassel-tolba/go-erp/internal/db"
	"github.com/bassel-tolba/go-erp/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	conn, err := sql.Open("postgres", "user=bassel dbname=go-erp sslmode=disable password=flstudio")
	if err != nil {
		log.Println(err.Error())
	}
	err = conn.Ping()
	if err != nil {
		log.Println(err.Error())
	}
	queries := db.New(conn)
	svc := service.NewCompanyService(queries)
	http.HandleFunc("/company", func(w http.ResponseWriter, r *http.Request) {
		company, err := svc.CreateCompany(r.Context(), "Test Supplier LLC")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		_, _ = w.Write([]byte("Created: " + company.Name))
	})
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
