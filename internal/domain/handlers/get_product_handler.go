package handlers

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	productsv1 "mapps_product/generated/mobileapps/proto/products/v1"
	db_gen "mapps_product/internal/db/gen"
)

func GetProductHandler(
	ctx context.Context,
	req *productsv1.GetProductRequest,
	logger *slog.Logger,
	q *db_gen.Queries,
) (*productsv1.GetProductResponse, error) {
	logger.Debug("GetProduct request received", "product_id", req.GetProductId())

	product, err := q.GetProductByID(ctx, req.GetProductId())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("product not found", "product_id", req.GetProductId())
			return nil, status.Errorf(codes.NotFound, "product %q not found", req.GetProductId())
		}
		logger.Error("GetProductByID failed", "product_id", req.GetProductId(), "error", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	logger.Debug("product fetched", "product_id", product.ID)
	return &productsv1.GetProductResponse{
		Product: &productsv1.Product{
			Id:             product.ID,
			Name:           product.Name,
			Description:    product.Description,
			Price:          product.Price,
			Category:       product.Category,
			VirtualImageId: product.VirtualImageID,
			ModelId:        product.ModelID,
		},
	}, nil
}
