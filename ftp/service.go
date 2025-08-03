package ftp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/m-persic/comp6231-assignment-02/fmp"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type FruitTotalPriceRequest struct {
	Fruit    string `json:"fruit"`
	Month    string `json:"month"`
	Quantity uint   `json:"quantity"`
}

type FruitTotalPriceResponse struct {
	Fruit    string  `json:"fruit"`
	Month    string  `json:"month"`
	FMP      float64 `json:"fmp"`
	Quantity float64 `json:"quantity"`
	Total    float64 `json:"total"`
	Port     string  `json:"port"`
}

type FruitTotalPriceService interface {
	GetFruitPriceTotal(fruit, month string, quantity uint) (FruitTotalPriceResponse, error)
}

type service struct {
	port          string
	fmpServiceURL string
}

// NewService creates a new FMP service instance
func NewService(port string, fmpServiceUrl string) FruitTotalPriceService {
	return &service{
		port:          port,
		fmpServiceURL: fmpServiceUrl,
	}
}

// GetFruitPriceMonth retrieves the price for a specific fruit in a specific month
func (s *service) GetFruitPriceTotal(fruit, month string, quantity uint) (FruitTotalPriceResponse, error) {
	if strings.TrimSpace(fruit) == "" {
		return FruitTotalPriceResponse{}, errors.New("fruit name is required")
	}
	if strings.TrimSpace(month) == "" {
		return FruitTotalPriceResponse{}, errors.New("month is required")
	}
	month = cases.Title(language.English).String(cases.Lower(language.English).String(month))
	resp, err := http.Get(fmt.Sprintf("%s/fruit-price/fruit/%s/month/%s", s.fmpServiceURL, fruit, month))
	if err != nil {
		return FruitTotalPriceResponse{}, fmt.Errorf("failed to call FruitMonthPrice service: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return FruitTotalPriceResponse{}, fmt.Errorf("FMP service returned status: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return FruitTotalPriceResponse{}, fmt.Errorf("failed to read FruitMonthPrice service response: %w", err)
	}
	var fmpResponse fmp.FruitMonthPriceResponse
	err = json.Unmarshal(body, &fmpResponse)
	if err != nil {
		return FruitTotalPriceResponse{}, fmt.Errorf("failed to parse FruitMonthPrice service response: %w", err)
	}
	total := fmpResponse.FMP * float64(quantity)
	return FruitTotalPriceResponse{
		Fruit:    strings.ToLower(fruit),
		Month:    strings.ToLower(month),
		FMP:      fmpResponse.FMP,
		Quantity: float64(quantity),
		Total:    total,
		Port:     s.port,
	}, nil
}
