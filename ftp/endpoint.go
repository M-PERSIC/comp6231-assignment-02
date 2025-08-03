package ftp

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetFruitPriceMonthEndpoint(service FruitTotalPriceService) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		req := request.(FruitTotalPriceRequest)
		resp, err := service.GetFruitPriceTotal(req.Fruit, req.Month, req.Quantity)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
