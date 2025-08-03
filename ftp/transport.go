package ftp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	transport "github.com/go-kit/kit/transport/http"
)

// StartFMPServer starts the Fruit Month Price microservice HTTP server
func StartFTPServer(service FruitTotalPriceService, port string) error {
	getFruitPriceHandler := transport.NewServer(
		MakeGetFruitPriceMonthEndpoint(service),
		decodeFruitTotalPriceRequest,
		encodeResponse,
	)
	http.Handle("/fruit-total/", getFruitPriceHandler)
	fmt.Printf("Starting FTP server! (Port %s)\n", port)
	return http.ListenAndServe(":"+port, nil)
}

// decodeFruitTotalPriceRequest decodes the HTTP request into the request struct
func decodeFruitTotalPriceRequest(_ context.Context, r *http.Request) (any, error) {
	path := strings.TrimPrefix(r.URL.Path, "/fruit-total/")
	parts := strings.Split(path, "/")
	if len(parts) != 6 || parts[0] != "fruit" || parts[2] != "month" || parts[4] != "quantity" {
		return nil, fmt.Errorf("invalid URL format! (expected: /fruit-total/fruit/{fruit}/month/{month}/quantity/{quantity})")
	}
	fruit := parts[1]
	month := parts[3]
	quantityStr := parts[5]
	if strings.TrimSpace(fruit) == "" {
		return nil, errors.New("fruit parameter is required")
	}
	if strings.TrimSpace(month) == "" {
		return nil, errors.New("month parameter is required")
	}
	if strings.TrimSpace(quantityStr) == "" {
		return nil, errors.New("quantity parameter is required")
	}
	quantity, err := strconv.ParseUint(quantityStr, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid quantity parameter: %s", quantityStr)
	}
	return FruitTotalPriceRequest{
		Fruit:    fruit,
		Month:    month,
		Quantity: uint(quantity),
	}, nil
}

// encodeResponse encodes the response as JSON
func encodeResponse(_ context.Context, w http.ResponseWriter, response any) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
