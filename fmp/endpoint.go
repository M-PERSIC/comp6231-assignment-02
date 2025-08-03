package fmp

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetFruitPriceMonthEndpoint(service FruitMonthPriceService) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		req := request.(FruitMonthPriceRequest)
		resp, err := service.GetFruitPriceMonth(req.Fruit, req.Month)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
