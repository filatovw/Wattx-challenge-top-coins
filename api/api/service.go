package api

import (
	"context"
	"log"

	"github.com/filatovw/Wattx-challenge-top-coins/pricelist/pricelist"
	"github.com/pkg/errors"
)

type Service struct {
	log       *log.Logger
	pricelist pricelist.Client
}

func NewService(log *log.Logger, pricelist pricelist.Client) Service {
	return Service{
		log:       log,
		pricelist: pricelist,
	}
}

type GetPricelistResponse struct {
	Pricelist []*pricelist.Position
}

type GetPricelistRequest struct {
	Limit int32
}

func (s Service) GetPricelist(ctx context.Context, req *GetPricelistRequest) (*GetPricelistResponse, error) {
	if req.Limit < 1 || req.Limit > 200 {
		return nil, ErrBadRequest.With(errors.New("limit should be in interval [1:200]"))
	}
	resp, err := s.pricelist.GetPricelist(ctx, &pricelist.GetPricelistRequest{
		Limit: req.Limit,
	})
	if err != nil {
		return nil, ErrInternalServerError.With(errors.Wrap(err, "failed to fetch pricelist"))
	}
	return &GetPricelistResponse{Pricelist: resp.GetPositions()}, nil
}
