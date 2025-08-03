package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/marcboeker/go-duckdb/v2"
)

// InitDB creates and returns a new DuckDB connection
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("duckdb", "")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database! %w", err)
	}
	return db, nil
}

// LoadExcelData creates a temporary file from embedded bytes and loads it
func LoadExcelData(db *sql.DB, excelData []byte, sheet, table string) error {
	temp, err := os.CreateTemp("", "FMP.xlsx")
	if err != nil {
		return fmt.Errorf("failed to create temporary file! %w", err)
	}
	defer os.Remove(temp.Name())
	if _, err := temp.Write(excelData); err != nil {
		return fmt.Errorf("failed to write data to temporary file! %w", err)
	}
	temp.Close()
	return LoadExcelFile(db, temp.Name(), sheet, table)
}

// LoadExcelFile loads the Excel file into DuckDB as a table
func LoadExcelFile(db *sql.DB, filename, sheet, table string) error {
	path, err := filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("failed to get absolute file path! %w", err)
	}
	query := fmt.Sprintf(`CREATE TABLE %s AS 
		SELECT * FROM read_xlsx('%s', sheet = '%s', header = true)`,
		table, path, sheet)
	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to load Excel file into table! %s - %w", table, err)
	}
	return nil
}

// QueryFruitPrice executes a basic query for fruit price by month
func QueryFruitPrice(db *sql.DB, table, fruit, month string) (float64, error) {
	query := fmt.Sprintf(`SELECT "%s" FROM %s WHERE LOWER("FRUIT/MONTH") = LOWER(?)`, month, table)
	var price float64
	err := db.QueryRow(query, fruit).Scan(&price)
	if err != nil {
		return 0, fmt.Errorf("database query failed! %w", err)
	}
	return price, nil
}
