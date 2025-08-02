package fmp

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	db "github.com/m-persic/comp6231-assignment-02/database"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type FruitMonthPriceRequest struct {
	Fruit string `json:"fruit"`
	Month string `json:"month"`
}

type FruitMonthPriceResponse struct {
	Fruit string  `json:"fruit"`
	Month string  `json:"month"`
	FMP   float64 `json:"fmp"`
	Port  string  `json:"port"`
}

type FruitMonthPriceService interface {
	GetFruitPriceMonth(fruit, month string) (FruitMonthPriceResponse, error)
}

type service struct {
	db    *sql.DB
	table string
	port  string
}

// NewService creates a new FMP service instance
func NewService(db *sql.DB, table, port string) FruitMonthPriceService {
	return &service{
		db:    db,
		table: table,
		port:  port,
	}
}

// GetFruitPriceMonth retrieves the price for a specific fruit in a specific month
func (s *service) GetFruitPriceMonth(fruit, month string) (FruitMonthPriceResponse, error) {
	if strings.TrimSpace(fruit) == "" {
		return FruitMonthPriceResponse{}, errors.New("fruit name is required")
	}
	if strings.TrimSpace(month) == "" {
		return FruitMonthPriceResponse{}, errors.New("month is required")
	}
	month = cases.Title(language.English).String(cases.Lower(language.English).String(month))
	price, err := db.QueryFruitPrice(s.db, s.table, fruit, month)
	if err != nil {
		return FruitMonthPriceResponse{}, fmt.Errorf("failed to get price for %s in %s: %w", fruit, month, err)
	}
	return FruitMonthPriceResponse{
		Fruit: strings.ToLower(fruit),
		Month: strings.ToLower(month),
		FMP:   price,
		Port:  s.port,
	}, nil
}
