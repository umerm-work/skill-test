package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// Student holds the data for a single student.
type Student struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	ClassName string `json:"class_name"`
}

// generatePDF creates a PDF report for a given student.
func generatePDF(student Student) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(40, 10, "Student Report")
	pdf.Ln(20)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Student ID: %d", student.ID))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Name: %s %s", student.FirstName, student.LastName))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Email: %s", student.Email))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Class: %s", student.ClassName))
	pdf.Ln(10)

	var buf strings.Builder
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return []byte(buf.String()), nil
}

// reportHandler handles the PDF report generation requests.
func reportHandler(w http.ResponseWriter, r *http.Request, backendURL string) {
	// Extract student ID from the URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	studentID := parts[4]
	fmt.Println("Fetching report for student ID:", studentID)
	// Fetch student data from the Node.js backend
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/students/%s", backendURL, studentID))
	if err != nil {
		http.Error(w, "Failed to fetch student data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	var student Student
	if err := json.NewDecoder(resp.Body).Decode(&student); err != nil {
		http.Error(w, "Failed to parse student data", http.StatusInternalServerError)
		return
	}

	// Generate PDF
	pdfBytes, err := generatePDF(student)
	if err != nil {
		http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
		return
	}

	// Set headers and write response
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=\"report.pdf\"")
	w.Write(pdfBytes)
}

func main() {
	backendURL := "http://localhost:5007"
	http.HandleFunc("/api/v1/students/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) == 6 && parts[5] == "report" {
			reportHandler(w, r, backendURL)
		} else {
			http.NotFound(w, r)
		}
	})
	fmt.Println("Go service listening on :8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
