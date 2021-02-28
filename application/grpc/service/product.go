package service

import (
	"time"
	"my_grpc_do_zero/model"
	"context"
	"my_grpc_do_zero/application/grpc/pb"
)

// ProductGrpcServer estrutura base
type ProductGrpcServer struct {
	pb.UnimplementedProductServiceServer
	Products *model.Products
}


// CreateProduct estrutura create product
func (p *ProductGrpcServer) CreateProduct(ctx context.Context, in *pb.Product) (*pb.ProductResult, error) {
	product := model.NewProduct()
	product.Name = in.Name
	p.Products.Add(product)

	return &pb.ProductResult{
		Id: product.ID,
		Name: product.Name,
	}, nil
}

// List de produtos
func (p *ProductGrpcServer) List(req *pb.Empty, stream pb.ProductService_ListServer) error{
	for _, product := range p.Products.Product {
		time.Sleep(time.Second * 5)
		// Retorna o resultado para o cliente
		stream.Send(&pb.ProductResult{Name: product.Name, Id:product.ID})
	}

	return nil
}

// NewProductGrpcServer novo produto 
func NewProductGrpcServer(products *model.Products) *ProductGrpcServer {
	return &ProductGrpcServer{
		Products: products,
	}
}