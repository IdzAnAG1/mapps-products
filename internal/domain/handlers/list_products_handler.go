package handlers

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	productsv1 "mapps_product/generated/mobileapps/proto/products/v1"
	db_gen "mapps_product/internal/db/gen"
)

func ListProductsHandler(
	ctx context.Context,
	req *productsv1.ListProductsRequest,
	logger *slog.Logger,
	q *db_gen.Queries,
) (*productsv1.ListProductsResponse, error) {
	logger.Debug("ListProducts request received", "category", req.GetCategory(), "page", req.GetPage())

	pageSize := int32(req.GetPageSize())
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}
	page := int32(req.GetPage())
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * pageSize

	var dbProducts []db_gen.Product
	var err error

	if req.GetCategory() != "" {
		dbProducts, err = q.ListProductsByCategory(ctx, db_gen.ListProductsByCategoryParams{
			Category: req.GetCategory(),
			Limit:    pageSize,
			Offset:   offset,
		})
	} else {
		dbProducts, err = q.ListProducts(ctx, db_gen.ListProductsParams{
			Limit:  pageSize,
			Offset: offset,
		})
	}
	if err != nil {
		logger.Error("ListProducts query failed", "error", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	products := make([]*productsv1.Product, 0, len(dbProducts))
	for _, p := range dbProducts {
		products = append(products, &productsv1.Product{
			Id:             p.ID,
			Name:           p.Name,
			Description:    p.Description,
			Price:          p.Price,
			Category:       p.Category,
			VirtualImageId: p.VirtualImageID,
			ModelId:        p.ModelID,
		})
	}

	logger.Debug("products listed", "count", len(products))
	return &productsv1.ListProductsResponse{Products: products}, nil
}
