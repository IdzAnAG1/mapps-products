package server

import (
	"context"
	"log/slog"

	productsv1 "mapps_product/generated/mobileapps/proto/products/v1"
	db_gen "mapps_product/internal/db/gen"
	"mapps_product/internal/domain/handlers"
)

type GrpcProductServer struct {
	productsv1.UnimplementedProductsServer
	Logger  *slog.Logger
	Queries *db_gen.Queries
}

func (gs *GrpcProductServer) GetProduct(ctx context.Context, req *productsv1.GetProductRequest) (*productsv1.GetProductResponse, error) {
	return handlers.GetProductHandler(ctx, req, gs.Logger, gs.Queries)
}

func (gs *GrpcProductServer) ListProducts(ctx context.Context, req *productsv1.ListProductsRequest) (*productsv1.ListProductsResponse, error) {
	return handlers.ListProductsHandler(ctx, req, gs.Logger, gs.Queries)
}

func (gs *GrpcProductServer) CreateProduct(ctx context.Context, req *productsv1.CreateProductRequest) (*productsv1.CreateProductResponse, error) {
	return handlers.CreateProductHandler(ctx, req, gs.Logger, gs.Queries)
}

func (gs *GrpcProductServer) UpdateProduct(ctx context.Context, req *productsv1.UpdateProductRequest) (*productsv1.UpdateProductResponse, error) {
	return handlers.UpdateProductHandler(ctx, req, gs.Logger, gs.Queries)
}
