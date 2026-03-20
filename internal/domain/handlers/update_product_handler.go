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

func UpdateProductHandler(
	ctx context.Context,
	req *productsv1.UpdateProductRequest,
	logger *slog.Logger,
	q *db_gen.Queries,
) (*productsv1.UpdateProductResponse, error) {
	logger.Debug("UpdateProduct request received", "product_id", req.GetProductId())

	if err := q.UpdateProduct(ctx, db_gen.UpdateProductParams{
		ID:             req.GetProductId(),
		Name:           req.GetName(),
		Description:    req.GetDescription(),
		Price:          req.GetPrice(),
		Category:       req.GetCategory(),
		VirtualImageID: req.GetVirtualImageId(),
		ModelID:        req.GetModelId(),
	}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Debug("product not found for update", "product_id", req.GetProductId())
			return nil, status.Errorf(codes.NotFound, "product %q not found", req.GetProductId())
		}
		logger.Error("UpdateProduct query failed", "product_id", req.GetProductId(), "error", err)
		return nil, status.Error(codes.Internal, "failed to update product")
	}

	updated, err := q.GetProductByID(ctx, req.GetProductId())
	if err != nil {
		logger.Error("GetProductByID after update failed", "product_id", req.GetProductId(), "error", err)
		return nil, status.Error(codes.Internal, "failed to fetch updated product")
	}

	logger.Debug("product updated", "product_id", updated.ID)
	return &productsv1.UpdateProductResponse{
		Product: &productsv1.Product{
			Id:             updated.ID,
			Name:           updated.Name,
			Description:    updated.Description,
			Price:          updated.Price,
			Category:       updated.Category,
			VirtualImageId: updated.VirtualImageID,
			ModelId:        updated.ModelID,
		},
	}, nil
}
