package handlers

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	productsv1 "mapps_product/generated/mobileapps/proto/products/v1"
	db_gen "mapps_product/internal/db/gen"
	"mapps_product/internal/domain/models"
)

func CreateProductHandler(
	ctx context.Context,
	req *productsv1.CreateProductRequest,
	logger *slog.Logger,
	q *db_gen.Queries,
) (*productsv1.CreateProductResponse, error) {
	logger.Debug("CreateProduct request received", "name", req.GetName())

	product := models.NewProduct(
		req.GetName(),
		req.GetDescription(),
		req.GetPrice(),
		req.GetCategory(),
		req.GetVirtualImageId(),
		req.GetModelId(),
	)

	if err := q.CreateProduct(ctx, db_gen.CreateProductParams{
		ID:             product.ID,
		Name:           product.Name,
		Description:    product.Description,
		Price:          product.Price,
		Category:       product.Category,
		VirtualImageID: product.VirtualImageID,
		ModelID:        product.ModelID,
	}); err != nil {
		logger.Error("CreateProduct query failed", "error", err)
		return nil, status.Error(codes.Internal, "failed to create product")
	}

	created, err := q.GetProductByID(ctx, product.ID)
	if err != nil {
		logger.Error("GetProductByID after create failed", "product_id", product.ID, "error", err)
		return nil, status.Error(codes.Internal, "failed to fetch created product")
	}

	logger.Debug("product created", "product_id", created.ID)
	return &productsv1.CreateProductResponse{
		Product: &productsv1.Product{
			Id:             created.ID,
			Name:           created.Name,
			Description:    created.Description,
			Price:          created.Price,
			Category:       created.Category,
			VirtualImageId: created.VirtualImageID,
			ModelId:        created.ModelID,
		},
	}, nil
}
