package library

import (
	"fmt"
	"strings"
	"time"
)

func ChangeFormatTime(timestamp string) (string, error) {
	// Parsing string timestamp ke dalam format waktu ISO 8601
	orderDate, err := time.Parse("2006-01-02T15:04:05.000000Z", timestamp)
	if err != nil {
		return "", fmt.Errorf("error parsing date: %v", err)
	}

	// Format tanggal ke bentuk "02 January 2006"
	formattedDate := orderDate.Format("02 January 2006")

	// Map nama bulan dari bahasa Inggris ke bahasa Indonesia
	indonesianMonth := map[string]string{
		"January":   "Januari",
		"February":  "Februari",
		"March":     "Maret",
		"April":     "April",
		"May":       "Mei",
		"June":      "Juni",
		"July":      "Juli",
		"August":    "Agustus",
		"September": "September",
		"October":   "Oktober",
		"November":  "November",
		"December":  "Desember",
	}

	// Mengganti nama bulan dalam bahasa Inggris ke bahasa Indonesia
	for english, indonesian := range indonesianMonth {
		formattedDate = strings.ReplaceAll(formattedDate, english, indonesian)
	}

	return formattedDate, nil
}

func main() {
	// Contoh penggunaan fungsi changeFormatTime
	orderDateString := "2024-11-08T00:04:45.804126Z"
	formattedDate, err := ChangeFormatTime(orderDateString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Formatted Date:", formattedDate) // Output: 08 November 2024
}
