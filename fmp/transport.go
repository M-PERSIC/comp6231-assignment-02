package fmp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	transport "github.com/go-kit/kit/transport/http"
)

// StartFMPServer starts the Fruit Month Price microservice HTTP server
func StartFMPServer(service FruitMonthPriceService, port string) error {
	getFruitPriceHandler := transport.NewServer(
		MakeGetFruitPriceMonthEndpoint(service),
		decodeFruitMonthPriceRequest,
		encodeResponse,
	)
	http.Handle("/fruit-price/", getFruitPriceHandler)
	fmt.Printf("Starting FMP server! (Port %s)\n", port)
	return http.ListenAndServe(":"+port, nil)
}

// decodeFruitMonthPriceRequest decodes the HTTP request into the request struct
func decodeFruitMonthPriceRequest(_ context.Context, r *http.Request) (any, error) {
	path := strings.TrimPrefix(r.URL.Path, "/fruit-price/")
	parts := strings.Split(path, "/")
	if len(parts) != 4 || parts[0] != "fruit" || parts[2] != "month" {
		return nil, fmt.Errorf("invalid URL format! (expected: /fruit-price/fruit/{fruit}/month/{month})")
	}
	fruit := parts[1]
	month := parts[3]
	if strings.TrimSpace(fruit) == "" {
		return nil, errors.New("fruit parameter is required")
	}
	if strings.TrimSpace(month) == "" {
		return nil, errors.New("month parameter is required")
	}
	return FruitMonthPriceRequest{
		Fruit: fruit,
		Month: month,
	}, nil
}

// encodeResponse encodes the response as JSON
func encodeResponse(_ context.Context, w http.ResponseWriter, response any) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
