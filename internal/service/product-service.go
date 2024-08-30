package service

import (
	"context"
	"fmt"

	pb "github.com/STO-KubSU/productpb"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	products map[string]*pb.GetProductResponse
}

func NewProductService() *ProductService {
	// Пример данных о продуктах
	products := map[string]*pb.GetProductResponse{
		"1": {Id: "1", Name: "Laptop", Price: 999.99},
		"2": {Id: "2", Name: "Smartphone", Price: 499.99},
	}
	return &ProductService{products: products}
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	product, exists := s.products[req.Id]
	if !exists {
		return nil, fmt.Errorf("product not found")
	}
	return product, nil
}

func (s *ProductService) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	var products []*pb.GetProductResponse
	for _, product := range s.products {
		products = append(products, product)
	}
	return &pb.ListProductsResponse{Products: products}, nil
}
