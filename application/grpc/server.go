package grpc

import (
	"my_grpc_do_zero/application/grpc/pb"
	"my_grpc_do_zero/application/grpc/service"
	"my_grpc_do_zero/model"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"google.golang.org/grpc"
)

// ProductList é a lista de novos produtos
var ProductList = model.NewProducts()

// StartGrpcServer inicializa um servidor GRPC
func StartGrpcServer() {
	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	
	grpcServer := grpc.NewServer()	

	// Adiciona o registro reflection para que o client evans encontre
	reflection.Register(grpcServer)

	productService := service.NewProductGrpcServer(ProductList)
	pb.RegisterProductServiceServer(grpcServer, productService)

	log.Println("gRPC Server has been started")
	
	// Passa o server para o metodo Serve da instancia do pacote grpc do google
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("could not connect in grpcServer: %s", err)
	}

}